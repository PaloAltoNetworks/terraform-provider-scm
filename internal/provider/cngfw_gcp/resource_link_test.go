package provider_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/paloaltonetworks/terraform-provider-scm/internal/provider"
)

// fakeNexus is a mock auth server plus the SCM gateway routes the link resource calls. Like
// the real gateway, it takes the caller's TSG from the token, not from the request body.
type fakeNexus struct {
	t   *testing.T
	srv *httptest.Server

	mu        sync.Mutex
	generates []map[string]string
	regions   []string
	unlinked  []string
}

func newFakeNexus(t *testing.T) *fakeNexus {
	f := &fakeNexus{t: t}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth", func(w http.ResponseWriter, r *http.Request) {
		tsg := strings.TrimPrefix(r.FormValue("scope"), "tsg_id:")
		f.json(w, http.StatusOK, map[string]any{"access_token": "jwt-" + tsg, "token_type": "Bearer", "expires_in": 900})
	})
	mux.HandleFunc("POST /cngfw-gcp/v2/link/generate", func(w http.ResponseWriter, r *http.Request) {
		tsg, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer jwt-")
		if !ok || tsg == "" {
			f.json(w, http.StatusUnauthorized, map[string]any{"_errors": []map[string]any{{"message": "no token"}}})
			return
		}
		var body map[string]string
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.mu.Lock()
		f.generates = append(f.generates, body)
		f.regions = append(f.regions, r.Header.Get("x-panw-region"))
		n := len(f.generates)
		f.mu.Unlock()
		if tsg == "99999999999" {
			f.json(w, http.StatusBadRequest, map[string]any{
				"Response":       nil,
				"ResponseStatus": map[string]any{"ErrorCode": 1, "Code": "SCM_NOT_PROVISIONED", "Reason": ""},
			})
			return
		}
		// The service looks up the TSG's CSP account ID and returns it.
		f.json(w, http.StatusOK, map[string]any{
			"Response":       map[string]any{"Code": fmt.Sprintf("CODE%06d", n), "TSGID": tsg, "CSPID": "1853605", "ExpiresAt": 1792171508 + n},
			"ResponseStatus": map[string]any{"ErrorCode": 0, "Code": "OK", "Reason": ""},
		})
	})
	// Like the real service, unlink is keyed by the link code and rejects a request without one.
	mux.HandleFunc("POST /cngfw-gcp/v2/link/unlink", func(w http.ResponseWriter, r *http.Request) {
		var body map[string]string
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["Code"] == "" {
			f.json(w, http.StatusBadRequest, map[string]any{
				"Response":       nil,
				"ResponseStatus": map[string]any{"ErrorCode": 1, "Code": "FIELD_MISSING", "Reason": ""},
			})
			return
		}
		f.mu.Lock()
		f.unlinked = append(f.unlinked, body["Code"]+" "+body["GCPTenantID"])
		f.mu.Unlock()
		f.json(w, http.StatusOK, map[string]any{
			"Response":       map[string]any{"GCPTenantID": body["GCPTenantID"], "Unlinked": true},
			"ResponseStatus": map[string]any{"ErrorCode": 0, "Code": "OK", "Reason": ""},
		})
	})
	f.srv = httptest.NewServer(mux)
	t.Cleanup(f.srv.Close)
	return f
}

func (f *fakeNexus) json(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		f.t.Error(err)
	}
}

func (f *fakeNexus) providerConfig() string {
	return f.providerConfigWithScope("tsg_id:1778860705")
}

func (f *fakeNexus) providerConfigWithScope(scope string) string {
	u, _ := url.Parse(f.srv.URL)
	return fmt.Sprintf(`
provider "scm" {
  auth_url      = "%s/auth"
  protocol      = "http"
  host          = %q
  port          = %s
  client_id     = "sa@1778860705.iam.panserviceaccount.com"
  client_secret = "secret"
  scope         = %q
}
`, f.srv.URL, u.Hostname(), u.Port(), scope)
}

// checkUnlinked asserts the given "<code> <gcp tenant>" pairs were unlinked, in order.
func (f *fakeNexus) checkUnlinked(want ...string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		f.mu.Lock()
		defer f.mu.Unlock()
		if fmt.Sprint(f.unlinked) != fmt.Sprint(want) {
			return fmt.Errorf("unlinked %v, want %v", f.unlinked, want)
		}
		return nil
	}
}

// checkLastGenerate asserts the body and region of the most recent generate call.
func (f *fakeNexus) checkLastGenerate(want map[string]string, wantRegion string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		f.mu.Lock()
		defer f.mu.Unlock()
		if len(f.generates) == 0 {
			return fmt.Errorf("generate was never called")
		}
		got := f.generates[len(f.generates)-1]
		if fmt.Sprint(got) != fmt.Sprint(want) {
			return fmt.Errorf("generate body %v, want %v", got, want)
		}
		if region := f.regions[len(f.regions)-1]; region != wantRegion {
			return fmt.Errorf("x-panw-region %q, want %q", region, wantRegion)
		}
		return nil
	}
}

func protoV6Factories() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"scm": providerserver.NewProtocol6WithError(provider.New("test")),
	}
}

func TestLinkResourceLifecycle(t *testing.T) {
	f := newFakeNexus(t)
	const name = "scm_cngfw_gcp_link.test"

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6Factories(),
		CheckDestroy:             f.checkUnlinked("CODE000001 projects/proj-one", "CODE000002 organizations/123456789012"),
		Steps: []resource.TestStep{
			// Only the GCP tenant is sent; tsg_id and csp_id come back in the generate response.
			{
				Config: f.providerConfig() + `
resource "scm_cngfw_gcp_link" "test" {
  gcp_tenant_id = "projects/proj-one"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", "projects/proj-one"),
					resource.TestCheckResourceAttr(name, "tsg_id", "1778860705"),
					resource.TestCheckResourceAttr(name, "link_code", "CODE000001"),
					resource.TestCheckResourceAttr(name, "expires_at", "1792171509"),
					resource.TestCheckResourceAttr(name, "csp_id", "1853605"),
					resource.TestCheckNoResourceAttr(name, "region"),
					f.checkLastGenerate(map[string]string{"GCPTenantID": "projects/proj-one"}, ""),
				),
			},
			// Re-applying the same config must not mint a new code.
			{
				Config: f.providerConfig() + `
resource "scm_cngfw_gcp_link" "test" {
  gcp_tenant_id = "projects/proj-one"
}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
			// Changing an argument replaces the link: unlink the old code, generate for the new tenant.
			// The region is not validated; any value is passed through as the x-panw-region header.
			{
				Config: f.providerConfig() + `
resource "scm_cngfw_gcp_link" "test" {
  gcp_tenant_id = "organizations/123456789012"
  region        = "asia"
}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					f.checkUnlinked("CODE000001 projects/proj-one"),
					resource.TestCheckResourceAttr(name, "id", "organizations/123456789012"),
					resource.TestCheckResourceAttr(name, "tsg_id", "1778860705"),
					resource.TestCheckResourceAttr(name, "link_code", "CODE000002"),
					f.checkLastGenerate(map[string]string{"GCPTenantID": "organizations/123456789012"}, "asia"),
				),
			},
		},
	})
}

func TestLinkResourceErrors(t *testing.T) {
	f := newFakeNexus(t)

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6Factories(),
		Steps: []resource.TestStep{
			{
				Config: f.providerConfig() + `
resource "scm_cngfw_gcp_link" "test" {
  gcp_tenant_id = "Bad_Project"
}`,
				ExpectError: regexp.MustCompile(`must be organizations/\{number\}`),
			},
			{
				// The fake answers SCM_NOT_PROVISIONED for TSG 99999999999.
				Config: f.providerConfigWithScope("tsg_id:99999999999") + `
resource "scm_cngfw_gcp_link" "test" {
  gcp_tenant_id = "projects/proj-one"
}`,
				ExpectError: regexp.MustCompile(`SCM_NOT_PROVISIONED`),
			},
		},
	})
}

// The TSG and CSP account always come from the link service.
func TestLinkResourceRejectsDerivedInputs(t *testing.T) {
	f := newFakeNexus(t)

	var steps []resource.TestStep
	for _, attr := range []string{"tsg_id", "csp_id"} {
		steps = append(steps, resource.TestStep{
			Config: f.providerConfig() + fmt.Sprintf(`
resource "scm_cngfw_gcp_link" "test" {
  gcp_tenant_id = "projects/proj-one"
  %s = "1234"
}`, attr),
			ExpectError: regexp.MustCompile(`Invalid Configuration for Read-Only Attribute`),
		})
	}
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6Factories(),
		Steps:                    steps,
	})
}

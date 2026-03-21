package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/paloaltonetworks/scm-go/generated/config_setup"
	models "github.com/paloaltonetworks/terraform-provider-scm/internal/models/config_setup"
)

// --- Tests for packDevicesFromSdk ---

func Test_packDevicesFromSdk_AllFields(t *testing.T) {
	ctx := context.Background()

	sdk := config_setup.Devices{
		Id:              "test-uuid-1234",
		Name:            "00112233445566",
		Folder:          "Production",
		Description:     strPtr("Test firewall"),
		DisplayName:     strPtr("fw-prod-001"),
		Hostname:        strPtr("fw-prod-001.example.com"),
		IpAddress:       strPtr("10.0.0.1"),
		IpV6Address:     strPtr("fd00::1"),
		MacAddress:      strPtr("aa:bb:cc:dd:ee:ff"),
		Family:          strPtr("vm"),
		Model:           strPtr("PA-VM"),
		SoftwareVersion: strPtr("11.1.0"),
		IsConnected:     boolPtr(true),
		Labels:          []string{"production", "azure"},
		Snippets:        []string{"snippet-1", "snippet-2"},
	}

	model, diags := packDevicesFromSdk(ctx, sdk)
	if diags.HasError() {
		t.Fatalf("packDevicesFromSdk returned errors: %v", diags.Errors())
	}

	// Required fields
	assertEqual(t, "Id", "test-uuid-1234", model.Id.ValueString())
	assertEqual(t, "Name", "00112233445566", model.Name.ValueString())
	assertEqual(t, "Folder", "Production", model.Folder.ValueString())

	// Optional string pointer fields
	assertStringPtrEqual(t, "Description", "Test firewall", model.Description)
	assertStringPtrEqual(t, "DisplayName", "fw-prod-001", model.DisplayName)
	assertStringPtrEqual(t, "Hostname", "fw-prod-001.example.com", model.Hostname)
	assertStringPtrEqual(t, "IpAddress", "10.0.0.1", model.IpAddress)
	assertStringPtrEqual(t, "IpV6Address", "fd00::1", model.IpV6Address)
	assertStringPtrEqual(t, "MacAddress", "aa:bb:cc:dd:ee:ff", model.MacAddress)
	assertStringPtrEqual(t, "Family", "vm", model.Family)
	assertStringPtrEqual(t, "Model", "PA-VM", model.Model)
	assertStringPtrEqual(t, "SoftwareVersion", "11.1.0", model.SoftwareVersion)

	// Bool field
	if model.IsConnected.IsNull() || model.IsConnected.IsUnknown() {
		t.Errorf("IsConnected should not be null/unknown")
	}
	if model.IsConnected.ValueBool() != true {
		t.Errorf("IsConnected: expected true, got false")
	}

	// List fields
	if model.Labels.IsNull() {
		t.Errorf("Labels should not be null")
	}
	if len(model.Labels.Elements()) != 2 {
		t.Errorf("Labels: expected 2 elements, got %d", len(model.Labels.Elements()))
	}
	if model.Snippets.IsNull() {
		t.Errorf("Snippets should not be null")
	}
	if len(model.Snippets.Elements()) != 2 {
		t.Errorf("Snippets: expected 2 elements, got %d", len(model.Snippets.Elements()))
	}
}

func Test_packDevicesFromSdk_NilOptionalFields(t *testing.T) {
	ctx := context.Background()

	// Only required fields set, everything else nil
	sdk := config_setup.Devices{
		Id:     "test-uuid-5678",
		Name:   "99887766554433",
		Folder: "All Firewalls",
	}

	model, diags := packDevicesFromSdk(ctx, sdk)
	if diags.HasError() {
		t.Fatalf("packDevicesFromSdk returned errors: %v", diags.Errors())
	}

	// Required fields should be populated
	assertEqual(t, "Id", "test-uuid-5678", model.Id.ValueString())
	assertEqual(t, "Name", "99887766554433", model.Name.ValueString())
	assertEqual(t, "Folder", "All Firewalls", model.Folder.ValueString())

	// Optional fields should be null
	assertNull(t, "Description", model.Description)
	assertNull(t, "DisplayName", model.DisplayName)
	assertNull(t, "Hostname", model.Hostname)
	assertNull(t, "IpAddress", model.IpAddress)
	assertNull(t, "IpV6Address", model.IpV6Address)
	assertNull(t, "MacAddress", model.MacAddress)
	assertNull(t, "Family", model.Family)
	assertNull(t, "Model", model.Model)
	assertNull(t, "SoftwareVersion", model.SoftwareVersion)
	assertNull(t, "IsConnected (bool)", model.IsConnected)

	// Lists should be null when nil
	if !model.Labels.IsNull() {
		t.Errorf("Labels should be null when SDK value is nil")
	}
	if !model.Snippets.IsNull() {
		t.Errorf("Snippets should be null when SDK value is nil")
	}
}

func Test_packDevicesFromSdk_EmptyLists(t *testing.T) {
	ctx := context.Background()

	sdk := config_setup.Devices{
		Id:       "test-uuid-9999",
		Name:     "emptylistdevice",
		Folder:   "TestFolder",
		Labels:   []string{},
		Snippets: []string{},
	}

	model, diags := packDevicesFromSdk(ctx, sdk)
	if diags.HasError() {
		t.Fatalf("packDevicesFromSdk returned errors: %v", diags.Errors())
	}

	// Empty lists (non-nil) should produce empty TF lists, not null
	if model.Labels.IsNull() {
		t.Errorf("Labels should not be null for an empty (non-nil) slice")
	}
	if len(model.Labels.Elements()) != 0 {
		t.Errorf("Labels: expected 0 elements, got %d", len(model.Labels.Elements()))
	}
	if model.Snippets.IsNull() {
		t.Errorf("Snippets should not be null for an empty (non-nil) slice")
	}
	if len(model.Snippets.Elements()) != 0 {
		t.Errorf("Snippets: expected 0 elements, got %d", len(model.Snippets.Elements()))
	}
}

// --- Tests for unpackDevicesToSdkPut ---

func Test_unpackDevicesToSdkPut_AllFields(t *testing.T) {
	ctx := context.Background()

	labels, _ := basetypes.NewListValueFrom(ctx, basetypes.StringType{}, []string{"label1", "label2"})
	snippets, _ := basetypes.NewListValueFrom(ctx, basetypes.StringType{}, []string{"snip1"})

	model := &devicesTfForTest{
		Folder:      basetypes.NewStringValue("Azure Firewalls/Production"),
		Description: basetypes.NewStringValue("Moved to production"),
		DisplayName: basetypes.NewStringValue("fw-prod-001"),
		Labels:      labels,
		Snippets:    snippets,
	}

	sdkPut, diags := unpackDevicesToSdkPut(ctx, model.toDevicesTf())
	if diags.HasError() {
		t.Fatalf("unpackDevicesToSdkPut returned errors: %v", diags.Errors())
	}

	if sdkPut.Folder == nil || *sdkPut.Folder != "Azure Firewalls/Production" {
		t.Errorf("Folder: expected 'Azure Firewalls/Production', got %v", sdkPut.Folder)
	}
	if sdkPut.Description == nil || *sdkPut.Description != "Moved to production" {
		t.Errorf("Description: expected 'Moved to production', got %v", sdkPut.Description)
	}
	if sdkPut.DisplayName == nil || *sdkPut.DisplayName != "fw-prod-001" {
		t.Errorf("DisplayName: expected 'fw-prod-001', got %v", sdkPut.DisplayName)
	}
	if len(sdkPut.Labels) != 2 {
		t.Errorf("Labels: expected 2 elements, got %d", len(sdkPut.Labels))
	}
	if len(sdkPut.Snippets) != 1 {
		t.Errorf("Snippets: expected 1 element, got %d", len(sdkPut.Snippets))
	}
}

func Test_unpackDevicesToSdkPut_NullOptionalFields(t *testing.T) {
	ctx := context.Background()

	model := &devicesTfForTest{
		Folder:      basetypes.NewStringValue("SomeFolder"),
		Description: basetypes.NewStringNull(),
		DisplayName: basetypes.NewStringNull(),
		Labels:      basetypes.NewListNull(basetypes.StringType{}),
		Snippets:    basetypes.NewListNull(basetypes.StringType{}),
	}

	sdkPut, diags := unpackDevicesToSdkPut(ctx, model.toDevicesTf())
	if diags.HasError() {
		t.Fatalf("unpackDevicesToSdkPut returned errors: %v", diags.Errors())
	}

	if sdkPut.Folder == nil || *sdkPut.Folder != "SomeFolder" {
		t.Errorf("Folder: expected 'SomeFolder', got %v", sdkPut.Folder)
	}
	if sdkPut.Description != nil {
		t.Errorf("Description: expected nil, got %v", sdkPut.Description)
	}
	if sdkPut.DisplayName != nil {
		t.Errorf("DisplayName: expected nil, got %v", sdkPut.DisplayName)
	}
	if sdkPut.Labels != nil {
		t.Errorf("Labels: expected nil, got %v", sdkPut.Labels)
	}
	if sdkPut.Snippets != nil {
		t.Errorf("Snippets: expected nil, got %v", sdkPut.Snippets)
	}
}

// --- Tests for packDevicesListFromSdk ---

func Test_packDevicesListFromSdk_MultipleDevices(t *testing.T) {
	ctx := context.Background()

	sdks := []config_setup.Devices{
		{Id: "id-1", Name: "serial-1", Folder: "Folder-A", Description: strPtr("desc-1")},
		{Id: "id-2", Name: "serial-2", Folder: "Folder-B"},
		{Id: "id-3", Name: "serial-3", Folder: "Folder-A", IsConnected: boolPtr(false)},
	}

	tfList, diags := packDevicesListFromSdk(ctx, sdks)
	if diags.HasError() {
		t.Fatalf("packDevicesListFromSdk returned errors: %v", diags.Errors())
	}

	if tfList.IsNull() {
		t.Fatal("Expected a non-null list")
	}
	if len(tfList.Elements()) != 3 {
		t.Errorf("Expected 3 elements in list, got %d", len(tfList.Elements()))
	}
}

func Test_packDevicesListFromSdk_EmptyList(t *testing.T) {
	ctx := context.Background()

	tfList, diags := packDevicesListFromSdk(ctx, []config_setup.Devices{})
	if diags.HasError() {
		t.Fatalf("packDevicesListFromSdk returned errors: %v", diags.Errors())
	}

	// An empty slice produces a list with 0 elements (may be null depending on
	// how NewListValueFrom handles a nil internal slice — both are acceptable).
	if !tfList.IsNull() && len(tfList.Elements()) != 0 {
		t.Errorf("Expected 0 elements or null list, got %d elements", len(tfList.Elements()))
	}
}

// --- Tests for round-trip pack/unpack ---

func Test_RoundTrip_PackThenUnpack(t *testing.T) {
	ctx := context.Background()

	// Start with SDK Devices (as returned from API read)
	original := config_setup.Devices{
		Id:          "round-trip-id",
		Name:        "round-trip-serial",
		Folder:      "TestFolder/SubFolder",
		Description: strPtr("round trip test"),
		DisplayName: strPtr("Round Trip FW"),
		Labels:      []string{"env:test"},
		Snippets:    []string{"s1", "s2"},
	}

	// Pack to TF model
	model, diags := packDevicesFromSdk(ctx, original)
	if diags.HasError() {
		t.Fatalf("packDevicesFromSdk returned errors: %v", diags.Errors())
	}

	// Unpack back to SDK DevicesPut (the mutable fields)
	sdkPut, diags := unpackDevicesToSdkPut(ctx, model)
	if diags.HasError() {
		t.Fatalf("unpackDevicesToSdkPut returned errors: %v", diags.Errors())
	}

	// Verify the mutable fields survived the round trip
	if *sdkPut.Folder != "TestFolder/SubFolder" {
		t.Errorf("Folder round-trip failed: expected 'TestFolder/SubFolder', got '%s'", *sdkPut.Folder)
	}
	if *sdkPut.Description != "round trip test" {
		t.Errorf("Description round-trip failed: expected 'round trip test', got '%s'", *sdkPut.Description)
	}
	if *sdkPut.DisplayName != "Round Trip FW" {
		t.Errorf("DisplayName round-trip failed: expected 'Round Trip FW', got '%s'", *sdkPut.DisplayName)
	}
	if len(sdkPut.Labels) != 1 || sdkPut.Labels[0] != "env:test" {
		t.Errorf("Labels round-trip failed: expected [env:test], got %v", sdkPut.Labels)
	}
	if len(sdkPut.Snippets) != 2 || sdkPut.Snippets[0] != "s1" || sdkPut.Snippets[1] != "s2" {
		t.Errorf("Snippets round-trip failed: expected [s1, s2], got %v", sdkPut.Snippets)
	}
}

// --- Test helpers ---

// devicesTfForTest is a helper to build DevicesTf instances for testing
// without needing to set all the computed fields
type devicesTfForTest struct {
	Folder      basetypes.StringValue
	Description basetypes.StringValue
	DisplayName basetypes.StringValue
	Labels      basetypes.ListValue
	Snippets    basetypes.ListValue
}

func (d *devicesTfForTest) toDevicesTf() *models.DevicesTf {
	return &models.DevicesTf{
		Id:          basetypes.NewStringValue("test-id"),
		Name:        basetypes.NewStringNull(),
		Folder:      d.Folder,
		Description: d.Description,
		DisplayName: d.DisplayName,
		Labels:      d.Labels,
		Snippets:    d.Snippets,
		Hostname:    basetypes.NewStringNull(),
		IpAddress:   basetypes.NewStringNull(),
		IpV6Address: basetypes.NewStringNull(),
		MacAddress:  basetypes.NewStringNull(),
		Family:      basetypes.NewStringNull(),
		Model:       basetypes.NewStringNull(),
		SoftwareVersion: basetypes.NewStringNull(),
		IsConnected: basetypes.NewBoolNull(),
	}
}

func strPtr(s string) *string { return &s }
func boolPtr(b bool) *bool    { return &b }

func assertEqual(t *testing.T, field, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Errorf("%s: expected %q, got %q", field, expected, actual)
	}
}

func assertStringPtrEqual(t *testing.T, field, expected string, actual basetypes.StringValue) {
	t.Helper()
	if actual.IsNull() || actual.IsUnknown() {
		t.Errorf("%s: expected %q, got null/unknown", field, expected)
		return
	}
	if actual.ValueString() != expected {
		t.Errorf("%s: expected %q, got %q", field, expected, actual.ValueString())
	}
}

func assertNull(t *testing.T, field string, actual interface{ IsNull() bool }) {
	t.Helper()
	if !actual.IsNull() {
		t.Errorf("%s: expected null, got non-null", field)
	}
}

Terraform Provider for Palo Alto Networks Strata Cloud Manager API
==================================================================

**NOTE**: This provider is auto-generated.

---
## Beta Release Disclaimer

**This software is a pre-release version and is not ready for production use.**

*   **No Warranty:** This software is provided "as is," without any warranty of any kind, either expressed or implied, including, but not limited to, the implied warranties of merchantability and fitness for a particular purpose.
*   **Instability:** The beta software may contain defects, may not operate correctly, and may be substantially modified or withdrawn at any time.
*   **Limitation of Liability:** In no event shall the authors or copyright holders be liable for any claim, damages, or other liability, whether in an action of contract, tort, or otherwise, arising from, out of, or in connection with the beta software or the use or other dealings in the beta software.
*   **Feedback:** We encourage and appreciate your feedback and bug reports. However, you acknowledge that any feedback you provide is non-confidential.

By using this software, you agree to these terms.
---

Warranty
--------

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

THIS SOFTWARE IS RELEASED AS A PROOF OF CONCEPT FOR EXPERIMENTAL PURPOSES ONLY. USE IT AT OWN RISK. THIS SOFTWARE IS NOT SUPPORTED.

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) v1+
- [Go](https://go.dev) v1.21+ (to build the provider from source)


Building the Provider
---------------------

1. Install [Go](https://go.dev/dl)

2. Clone the SDK repo:

```sh
git clone https://github.com/paloaltonetworks/scm-go
```

3. Clone this repo:

```sh
git clone https://github.com/paloaltonetworks/terraform-provider-scm
```

4. Build the provider:

```sh
cd terraform-provider-scm
go build
```

5. Specify the `dev_overrides` configuration per the next section below. This tells Terraform where to find the provider you just built. The directory to specify is the full path to the cloned provider repo.


Developing the Provider
-----------------------

With Terraform v1 and later, [development overrides for provider developers](https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers) can be leveraged in order to use the provider built from source.

To do this, populate a Terraform CLI configuration file (`~/.terraformrc` for all platforms other than Windows; `terraform.rc` in the `%APPDATA%` directory when using Windows) with at least the following options:

```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/paloaltonetworks-local/scm" = "/directory/containing/the/provider/binary/here"
  }

  direct {}
}
```

Then when referencing the locally built provider, use the local name in the `terraform` configuration block like so:

```hcl
terraform {
    required_providers {
        scm = {
            source = "paloaltonetworks-local/scm"
            version = "1.0.0"
        }
    }
}
```

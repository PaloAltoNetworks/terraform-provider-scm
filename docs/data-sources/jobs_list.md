---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_jobs_list Data Source - scm"
subcategory: "NetSec"
description: |-
  Retrieves a listing of config items.
---

# scm_jobs_list (Data Source)

Retrieves a listing of config items.

## Example Usage

```terraform
data "scm_jobs_list" "example" {
  folder = "Shared"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `data` (Attributes List) The Data param. (see [below for nested schema](#nestedatt--data))
- `limit` (Number) The Limit param. A limit of -1 will return all configured items. Default: `200`.
- `offset` (Number) The Offset param. Default: `0`.
- `tfid` (String) The Terraform ID.
- `total` (Number) The Total param.

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `description` (String) A description provided by the administrator or service account.
- `device_name` (String) The name of the device.
- `end_ts` (String) The timestamp indicating when the job was finished.
- `id` (Number) The job ID.
- `job_result` (Number) The job result.
- `job_status` (Number) The current status of the job.
- `job_type` (Number) The job type.
- `parent_id` (Number) The parent job ID.
- `percent` (Number) Job completion percentage. Value must be less than or equal to 100.
- `result_str` (String) The result of the job. String must be one of these: `"OK"`, `"FAIL"`, `"PEND"`, `"WAIT"`, `"CANCELLED"`.
- `start_ts` (String) The timestamp indicating when the job was created.
- `status_str` (String) The current status of the job. String must be one of these: `"ACT"`, `"FIN"`, `"PEND"`, `"PUSHSENT"`, `"PUSHFAIL"`.
- `summary` (String) The completion summary of the job.
- `type_str` (String) The job type. String must be one of these: `"CommitAll"`, `"CommitAndPush"`, `"NGFW-Bootstrap-Push"`, `"Validate"`.
- `uname` (String) The administrator or service account that created the job.

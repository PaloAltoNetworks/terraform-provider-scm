---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scm_schedule Resource - scm"
subcategory: "NetSec"
description: |-
  Retrieves a config item.
---

# scm_schedule (Resource)

Retrieves a config item.

## Example Usage

```terraform
resource "scm_schedule" "example" {
  # Resource params
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Alphanumeric string [ 0-9a-zA-Z._-]. String length must not exceed 31 characters.
- `schedule_type` (Attributes) The ScheduleType param. (see [below for nested schema](#nestedatt--schedule_type))

### Optional

- `device` (String) The Device param.
- `folder` (String) The Folder param.
- `snippet` (String) The Snippet param.

### Read-Only

- `id` (String) UUID of the resource.
- `tfid` (String) The Terraform ID.

<a id="nestedatt--schedule_type"></a>
### Nested Schema for `schedule_type`

Optional:

- `non_recurring_list` (List of String) The NonRecurringList param. Individual elements in this list are subject to additional validation. String length must be between 33 and 33 characters. String validation regex: `[0-9][0-9][0-9][0-9]\/([0][1-9]|[1][0-2])\/([0-2][0-9]|[3][0-1])@([01][0-9]|[2][0-3]):([0-5][0-9])-[0-9][0-9][0-9][0-9]\/([0][1-9]|[1][0-2])\/([0-2][0-9]|[3][0-1])@([01][0-9]|[2][0-3]):([0-5][0-9])`. Ensure that only one of the following is specified: `non_recurring`, `recurring`
- `recurring` (Attributes) The Recurring param. Ensure that only one of the following is specified: `non_recurring`, `recurring` (see [below for nested schema](#nestedatt--schedule_type--recurring))

<a id="nestedatt--schedule_type--recurring"></a>
### Nested Schema for `schedule_type.recurring`

Optional:

- `daily_list` (List of String) The DailyList param. Individual elements in this list are subject to additional validation. String length must be between 11 and 11 characters. String validation regex: `([01][0-9]|[2][0-3]):([0-5][0-9])-([01][0-9]|[2][0-3]):([0-5][0-9])`. Ensure that only one of the following is specified: `daily`, `weekly`
- `weekly` (Attributes) The Weekly param. Ensure that only one of the following is specified: `daily`, `weekly` (see [below for nested schema](#nestedatt--schedule_type--recurring--weekly))

<a id="nestedatt--schedule_type--recurring--weekly"></a>
### Nested Schema for `schedule_type.recurring.weekly`

Optional:

- `friday_list` (List of String) The FridayList param. Individual elements in this list are subject to additional validation. String length must be between 11 and 11 characters. String validation regex: `([01][0-9]|[2][0-3]):([0-5][0-9])-([01][0-9]|[2][0-3]):([0-5][0-9])`.
- `monday_list` (List of String) The MondayList param. Individual elements in this list are subject to additional validation. String length must be between 11 and 11 characters. String validation regex: `([01][0-9]|[2][0-3]):([0-5][0-9])-([01][0-9]|[2][0-3]):([0-5][0-9])`.
- `saturday_list` (List of String) The SaturdayList param. Individual elements in this list are subject to additional validation. String length must be between 11 and 11 characters. String validation regex: `([01][0-9]|[2][0-3]):([0-5][0-9])-([01][0-9]|[2][0-3]):([0-5][0-9])`.
- `sunday_list` (List of String) The SundayList param. Individual elements in this list are subject to additional validation. String length must be between 11 and 11 characters. String validation regex: `([01][0-9]|[2][0-3]):([0-5][0-9])-([01][0-9]|[2][0-3]):([0-5][0-9])`.
- `thursday_list` (List of String) The ThursdayList param. Individual elements in this list are subject to additional validation. String length must be between 11 and 11 characters. String validation regex: `([01][0-9]|[2][0-3]):([0-5][0-9])-([01][0-9]|[2][0-3]):([0-5][0-9])`.
- `tuesday_list` (List of String) The TuesdayList param. Individual elements in this list are subject to additional validation. String length must be between 11 and 11 characters. String validation regex: `([01][0-9]|[2][0-3]):([0-5][0-9])-([01][0-9]|[2][0-3]):([0-5][0-9])`.
- `wednesday_list` (List of String) The WednesdayList param. Individual elements in this list are subject to additional validation. String length must be between 11 and 11 characters. String validation regex: `([01][0-9]|[2][0-3]):([0-5][0-9])-([01][0-9]|[2][0-3]):([0-5][0-9])`.

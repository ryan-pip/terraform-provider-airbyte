---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_twilio_taskrouter Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceTwilioTaskrouter DataSource
---

# airbyte_source_twilio_taskrouter (Data Source)

SourceTwilioTaskrouter DataSource

## Example Usage

```terraform
data "airbyte_source_twilio_taskrouter" "my_source_twiliotaskrouter" {
  secret_id = "...my_secret_id..."
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `account_sid` (String) Twilio Account ID
- `auth_token` (String) Twilio Auth Token
- `source_type` (String) must be one of ["twilio-taskrouter"]



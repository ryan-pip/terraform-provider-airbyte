---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_twilio_taskrouter Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceTwilioTaskrouter Resource
---

# airbyte_source_twilio_taskrouter (Resource)

SourceTwilioTaskrouter Resource

## Example Usage

```terraform
resource "airbyte_source_twilio_taskrouter" "my_source_twiliotaskrouter" {
  configuration = {
    account_sid = "...my_account_sid..."
    auth_token  = "...my_auth_token..."
    source_type = "twilio-taskrouter"
  }
  name         = "Marta Hodkiewicz"
  secret_id    = "...my_secret_id..."
  workspace_id = "8c99c722-d2bc-40f9-8087-d9caae042dd7"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `account_sid` (String) Twilio Account ID
- `auth_token` (String) Twilio Auth Token
- `source_type` (String) must be one of ["twilio-taskrouter"]



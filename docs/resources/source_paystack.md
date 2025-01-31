---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_paystack Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePaystack Resource
---

# airbyte_source_paystack (Resource)

SourcePaystack Resource

## Example Usage

```terraform
resource "airbyte_source_paystack" "my_source_paystack" {
  configuration = {
    lookback_window_days = 7
    secret_key           = "...my_secret_key..."
    source_type          = "paystack"
    start_date           = "2017-01-25T00:00:00Z"
  }
  name         = "Mattie Gutkowski"
  secret_id    = "...my_secret_id..."
  workspace_id = "99257d04-f408-447a-b42d-84496cbdeecf"
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

- `secret_key` (String) The Paystack API key (usually starts with 'sk_live_'; find yours <a href="https://dashboard.paystack.com/#/settings/developer">here</a>).
- `source_type` (String) must be one of ["paystack"]
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.

Optional:

- `lookback_window_days` (Number) When set, the connector will always reload data from the past N days, where N is the value set here. This is useful if your data is updated after creation.



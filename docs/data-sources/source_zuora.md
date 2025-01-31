---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_zuora Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceZuora DataSource
---

# airbyte_source_zuora (Data Source)

SourceZuora DataSource

## Example Usage

```terraform
data "airbyte_source_zuora" "my_source_zuora" {
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

- `client_id` (String) Your OAuth user Client ID
- `client_secret` (String) Your OAuth user Client Secret
- `data_query` (String) must be one of ["Live", "Unlimited"]
Choose between `Live`, or `Unlimited` - the optimized, replicated database at 12 hours freshness for high volume extraction <a href="https://knowledgecenter.zuora.com/Central_Platform/Query/Data_Query/A_Overview_of_Data_Query#Query_Processing_Limitations">Link</a>
- `source_type` (String) must be one of ["zuora"]
- `start_date` (String) Start Date in format: YYYY-MM-DD
- `tenant_endpoint` (String) must be one of ["US Production", "US Cloud Production", "US API Sandbox", "US Cloud API Sandbox", "US Central Sandbox", "US Performance Test", "EU Production", "EU API Sandbox", "EU Central Sandbox"]
Please choose the right endpoint where your Tenant is located. More info by this <a href="https://www.zuora.com/developer/api-reference/#section/Introduction/Access-to-the-API">Link</a>
- `window_in_days` (String) The amount of days for each data-chunk begining from start_date. Bigger the value - faster the fetch. (0.1 - as for couple of hours, 1 - as for a Day; 364 - as for a Year).



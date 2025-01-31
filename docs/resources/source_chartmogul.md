---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_chartmogul Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceChartmogul Resource
---

# airbyte_source_chartmogul (Resource)

SourceChartmogul Resource

## Example Usage

```terraform
resource "airbyte_source_chartmogul" "my_source_chartmogul" {
  configuration = {
    api_key     = "...my_api_key..."
    interval    = "quarter"
    source_type = "chartmogul"
    start_date  = "2017-01-25T00:00:00Z"
  }
  name         = "Mr. Kristopher Torphy"
  secret_id    = "...my_secret_id..."
  workspace_id = "0ce2169e-5100-419c-adc5-e34762799bfb"
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

- `api_key` (String) Your Chartmogul API key. See <a href="https://help.chartmogul.com/hc/en-us/articles/4407796325906-Creating-and-Managing-API-keys#creating-an-api-key"> the docs </a> for info on how to obtain this.
- `interval` (String) must be one of ["day", "week", "month", "quarter"]
Some APIs such as <a href="https://dev.chartmogul.com/reference/endpoint-overview-metrics-api">Metrics</a> require intervals to cluster data.
- `source_type` (String) must be one of ["chartmogul"]
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. When feasible, any data before this date will not be replicated.



---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_us_census Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceUsCensus Resource
---

# airbyte_source_us_census (Resource)

SourceUsCensus Resource

## Example Usage

```terraform
resource "airbyte_source_us_census" "my_source_uscensus" {
  configuration = {
    api_key      = "...my_api_key..."
    query_params = "get=NAME,NAICS2017_LABEL,LFO_LABEL,EMPSZES_LABEL,ESTAB,PAYANN,PAYQTR1,EMP&for=us:*&NAICS2017=72&LFO=001&EMPSZES=001"
    query_path   = "data/timeseries/healthins/sahie"
    source_type  = "us-census"
  }
  name         = "Brandon Rogahn"
  secret_id    = "...my_secret_id..."
  workspace_id = "2a00bef6-9e10-4015-b630-bda7afded84a"
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

- `api_key` (String) Your API Key. Get your key <a href="https://api.census.gov/data/key_signup.html">here</a>.
- `query_path` (String) The path portion of the GET request
- `source_type` (String) must be one of ["us-census"]

Optional:

- `query_params` (String) The query parameters portion of the GET request, without the api key



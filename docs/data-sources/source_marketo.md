---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_marketo Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceMarketo DataSource
---

# airbyte_source_marketo (Data Source)

SourceMarketo DataSource

## Example Usage

```terraform
data "airbyte_source_marketo" "my_source_marketo" {
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

- `client_id` (String) The Client ID of your Marketo developer application. See <a href="https://docs.airbyte.com/integrations/sources/marketo"> the docs </a> for info on how to obtain this.
- `client_secret` (String) The Client Secret of your Marketo developer application. See <a href="https://docs.airbyte.com/integrations/sources/marketo"> the docs </a> for info on how to obtain this.
- `domain_url` (String) Your Marketo Base URL. See <a href="https://docs.airbyte.com/integrations/sources/marketo"> the docs </a> for info on how to obtain this.
- `source_type` (String) must be one of ["marketo"]
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.



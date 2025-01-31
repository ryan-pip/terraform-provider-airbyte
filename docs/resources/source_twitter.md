---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_twitter Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceTwitter Resource
---

# airbyte_source_twitter (Resource)

SourceTwitter Resource

## Example Usage

```terraform
resource "airbyte_source_twitter" "my_source_twitter" {
  configuration = {
    api_key     = "...my_api_key..."
    end_date    = "2020-12-31T16:49:13.658Z"
    query       = "...my_query..."
    source_type = "twitter"
    start_date  = "2021-05-21T04:49:52.479Z"
  }
  name         = "Marco Gleichner"
  secret_id    = "...my_secret_id..."
  workspace_id = "a1cfe9e1-5df9-4039-87f3-7831983d42e5"
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

- `api_key` (String) App only Bearer Token. See the <a href="https://developer.twitter.com/en/docs/authentication/oauth-2-0/bearer-tokens">docs</a> for more information on how to obtain this token.
- `query` (String) Query for matching Tweets. You can learn how to build this query by reading <a href="https://developer.twitter.com/en/docs/twitter-api/tweets/search/integrate/build-a-query"> build a query guide </a>.
- `source_type` (String) must be one of ["twitter"]

Optional:

- `end_date` (String) The end date for retrieving tweets must be a minimum of 10 seconds prior to the request time.
- `start_date` (String) The start date for retrieving tweets cannot be more than 7 days in the past.



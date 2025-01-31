---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_airtable Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAirtable Resource
---

# airbyte_source_airtable (Resource)

SourceAirtable Resource

## Example Usage

```terraform
resource "airbyte_source_airtable" "my_source_airtable" {
  configuration = {
    credentials = {
      source_airtable_authentication_o_auth2_0 = {
        access_token      = "...my_access_token..."
        auth_method       = "oauth2.0"
        client_id         = "...my_client_id..."
        client_secret     = "...my_client_secret..."
        refresh_token     = "...my_refresh_token..."
        token_expiry_date = "2020-12-11T09:39:15.481Z"
      }
    }
    source_type = "airtable"
  }
  name         = "Vincent Frami"
  secret_id    = "...my_secret_id..."
  workspace_id = "c2beb477-373c-48d7-af64-d1db1f2c4310"
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

Optional:

- `credentials` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials))
- `source_type` (String) must be one of ["airtable"]

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `source_airtable_authentication_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_airtable_authentication_o_auth2_0))
- `source_airtable_authentication_personal_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_airtable_authentication_personal_access_token))
- `source_airtable_update_authentication_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_airtable_update_authentication_o_auth2_0))
- `source_airtable_update_authentication_personal_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_airtable_update_authentication_personal_access_token))

<a id="nestedatt--configuration--credentials--source_airtable_authentication_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_airtable_authentication_o_auth2_0`

Required:

- `client_id` (String) The client ID of the Airtable developer application.
- `client_secret` (String) The client secret the Airtable developer application.
- `refresh_token` (String) The key to refresh the expired access token.

Optional:

- `access_token` (String) Access Token for making authenticated requests.
- `auth_method` (String) must be one of ["oauth2.0"]
- `token_expiry_date` (String) The date-time when the access token should be refreshed.


<a id="nestedatt--configuration--credentials--source_airtable_authentication_personal_access_token"></a>
### Nested Schema for `configuration.credentials.source_airtable_authentication_personal_access_token`

Required:

- `api_key` (String) The Personal Access Token for the Airtable account. See the <a href="https://airtable.com/developers/web/guides/personal-access-tokens">Support Guide</a> for more information on how to obtain this token.

Optional:

- `auth_method` (String) must be one of ["api_key"]


<a id="nestedatt--configuration--credentials--source_airtable_update_authentication_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_airtable_update_authentication_o_auth2_0`

Required:

- `client_id` (String) The client ID of the Airtable developer application.
- `client_secret` (String) The client secret the Airtable developer application.
- `refresh_token` (String) The key to refresh the expired access token.

Optional:

- `access_token` (String) Access Token for making authenticated requests.
- `auth_method` (String) must be one of ["oauth2.0"]
- `token_expiry_date` (String) The date-time when the access token should be refreshed.


<a id="nestedatt--configuration--credentials--source_airtable_update_authentication_personal_access_token"></a>
### Nested Schema for `configuration.credentials.source_airtable_update_authentication_personal_access_token`

Required:

- `api_key` (String) The Personal Access Token for the Airtable account. See the <a href="https://airtable.com/developers/web/guides/personal-access-tokens">Support Guide</a> for more information on how to obtain this token.

Optional:

- `auth_method` (String) must be one of ["api_key"]



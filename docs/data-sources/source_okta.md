---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_okta Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceOkta DataSource
---

# airbyte_source_okta (Data Source)

SourceOkta DataSource

## Example Usage

```terraform
data "airbyte_source_okta" "my_source_okta" {
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

- `credentials` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials))
- `domain` (String) The Okta domain. See the <a href="https://docs.airbyte.com/integrations/sources/okta">docs</a> for instructions on how to find it.
- `source_type` (String) must be one of ["okta"]
- `start_date` (String) UTC date and time in the format YYYY-MM-DDTHH:MM:SSZ. Any data before this date will not be replicated.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `source_okta_authorization_method_api_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_okta_authorization_method_api_token))
- `source_okta_authorization_method_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_okta_authorization_method_o_auth2_0))
- `source_okta_update_authorization_method_api_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_okta_update_authorization_method_api_token))
- `source_okta_update_authorization_method_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_okta_update_authorization_method_o_auth2_0))

<a id="nestedatt--configuration--credentials--source_okta_authorization_method_api_token"></a>
### Nested Schema for `configuration.credentials.source_okta_authorization_method_api_token`

Read-Only:

- `api_token` (String) An Okta token. See the <a href="https://docs.airbyte.com/integrations/sources/okta">docs</a> for instructions on how to generate it.
- `auth_type` (String) must be one of ["api_token"]


<a id="nestedatt--configuration--credentials--source_okta_authorization_method_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_okta_authorization_method_o_auth2_0`

Read-Only:

- `auth_type` (String) must be one of ["oauth2.0"]
- `client_id` (String) The Client ID of your OAuth application.
- `client_secret` (String) The Client Secret of your OAuth application.
- `refresh_token` (String) Refresh Token to obtain new Access Token, when it's expired.


<a id="nestedatt--configuration--credentials--source_okta_update_authorization_method_api_token"></a>
### Nested Schema for `configuration.credentials.source_okta_update_authorization_method_api_token`

Read-Only:

- `api_token` (String) An Okta token. See the <a href="https://docs.airbyte.com/integrations/sources/okta">docs</a> for instructions on how to generate it.
- `auth_type` (String) must be one of ["api_token"]


<a id="nestedatt--configuration--credentials--source_okta_update_authorization_method_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_okta_update_authorization_method_o_auth2_0`

Read-Only:

- `auth_type` (String) must be one of ["oauth2.0"]
- `client_id` (String) The Client ID of your OAuth application.
- `client_secret` (String) The Client Secret of your OAuth application.
- `refresh_token` (String) Refresh Token to obtain new Access Token, when it's expired.



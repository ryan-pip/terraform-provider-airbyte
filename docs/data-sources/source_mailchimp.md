---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_mailchimp Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceMailchimp DataSource
---

# airbyte_source_mailchimp (Data Source)

SourceMailchimp DataSource

## Example Usage

```terraform
data "airbyte_source_mailchimp" "my_source_mailchimp" {
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

- `campaign_id` (String)
- `credentials` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials))
- `source_type` (String) must be one of ["mailchimp"]

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `source_mailchimp_authentication_api_key` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_mailchimp_authentication_api_key))
- `source_mailchimp_authentication_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_mailchimp_authentication_o_auth2_0))
- `source_mailchimp_update_authentication_api_key` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_mailchimp_update_authentication_api_key))
- `source_mailchimp_update_authentication_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_mailchimp_update_authentication_o_auth2_0))

<a id="nestedatt--configuration--credentials--source_mailchimp_authentication_api_key"></a>
### Nested Schema for `configuration.credentials.source_mailchimp_authentication_api_key`

Read-Only:

- `apikey` (String) Mailchimp API Key. See the <a href="https://docs.airbyte.com/integrations/sources/mailchimp">docs</a> for information on how to generate this key.
- `auth_type` (String) must be one of ["apikey"]


<a id="nestedatt--configuration--credentials--source_mailchimp_authentication_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_mailchimp_authentication_o_auth2_0`

Read-Only:

- `access_token` (String) An access token generated using the above client ID and secret.
- `auth_type` (String) must be one of ["oauth2.0"]
- `client_id` (String) The Client ID of your OAuth application.
- `client_secret` (String) The Client Secret of your OAuth application.


<a id="nestedatt--configuration--credentials--source_mailchimp_update_authentication_api_key"></a>
### Nested Schema for `configuration.credentials.source_mailchimp_update_authentication_api_key`

Read-Only:

- `apikey` (String) Mailchimp API Key. See the <a href="https://docs.airbyte.com/integrations/sources/mailchimp">docs</a> for information on how to generate this key.
- `auth_type` (String) must be one of ["apikey"]


<a id="nestedatt--configuration--credentials--source_mailchimp_update_authentication_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_mailchimp_update_authentication_o_auth2_0`

Read-Only:

- `access_token` (String) An access token generated using the above client ID and secret.
- `auth_type` (String) must be one of ["oauth2.0"]
- `client_id` (String) The Client ID of your OAuth application.
- `client_secret` (String) The Client Secret of your OAuth application.



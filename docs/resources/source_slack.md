---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_slack Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSlack Resource
---

# airbyte_source_slack (Resource)

SourceSlack Resource

## Example Usage

```terraform
resource "airbyte_source_slack" "my_source_slack" {
  configuration = {
    channel_filter = [
      "...",
    ]
    credentials = {
      source_slack_authentication_mechanism_api_token = {
        api_token    = "...my_api_token..."
        option_title = "API Token Credentials"
      }
    }
    join_channels   = true
    lookback_window = 14
    source_type     = "slack"
    start_date      = "2017-01-25T00:00:00Z"
  }
  name         = "Eduardo Gottlieb"
  secret_id    = "...my_secret_id..."
  workspace_id = "43664a8f-0af8-4c69-9d73-2d9fbaf9476a"
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

- `join_channels` (Boolean) Whether to join all channels or to sync data only from channels the bot is already in.  If false, you'll need to manually add the bot to all the channels from which you'd like to sync messages.
- `lookback_window` (Number) How far into the past to look for messages in threads, default is 0 days
- `source_type` (String) must be one of ["slack"]
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.

Optional:

- `channel_filter` (List of String) A channel name list (without leading '#' char) which limit the channels from which you'd like to sync. Empty list means no filter.
- `credentials` (Attributes) Choose how to authenticate into Slack (see [below for nested schema](#nestedatt--configuration--credentials))

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `source_slack_authentication_mechanism_api_token` (Attributes) Choose how to authenticate into Slack (see [below for nested schema](#nestedatt--configuration--credentials--source_slack_authentication_mechanism_api_token))
- `source_slack_authentication_mechanism_sign_in_via_slack_o_auth` (Attributes) Choose how to authenticate into Slack (see [below for nested schema](#nestedatt--configuration--credentials--source_slack_authentication_mechanism_sign_in_via_slack_o_auth))
- `source_slack_update_authentication_mechanism_api_token` (Attributes) Choose how to authenticate into Slack (see [below for nested schema](#nestedatt--configuration--credentials--source_slack_update_authentication_mechanism_api_token))
- `source_slack_update_authentication_mechanism_sign_in_via_slack_o_auth` (Attributes) Choose how to authenticate into Slack (see [below for nested schema](#nestedatt--configuration--credentials--source_slack_update_authentication_mechanism_sign_in_via_slack_o_auth))

<a id="nestedatt--configuration--credentials--source_slack_authentication_mechanism_api_token"></a>
### Nested Schema for `configuration.credentials.source_slack_authentication_mechanism_api_token`

Required:

- `api_token` (String) A Slack bot token. See the <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> for instructions on how to generate it.
- `option_title` (String) must be one of ["API Token Credentials"]


<a id="nestedatt--configuration--credentials--source_slack_authentication_mechanism_sign_in_via_slack_o_auth"></a>
### Nested Schema for `configuration.credentials.source_slack_authentication_mechanism_sign_in_via_slack_o_auth`

Required:

- `access_token` (String) Slack access_token. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help generating the token.
- `client_id` (String) Slack client_id. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this id.
- `client_secret` (String) Slack client_secret. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this secret.
- `option_title` (String) must be one of ["Default OAuth2.0 authorization"]


<a id="nestedatt--configuration--credentials--source_slack_update_authentication_mechanism_api_token"></a>
### Nested Schema for `configuration.credentials.source_slack_update_authentication_mechanism_api_token`

Required:

- `api_token` (String) A Slack bot token. See the <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> for instructions on how to generate it.
- `option_title` (String) must be one of ["API Token Credentials"]


<a id="nestedatt--configuration--credentials--source_slack_update_authentication_mechanism_sign_in_via_slack_o_auth"></a>
### Nested Schema for `configuration.credentials.source_slack_update_authentication_mechanism_sign_in_via_slack_o_auth`

Required:

- `access_token` (String) Slack access_token. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help generating the token.
- `client_id` (String) Slack client_id. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this id.
- `client_secret` (String) Slack client_secret. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this secret.
- `option_title` (String) must be one of ["Default OAuth2.0 authorization"]



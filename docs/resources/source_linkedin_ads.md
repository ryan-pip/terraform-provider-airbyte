---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_linkedin_ads Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceLinkedinAds Resource
---

# airbyte_source_linkedin_ads (Resource)

SourceLinkedinAds Resource

## Example Usage

```terraform
resource "airbyte_source_linkedin_ads" "my_source_linkedinads" {
  configuration = {
    account_ids = [
      1,
    ]
    ad_analytics_reports = [
      {
        name             = "Mable Stroman"
        pivot_by         = "MEMBER_COMPANY_SIZE"
        time_granularity = "MONTHLY"
      },
    ]
    credentials = {
      source_linkedin_ads_authentication_access_token = {
        access_token = "...my_access_token..."
        auth_method  = "access_token"
      }
    }
    source_type = "linkedin-ads"
    start_date  = "2021-05-17"
  }
  name         = "Leigh Kuhic"
  secret_id    = "...my_secret_id..."
  workspace_id = "1cbe6d95-02f0-4ea9-b0b6-9f7ac2f72f88"
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

- `source_type` (String) must be one of ["linkedin-ads"]
- `start_date` (String) UTC date in the format 2020-09-17. Any data before this date will not be replicated.

Optional:

- `account_ids` (List of Number) Specify the account IDs separated by a space, to pull the data from. Leave empty, if you want to pull the data from all associated accounts. See the <a href="https://www.linkedin.com/help/linkedin/answer/a424270/find-linkedin-ads-account-details?lang=en">LinkedIn Ads docs</a> for more info.
- `ad_analytics_reports` (Attributes List) (see [below for nested schema](#nestedatt--configuration--ad_analytics_reports))
- `credentials` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials))

<a id="nestedatt--configuration--ad_analytics_reports"></a>
### Nested Schema for `configuration.ad_analytics_reports`

Required:

- `name` (String) The name for the report
- `pivot_by` (String) must be one of ["COMPANY", "ACCOUNT", "SHARE", "CAMPAIGN", "CREATIVE", "CAMPAIGN_GROUP", "CONVERSION", "CONVERSATION_NODE", "CONVERSATION_NODE_OPTION_INDEX", "SERVING_LOCATION", "CARD_INDEX", "MEMBER_COMPANY_SIZE", "MEMBER_INDUSTRY", "MEMBER_SENIORITY", "MEMBER_JOB_TITLE ", "MEMBER_JOB_FUNCTION ", "MEMBER_COUNTRY_V2 ", "MEMBER_REGION_V2", "MEMBER_COMPANY", "PLACEMENT_NAME", "IMPRESSION_DEVICE_TYPE"]
Select value from list to pivot by
- `time_granularity` (String) must be one of ["ALL", "DAILY", "MONTHLY", "YEARLY"]
Set time granularity for report: 
ALL - Results grouped into a single result across the entire time range of the report.
DAILY - Results grouped by day.
MONTHLY - Results grouped by month.
YEARLY - Results grouped by year.


<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `source_linkedin_ads_authentication_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_linkedin_ads_authentication_access_token))
- `source_linkedin_ads_authentication_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_linkedin_ads_authentication_o_auth2_0))
- `source_linkedin_ads_update_authentication_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_linkedin_ads_update_authentication_access_token))
- `source_linkedin_ads_update_authentication_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_linkedin_ads_update_authentication_o_auth2_0))

<a id="nestedatt--configuration--credentials--source_linkedin_ads_authentication_access_token"></a>
### Nested Schema for `configuration.credentials.source_linkedin_ads_authentication_access_token`

Required:

- `access_token` (String) The token value generated using the authentication code. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-ads#authentication">docs</a> to obtain yours.

Optional:

- `auth_method` (String) must be one of ["access_token"]


<a id="nestedatt--configuration--credentials--source_linkedin_ads_authentication_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_linkedin_ads_authentication_o_auth2_0`

Required:

- `client_id` (String) The client ID of the LinkedIn Ads developer application.
- `client_secret` (String) The client secret the LinkedIn Ads developer application.
- `refresh_token` (String) The key to refresh the expired access token.

Optional:

- `auth_method` (String) must be one of ["oAuth2.0"]


<a id="nestedatt--configuration--credentials--source_linkedin_ads_update_authentication_access_token"></a>
### Nested Schema for `configuration.credentials.source_linkedin_ads_update_authentication_access_token`

Required:

- `access_token` (String) The token value generated using the authentication code. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-ads#authentication">docs</a> to obtain yours.

Optional:

- `auth_method` (String) must be one of ["access_token"]


<a id="nestedatt--configuration--credentials--source_linkedin_ads_update_authentication_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_linkedin_ads_update_authentication_o_auth2_0`

Required:

- `client_id` (String) The client ID of the LinkedIn Ads developer application.
- `client_secret` (String) The client secret the LinkedIn Ads developer application.
- `refresh_token` (String) The key to refresh the expired access token.

Optional:

- `auth_method` (String) must be one of ["oAuth2.0"]



---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_google_pagespeed_insights Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGooglePagespeedInsights Resource
---

# airbyte_source_google_pagespeed_insights (Resource)

SourceGooglePagespeedInsights Resource

## Example Usage

```terraform
resource "airbyte_source_google_pagespeed_insights" "my_source_googlepagespeedinsights" {
  configuration = {
    api_key = "...my_api_key..."
    categories = [
      "pwa",
    ]
    source_type = "google-pagespeed-insights"
    strategies = [
      "desktop",
    ]
    urls = [
      "...",
    ]
  }
  name         = "Miss Terrell Satterfield"
  secret_id    = "...my_secret_id..."
  workspace_id = "6230f841-fb1b-4d23-bdb1-4db6be5a6859"
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

- `categories` (List of String) Defines which Lighthouse category to run. One or many of: "accessibility", "best-practices", "performance", "pwa", "seo".
- `source_type` (String) must be one of ["google-pagespeed-insights"]
- `strategies` (List of String) The analyses strategy to use. Either "desktop" or "mobile".
- `urls` (List of String) The URLs to retrieve pagespeed information from. The connector will attempt to sync PageSpeed reports for all the defined URLs. Format: https://(www.)url.domain

Optional:

- `api_key` (String) Google PageSpeed API Key. See <a href="https://developers.google.com/speed/docs/insights/v5/get-started#APIKey">here</a>. The key is optional - however the API is heavily rate limited when using without API Key. Creating and using the API key therefore is recommended. The key is case sensitive.



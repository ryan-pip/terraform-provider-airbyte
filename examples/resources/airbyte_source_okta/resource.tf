resource "airbyte_source_okta" "my_source_okta" {
  configuration = {
    credentials = {
      source_okta_authorization_method_api_token = {
        api_token = "...my_api_token..."
        auth_type = "api_token"
      }
    }
    domain      = "...my_domain..."
    source_type = "okta"
    start_date  = "2022-07-22T00:00:00Z"
  }
  name         = "Randolph Russel"
  secret_id    = "...my_secret_id..."
  workspace_id = "47b7684e-ff50-4126-971c-ffbd0eb74b84"
}
resource "airbyte_source_twilio" "my_source_twilio" {
  configuration = {
    account_sid     = "...my_account_sid..."
    auth_token      = "...my_auth_token..."
    lookback_window = 60
    source_type     = "twilio"
    start_date      = "2020-10-01T00:00:00Z"
  }
  name         = "Oliver Kautzer"
  secret_id    = "...my_secret_id..."
  workspace_id = "3b90a1b8-c95b-4e12-94b7-39f4fe77210d"
}
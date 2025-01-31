resource "airbyte_source_faker" "my_source_faker" {
  configuration = {
    always_updated    = false
    count             = 4
    parallelism       = 1
    records_per_slice = 4
    seed              = 4
    source_type       = "faker"
  }
  name         = "Beth McKenzie"
  secret_id    = "...my_secret_id..."
  workspace_id = "1ec00221-b335-4d89-acb3-ecfda8d0c549"
}
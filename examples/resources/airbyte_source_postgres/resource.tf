resource "airbyte_source_postgres" "my_source_postgres" {
  configuration = {
    database        = "...my_database..."
    host            = "...my_host..."
    jdbc_url_params = "...my_jdbc_url_params..."
    password        = "...my_password..."
    port            = 5432
    replication_method = {
      source_postgres_replication_method_logical_replication_cdc_ = {
        initial_waiting_seconds = 4
        lsn_commit_behaviour    = "While reading Data"
        method                  = "CDC"
        plugin                  = "pgoutput"
        publication             = "...my_publication..."
        queue_size              = 4
        replication_slot        = "...my_replication_slot..."
      }
    }
    schemas = [
      "...",
    ]
    source_type = "postgres"
    ssl_mode = {
      source_postgres_ssl_modes_allow = {
        mode = "allow"
      }
    }
    tunnel_method = {
      source_postgres_ssh_tunnel_method_no_tunnel = {
        tunnel_method = "NO_TUNNEL"
      }
    }
    username = "Kendrick52"
  }
  name         = "Vivian Dietrich"
  secret_id    = "...my_secret_id..."
  workspace_id = "10006bef-4921-4ec2-853b-749366ac8ee0"
}
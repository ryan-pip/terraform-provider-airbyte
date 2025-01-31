---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_file_secure Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceFileSecure DataSource
---

# airbyte_source_file_secure (Data Source)

SourceFileSecure DataSource

## Example Usage

```terraform
data "airbyte_source_file_secure" "my_source_filesecure" {
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

- `dataset_name` (String) The Name of the final table to replicate this file into (should include letters, numbers dash and underscores only).
- `format` (String) must be one of ["csv", "json", "jsonl", "excel", "excel_binary", "feather", "parquet", "yaml"]
The Format of the file which should be replicated (Warning: some formats may be experimental, please refer to the docs).
- `provider` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider))
- `reader_options` (String) This should be a string in JSON format. It depends on the chosen file format to provide additional options and tune its behavior.
- `source_type` (String) must be one of ["file-secure"]
- `url` (String) The URL path to access the file which should be replicated.

<a id="nestedatt--configuration--provider"></a>
### Nested Schema for `configuration.provider`

Read-Only:

- `source_file_secure_storage_provider_az_blob_azure_blob_storage` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_storage_provider_az_blob_azure_blob_storage))
- `source_file_secure_storage_provider_gcs_google_cloud_storage` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_storage_provider_gcs_google_cloud_storage))
- `source_file_secure_storage_provider_https_public_web` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_storage_provider_https_public_web))
- `source_file_secure_storage_provider_s3_amazon_web_services` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_storage_provider_s3_amazon_web_services))
- `source_file_secure_storage_provider_scp_secure_copy_protocol` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_storage_provider_scp_secure_copy_protocol))
- `source_file_secure_storage_provider_sftp_secure_file_transfer_protocol` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_storage_provider_sftp_secure_file_transfer_protocol))
- `source_file_secure_storage_provider_ssh_secure_shell` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_storage_provider_ssh_secure_shell))
- `source_file_secure_update_storage_provider_az_blob_azure_blob_storage` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_update_storage_provider_az_blob_azure_blob_storage))
- `source_file_secure_update_storage_provider_gcs_google_cloud_storage` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_update_storage_provider_gcs_google_cloud_storage))
- `source_file_secure_update_storage_provider_https_public_web` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_update_storage_provider_https_public_web))
- `source_file_secure_update_storage_provider_s3_amazon_web_services` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_update_storage_provider_s3_amazon_web_services))
- `source_file_secure_update_storage_provider_scp_secure_copy_protocol` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_update_storage_provider_scp_secure_copy_protocol))
- `source_file_secure_update_storage_provider_sftp_secure_file_transfer_protocol` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_update_storage_provider_sftp_secure_file_transfer_protocol))
- `source_file_secure_update_storage_provider_ssh_secure_shell` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--source_file_secure_update_storage_provider_ssh_secure_shell))

<a id="nestedatt--configuration--provider--source_file_secure_storage_provider_az_blob_azure_blob_storage"></a>
### Nested Schema for `configuration.provider.source_file_secure_storage_provider_az_blob_azure_blob_storage`

Read-Only:

- `sas_token` (String) To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a SAS (Shared Access Signature) token. If accessing publicly available data, this field is not necessary.
- `shared_key` (String) To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a storage account shared key (aka account key or access key). If accessing publicly available data, this field is not necessary.
- `storage` (String) must be one of ["AzBlob"]
- `storage_account` (String) The globally unique name of the storage account that the desired blob sits within. See <a href="https://docs.microsoft.com/en-us/azure/storage/common/storage-account-overview" target="_blank">here</a> for more details.


<a id="nestedatt--configuration--provider--source_file_secure_storage_provider_gcs_google_cloud_storage"></a>
### Nested Schema for `configuration.provider.source_file_secure_storage_provider_gcs_google_cloud_storage`

Read-Only:

- `service_account_json` (String) In order to access private Buckets stored on Google Cloud, this connector would need a service account json credentials with the proper permissions as described <a href="https://cloud.google.com/iam/docs/service-accounts" target="_blank">here</a>. Please generate the credentials.json file and copy/paste its content to this field (expecting JSON formats). If accessing publicly available data, this field is not necessary.
- `storage` (String) must be one of ["GCS"]


<a id="nestedatt--configuration--provider--source_file_secure_storage_provider_https_public_web"></a>
### Nested Schema for `configuration.provider.source_file_secure_storage_provider_https_public_web`

Read-Only:

- `storage` (String) must be one of ["HTTPS"]
- `user_agent` (Boolean) Add User-Agent to request


<a id="nestedatt--configuration--provider--source_file_secure_storage_provider_s3_amazon_web_services"></a>
### Nested Schema for `configuration.provider.source_file_secure_storage_provider_s3_amazon_web_services`

Read-Only:

- `aws_access_key_id` (String) In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
- `aws_secret_access_key` (String) In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
- `storage` (String) must be one of ["S3"]


<a id="nestedatt--configuration--provider--source_file_secure_storage_provider_scp_secure_copy_protocol"></a>
### Nested Schema for `configuration.provider.source_file_secure_storage_provider_scp_secure_copy_protocol`

Read-Only:

- `host` (String)
- `password` (String)
- `port` (String)
- `storage` (String) must be one of ["SCP"]
- `user` (String)


<a id="nestedatt--configuration--provider--source_file_secure_storage_provider_sftp_secure_file_transfer_protocol"></a>
### Nested Schema for `configuration.provider.source_file_secure_storage_provider_sftp_secure_file_transfer_protocol`

Read-Only:

- `host` (String)
- `password` (String)
- `port` (String)
- `storage` (String) must be one of ["SFTP"]
- `user` (String)


<a id="nestedatt--configuration--provider--source_file_secure_storage_provider_ssh_secure_shell"></a>
### Nested Schema for `configuration.provider.source_file_secure_storage_provider_ssh_secure_shell`

Read-Only:

- `host` (String)
- `password` (String)
- `port` (String)
- `storage` (String) must be one of ["SSH"]
- `user` (String)


<a id="nestedatt--configuration--provider--source_file_secure_update_storage_provider_az_blob_azure_blob_storage"></a>
### Nested Schema for `configuration.provider.source_file_secure_update_storage_provider_az_blob_azure_blob_storage`

Read-Only:

- `sas_token` (String) To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a SAS (Shared Access Signature) token. If accessing publicly available data, this field is not necessary.
- `shared_key` (String) To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a storage account shared key (aka account key or access key). If accessing publicly available data, this field is not necessary.
- `storage` (String) must be one of ["AzBlob"]
- `storage_account` (String) The globally unique name of the storage account that the desired blob sits within. See <a href="https://docs.microsoft.com/en-us/azure/storage/common/storage-account-overview" target="_blank">here</a> for more details.


<a id="nestedatt--configuration--provider--source_file_secure_update_storage_provider_gcs_google_cloud_storage"></a>
### Nested Schema for `configuration.provider.source_file_secure_update_storage_provider_gcs_google_cloud_storage`

Read-Only:

- `service_account_json` (String) In order to access private Buckets stored on Google Cloud, this connector would need a service account json credentials with the proper permissions as described <a href="https://cloud.google.com/iam/docs/service-accounts" target="_blank">here</a>. Please generate the credentials.json file and copy/paste its content to this field (expecting JSON formats). If accessing publicly available data, this field is not necessary.
- `storage` (String) must be one of ["GCS"]


<a id="nestedatt--configuration--provider--source_file_secure_update_storage_provider_https_public_web"></a>
### Nested Schema for `configuration.provider.source_file_secure_update_storage_provider_https_public_web`

Read-Only:

- `storage` (String) must be one of ["HTTPS"]
- `user_agent` (Boolean) Add User-Agent to request


<a id="nestedatt--configuration--provider--source_file_secure_update_storage_provider_s3_amazon_web_services"></a>
### Nested Schema for `configuration.provider.source_file_secure_update_storage_provider_s3_amazon_web_services`

Read-Only:

- `aws_access_key_id` (String) In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
- `aws_secret_access_key` (String) In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
- `storage` (String) must be one of ["S3"]


<a id="nestedatt--configuration--provider--source_file_secure_update_storage_provider_scp_secure_copy_protocol"></a>
### Nested Schema for `configuration.provider.source_file_secure_update_storage_provider_scp_secure_copy_protocol`

Read-Only:

- `host` (String)
- `password` (String)
- `port` (String)
- `storage` (String) must be one of ["SCP"]
- `user` (String)


<a id="nestedatt--configuration--provider--source_file_secure_update_storage_provider_sftp_secure_file_transfer_protocol"></a>
### Nested Schema for `configuration.provider.source_file_secure_update_storage_provider_sftp_secure_file_transfer_protocol`

Read-Only:

- `host` (String)
- `password` (String)
- `port` (String)
- `storage` (String) must be one of ["SFTP"]
- `user` (String)


<a id="nestedatt--configuration--provider--source_file_secure_update_storage_provider_ssh_secure_shell"></a>
### Nested Schema for `configuration.provider.source_file_secure_update_storage_provider_ssh_secure_shell`

Read-Only:

- `host` (String)
- `password` (String)
- `port` (String)
- `storage` (String) must be one of ["SSH"]
- `user` (String)



// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk"
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"github.com/ryan-pip/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceMssqlDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceMssqlDataSource{}

func NewSourceMssqlDataSource() datasource.DataSource {
	return &SourceMssqlDataSource{}
}

// SourceMssqlDataSource is the data source implementation.
type SourceMssqlDataSource struct {
	client *sdk.SDK
}

// SourceMssqlDataSourceModel describes the data model.
type SourceMssqlDataSourceModel struct {
	Configuration SourceMssql  `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceMssqlDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_mssql"
}

// Schema defines the schema for the data source.
func (r *SourceMssqlDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceMssql DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the database.`,
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Description: `The hostname of the database.`,
					},
					"jdbc_url_params": schema.StringAttribute{
						Computed:    true,
						Description: `Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).`,
					},
					"password": schema.StringAttribute{
						Computed:    true,
						Description: `The password associated with the username.`,
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Description: `The port of the database.`,
					},
					"replication_method": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_mssql_replication_method_logical_replication_cdc": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"data_to_sync": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Existing and New",
												"New Changes Only",
											),
										},
										MarkdownDescription: `must be one of ["Existing and New", "New Changes Only"]` + "\n" +
											`What data should be synced under the CDC. "Existing and New" will read existing data as a snapshot, and sync new changes through CDC. "New Changes Only" will skip the initial snapshot, and only sync new changes through CDC.`,
									},
									"initial_waiting_seconds": schema.Int64Attribute{
										Computed:    true,
										Description: `The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">initial waiting time</a>.`,
									},
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CDC",
											),
										},
										Description: `must be one of ["CDC"]`,
									},
									"snapshot_isolation": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Snapshot",
												"Read Committed",
											),
										},
										MarkdownDescription: `must be one of ["Snapshot", "Read Committed"]` + "\n" +
											`Existing data in the database are synced through an initial snapshot. This parameter controls the isolation level that will be used during the initial snapshotting. If you choose the "Snapshot" level, you must enable the <a href="https://docs.microsoft.com/en-us/dotnet/framework/data/adonet/sql/snapshot-isolation-in-sql-server">snapshot isolation mode</a> on the database.`,
									},
								},
								Description: `CDC uses {TBC} to detect inserts, updates, and deletes. This needs to be configured on the source database itself.`,
							},
							"source_mssql_replication_method_standard": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"STANDARD",
											),
										},
										Description: `must be one of ["STANDARD"]`,
									},
								},
								Description: `Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
							"source_mssql_update_replication_method_logical_replication_cdc": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"data_to_sync": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Existing and New",
												"New Changes Only",
											),
										},
										MarkdownDescription: `must be one of ["Existing and New", "New Changes Only"]` + "\n" +
											`What data should be synced under the CDC. "Existing and New" will read existing data as a snapshot, and sync new changes through CDC. "New Changes Only" will skip the initial snapshot, and only sync new changes through CDC.`,
									},
									"initial_waiting_seconds": schema.Int64Attribute{
										Computed:    true,
										Description: `The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">initial waiting time</a>.`,
									},
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CDC",
											),
										},
										Description: `must be one of ["CDC"]`,
									},
									"snapshot_isolation": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Snapshot",
												"Read Committed",
											),
										},
										MarkdownDescription: `must be one of ["Snapshot", "Read Committed"]` + "\n" +
											`Existing data in the database are synced through an initial snapshot. This parameter controls the isolation level that will be used during the initial snapshotting. If you choose the "Snapshot" level, you must enable the <a href="https://docs.microsoft.com/en-us/dotnet/framework/data/adonet/sql/snapshot-isolation-in-sql-server">snapshot isolation mode</a> on the database.`,
									},
								},
								Description: `CDC uses {TBC} to detect inserts, updates, and deletes. This needs to be configured on the source database itself.`,
							},
							"source_mssql_update_replication_method_standard": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"STANDARD",
											),
										},
										Description: `must be one of ["STANDARD"]`,
									},
								},
								Description: `Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `The replication method used for extracting data from the database. STANDARD replication requires no setup on the DB side but will not be able to represent deletions incrementally. CDC uses {TBC} to detect inserts, updates, and deletes. This needs to be configured on the source database itself.`,
					},
					"schemas": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `The list of schemas to sync from. Defaults to user. Case sensitive.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"mssql",
							),
						},
						Description: `must be one of ["mssql"]`,
					},
					"ssl_method": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_mssql_ssl_method_encrypted_trust_server_certificate": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"ssl_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"encrypted_trust_server_certificate",
											),
										},
										Description: `must be one of ["encrypted_trust_server_certificate"]`,
									},
								},
								Description: `Use the certificate provided by the server without verification. (For testing purposes only!)`,
							},
							"source_mssql_ssl_method_encrypted_verify_certificate": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"host_name_in_certificate": schema.StringAttribute{
										Computed:    true,
										Description: `Specifies the host name of the server. The value of this property must match the subject property of the certificate.`,
									},
									"ssl_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"encrypted_verify_certificate",
											),
										},
										Description: `must be one of ["encrypted_verify_certificate"]`,
									},
								},
								Description: `Verify and use the certificate provided by the server.`,
							},
							"source_mssql_update_ssl_method_encrypted_trust_server_certificate": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"ssl_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"encrypted_trust_server_certificate",
											),
										},
										Description: `must be one of ["encrypted_trust_server_certificate"]`,
									},
								},
								Description: `Use the certificate provided by the server without verification. (For testing purposes only!)`,
							},
							"source_mssql_update_ssl_method_encrypted_verify_certificate": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"host_name_in_certificate": schema.StringAttribute{
										Computed:    true,
										Description: `Specifies the host name of the server. The value of this property must match the subject property of the certificate.`,
									},
									"ssl_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"encrypted_verify_certificate",
											),
										},
										Description: `must be one of ["encrypted_verify_certificate"]`,
									},
								},
								Description: `Verify and use the certificate provided by the server.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `The encryption method which is used when communicating with the database.`,
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_mssql_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"NO_TUNNEL",
											),
										},
										MarkdownDescription: `must be one of ["NO_TUNNEL"]` + "\n" +
											`No ssh tunnel needed to connect to database`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_mssql_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Computed:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_PASSWORD_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level username for logging into the jump server host`,
									},
									"tunnel_user_password": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_mssql_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )`,
									},
									"tunnel_host": schema.StringAttribute{
										Computed:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_KEY_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level username for logging into the jump server host.`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_mssql_update_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"NO_TUNNEL",
											),
										},
										MarkdownDescription: `must be one of ["NO_TUNNEL"]` + "\n" +
											`No ssh tunnel needed to connect to database`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_mssql_update_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Computed:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_PASSWORD_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level username for logging into the jump server host`,
									},
									"tunnel_user_password": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_mssql_update_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )`,
									},
									"tunnel_host": schema.StringAttribute{
										Computed:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_KEY_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level username for logging into the jump server host.`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
					},
					"username": schema.StringAttribute{
						Computed:    true,
						Description: `The username which is used to access the database.`,
					},
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"secret_id": schema.StringAttribute{
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *SourceMssqlDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceMssqlDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceMssqlDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.GetSourceMssqlRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceMssql(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

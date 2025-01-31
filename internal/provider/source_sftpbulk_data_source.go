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
var _ datasource.DataSource = &SourceSftpBulkDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceSftpBulkDataSource{}

func NewSourceSftpBulkDataSource() datasource.DataSource {
	return &SourceSftpBulkDataSource{}
}

// SourceSftpBulkDataSource is the data source implementation.
type SourceSftpBulkDataSource struct {
	client *sdk.SDK
}

// SourceSftpBulkDataSourceModel describes the data model.
type SourceSftpBulkDataSourceModel struct {
	Configuration SourceSftpBulk `tfsdk:"configuration"`
	Name          types.String   `tfsdk:"name"`
	SecretID      types.String   `tfsdk:"secret_id"`
	SourceID      types.String   `tfsdk:"source_id"`
	WorkspaceID   types.String   `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceSftpBulkDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_sftp_bulk"
}

// Schema defines the schema for the data source.
func (r *SourceSftpBulkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceSftpBulk DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"file_most_recent": schema.BoolAttribute{
						Computed:    true,
						Description: `Sync only the most recent file for the configured folder path and file pattern`,
					},
					"file_pattern": schema.StringAttribute{
						Computed:    true,
						Description: `The regular expression to specify files for sync in a chosen Folder Path`,
					},
					"file_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"csv",
								"json",
							),
						},
						MarkdownDescription: `must be one of ["csv", "json"]` + "\n" +
							`The file type you want to sync. Currently only 'csv' and 'json' files are supported.`,
					},
					"folder_path": schema.StringAttribute{
						Computed:    true,
						Description: `The directory to search files for sync`,
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Description: `The server host address`,
					},
					"password": schema.StringAttribute{
						Computed:    true,
						Description: `OS-level password for logging into the jump server host`,
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Description: `The server port`,
					},
					"private_key": schema.StringAttribute{
						Computed:    true,
						Description: `The private key`,
					},
					"separator": schema.StringAttribute{
						Computed:    true,
						Description: `The separator used in the CSV files. Define None if you want to use the Sniffer functionality`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"sftp-bulk",
							),
						},
						Description: `must be one of ["sftp-bulk"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `The date from which you'd like to replicate data for all incremental streams, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.`,
					},
					"stream_name": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the stream or table you want to create`,
					},
					"username": schema.StringAttribute{
						Computed:    true,
						Description: `The server user`,
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

func (r *SourceSftpBulkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceSftpBulkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceSftpBulkDataSourceModel
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
	request := operations.GetSourceSftpBulkRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceSftpBulk(ctx, request)
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

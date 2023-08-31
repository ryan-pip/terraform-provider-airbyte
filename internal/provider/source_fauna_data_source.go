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
var _ datasource.DataSource = &SourceFaunaDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceFaunaDataSource{}

func NewSourceFaunaDataSource() datasource.DataSource {
	return &SourceFaunaDataSource{}
}

// SourceFaunaDataSource is the data source implementation.
type SourceFaunaDataSource struct {
	client *sdk.SDK
}

// SourceFaunaDataSourceModel describes the data model.
type SourceFaunaDataSourceModel struct {
	Configuration SourceFauna  `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceFaunaDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_fauna"
}

// Schema defines the schema for the data source.
func (r *SourceFaunaDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceFauna DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"collection": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"deletions": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"source_fauna_collection_deletion_mode_disabled": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"deletion_mode": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"ignore",
													),
												},
												Description: `must be one of ["ignore"]`,
											},
										},
										MarkdownDescription: `<b>This only applies to incremental syncs.</b> <br>` + "\n" +
											`Enabling deletion mode informs your destination of deleted documents.<br>` + "\n" +
											`Disabled - Leave this feature disabled, and ignore deleted documents.<br>` + "\n" +
											`Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.`,
									},
									"source_fauna_collection_deletion_mode_enabled": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"column": schema.StringAttribute{
												Computed:    true,
												Description: `Name of the "deleted at" column.`,
											},
											"deletion_mode": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"deleted_field",
													),
												},
												Description: `must be one of ["deleted_field"]`,
											},
										},
										MarkdownDescription: `<b>This only applies to incremental syncs.</b> <br>` + "\n" +
											`Enabling deletion mode informs your destination of deleted documents.<br>` + "\n" +
											`Disabled - Leave this feature disabled, and ignore deleted documents.<br>` + "\n" +
											`Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.`,
									},
									"source_fauna_update_collection_deletion_mode_disabled": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"deletion_mode": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"ignore",
													),
												},
												Description: `must be one of ["ignore"]`,
											},
										},
										MarkdownDescription: `<b>This only applies to incremental syncs.</b> <br>` + "\n" +
											`Enabling deletion mode informs your destination of deleted documents.<br>` + "\n" +
											`Disabled - Leave this feature disabled, and ignore deleted documents.<br>` + "\n" +
											`Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.`,
									},
									"source_fauna_update_collection_deletion_mode_enabled": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"column": schema.StringAttribute{
												Computed:    true,
												Description: `Name of the "deleted at" column.`,
											},
											"deletion_mode": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"deleted_field",
													),
												},
												Description: `must be one of ["deleted_field"]`,
											},
										},
										MarkdownDescription: `<b>This only applies to incremental syncs.</b> <br>` + "\n" +
											`Enabling deletion mode informs your destination of deleted documents.<br>` + "\n" +
											`Disabled - Leave this feature disabled, and ignore deleted documents.<br>` + "\n" +
											`Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.`,
									},
								},
								Validators: []validator.Object{
									validators.ExactlyOneChild(),
								},
								MarkdownDescription: `<b>This only applies to incremental syncs.</b> <br>` + "\n" +
									`Enabling deletion mode informs your destination of deleted documents.<br>` + "\n" +
									`Disabled - Leave this feature disabled, and ignore deleted documents.<br>` + "\n" +
									`Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.`,
							},
							"page_size": schema.Int64Attribute{
								Computed: true,
								MarkdownDescription: `The page size used when reading documents from the database. The larger the page size, the faster the connector processes documents. However, if a page is too large, the connector may fail. <br>` + "\n" +
									`Choose your page size based on how large the documents are. <br>` + "\n" +
									`See <a href="https://docs.fauna.com/fauna/current/learn/understanding/types#page">the docs</a>.`,
							},
						},
						Description: `Settings for the Fauna Collection.`,
					},
					"domain": schema.StringAttribute{
						Computed:    true,
						Description: `Domain of Fauna to query. Defaults db.fauna.com. See <a href=https://docs.fauna.com/fauna/current/learn/understanding/region_groups#how-to-use-region-groups>the docs</a>.`,
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Description: `Endpoint port.`,
					},
					"scheme": schema.StringAttribute{
						Computed:    true,
						Description: `URL scheme.`,
					},
					"secret": schema.StringAttribute{
						Computed:    true,
						Description: `Fauna secret, used when authenticating with the database.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"fauna",
							),
						},
						Description: `must be one of ["fauna"]`,
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

func (r *SourceFaunaDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceFaunaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceFaunaDataSourceModel
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
	request := operations.GetSourceFaunaRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceFauna(ctx, request)
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

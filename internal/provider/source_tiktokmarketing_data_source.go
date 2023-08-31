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
var _ datasource.DataSource = &SourceTiktokMarketingDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceTiktokMarketingDataSource{}

func NewSourceTiktokMarketingDataSource() datasource.DataSource {
	return &SourceTiktokMarketingDataSource{}
}

// SourceTiktokMarketingDataSource is the data source implementation.
type SourceTiktokMarketingDataSource struct {
	client *sdk.SDK
}

// SourceTiktokMarketingDataSourceModel describes the data model.
type SourceTiktokMarketingDataSourceModel struct {
	Configuration SourceTiktokMarketing `tfsdk:"configuration"`
	Name          types.String          `tfsdk:"name"`
	SecretID      types.String          `tfsdk:"secret_id"`
	SourceID      types.String          `tfsdk:"source_id"`
	WorkspaceID   types.String          `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceTiktokMarketingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_tiktok_marketing"
}

// Schema defines the schema for the data source.
func (r *SourceTiktokMarketingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceTiktokMarketing DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"attribution_window": schema.Int64Attribute{
						Computed:    true,
						Description: `The attribution window in days.`,
					},
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_tiktok_marketing_authentication_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Long-term Authorized Access Token.`,
									},
									"advertiser_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Advertiser ID to filter reports and streams. Let this empty to retrieve all.`,
									},
									"app_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Developer Application App ID.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
										Description: `must be one of ["oauth2.0"]`,
									},
									"secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Developer Application Secret.`,
									},
								},
								Description: `Authentication method`,
							},
							"source_tiktok_marketing_authentication_method_sandbox_access_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `The long-term authorized access token.`,
									},
									"advertiser_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Advertiser ID which generated for the developer's Sandbox application.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"sandbox_access_token",
											),
										},
										Description: `must be one of ["sandbox_access_token"]`,
									},
								},
								Description: `Authentication method`,
							},
							"source_tiktok_marketing_update_authentication_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Long-term Authorized Access Token.`,
									},
									"advertiser_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Advertiser ID to filter reports and streams. Let this empty to retrieve all.`,
									},
									"app_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Developer Application App ID.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
										Description: `must be one of ["oauth2.0"]`,
									},
									"secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Developer Application Secret.`,
									},
								},
								Description: `Authentication method`,
							},
							"source_tiktok_marketing_update_authentication_method_sandbox_access_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `The long-term authorized access token.`,
									},
									"advertiser_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Advertiser ID which generated for the developer's Sandbox application.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"sandbox_access_token",
											),
										},
										Description: `must be one of ["sandbox_access_token"]`,
									},
								},
								Description: `Authentication method`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Authentication method`,
					},
					"end_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `The date until which you'd like to replicate data for all incremental streams, in the format YYYY-MM-DD. All data generated between start_date and this date will be replicated. Not setting this option will result in always syncing the data till the current date.`,
					},
					"include_deleted": schema.BoolAttribute{
						Computed:    true,
						Description: `Set to active if you want to include deleted data in reports.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"tiktok-marketing",
							),
						},
						Description: `must be one of ["tiktok-marketing"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `The Start Date in format: YYYY-MM-DD. Any data before this date will not be replicated. If this parameter is not set, all data will be replicated.`,
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

func (r *SourceTiktokMarketingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceTiktokMarketingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceTiktokMarketingDataSourceModel
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
	request := operations.GetSourceTiktokMarketingRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceTiktokMarketing(ctx, request)
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

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
var _ datasource.DataSource = &SourceSlackDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceSlackDataSource{}

func NewSourceSlackDataSource() datasource.DataSource {
	return &SourceSlackDataSource{}
}

// SourceSlackDataSource is the data source implementation.
type SourceSlackDataSource struct {
	client *sdk.SDK
}

// SourceSlackDataSourceModel describes the data model.
type SourceSlackDataSourceModel struct {
	Configuration SourceSlack  `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceSlackDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_slack"
}

// Schema defines the schema for the data source.
func (r *SourceSlackDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceSlack DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"channel_filter": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `A channel name list (without leading '#' char) which limit the channels from which you'd like to sync. Empty list means no filter.`,
					},
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_slack_authentication_mechanism_api_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"api_token": schema.StringAttribute{
										Computed:    true,
										Description: `A Slack bot token. See the <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> for instructions on how to generate it.`,
									},
									"option_title": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"API Token Credentials",
											),
										},
										Description: `must be one of ["API Token Credentials"]`,
									},
								},
								Description: `Choose how to authenticate into Slack`,
							},
							"source_slack_authentication_mechanism_sign_in_via_slack_o_auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Slack access_token. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help generating the token.`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `Slack client_id. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this id.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `Slack client_secret. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this secret.`,
									},
									"option_title": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Default OAuth2.0 authorization",
											),
										},
										Description: `must be one of ["Default OAuth2.0 authorization"]`,
									},
								},
								Description: `Choose how to authenticate into Slack`,
							},
							"source_slack_update_authentication_mechanism_api_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"api_token": schema.StringAttribute{
										Computed:    true,
										Description: `A Slack bot token. See the <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> for instructions on how to generate it.`,
									},
									"option_title": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"API Token Credentials",
											),
										},
										Description: `must be one of ["API Token Credentials"]`,
									},
								},
								Description: `Choose how to authenticate into Slack`,
							},
							"source_slack_update_authentication_mechanism_sign_in_via_slack_o_auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Slack access_token. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help generating the token.`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `Slack client_id. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this id.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `Slack client_secret. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this secret.`,
									},
									"option_title": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Default OAuth2.0 authorization",
											),
										},
										Description: `must be one of ["Default OAuth2.0 authorization"]`,
									},
								},
								Description: `Choose how to authenticate into Slack`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Choose how to authenticate into Slack`,
					},
					"join_channels": schema.BoolAttribute{
						Computed:    true,
						Description: `Whether to join all channels or to sync data only from channels the bot is already in.  If false, you'll need to manually add the bot to all the channels from which you'd like to sync messages. `,
					},
					"lookback_window": schema.Int64Attribute{
						Computed:    true,
						Description: `How far into the past to look for messages in threads, default is 0 days`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"slack",
							),
						},
						Description: `must be one of ["slack"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.`,
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

func (r *SourceSlackDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceSlackDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceSlackDataSourceModel
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
	request := operations.GetSourceSlackRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceSlack(ctx, request)
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

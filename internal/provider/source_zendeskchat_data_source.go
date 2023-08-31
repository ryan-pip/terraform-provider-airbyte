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
var _ datasource.DataSource = &SourceZendeskChatDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceZendeskChatDataSource{}

func NewSourceZendeskChatDataSource() datasource.DataSource {
	return &SourceZendeskChatDataSource{}
}

// SourceZendeskChatDataSource is the data source implementation.
type SourceZendeskChatDataSource struct {
	client *sdk.SDK
}

// SourceZendeskChatDataSourceModel describes the data model.
type SourceZendeskChatDataSourceModel struct {
	Configuration SourceZendeskChat `tfsdk:"configuration"`
	Name          types.String      `tfsdk:"name"`
	SecretID      types.String      `tfsdk:"secret_id"`
	SourceID      types.String      `tfsdk:"source_id"`
	WorkspaceID   types.String      `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceZendeskChatDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_zendesk_chat"
}

// Schema defines the schema for the data source.
func (r *SourceZendeskChatDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceZendeskChat DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_zendesk_chat_authorization_method_access_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `The Access Token to make authenticated requests.`,
									},
									"credentials": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"access_token",
											),
										},
										Description: `must be one of ["access_token"]`,
									},
								},
							},
							"source_zendesk_chat_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Access Token for making authenticated requests.`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of your OAuth application`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Client Secret of your OAuth application.`,
									},
									"credentials": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
										Description: `must be one of ["oauth2.0"]`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Refresh Token to obtain new Access Token, when it's expired.`,
									},
								},
							},
							"source_zendesk_chat_update_authorization_method_access_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `The Access Token to make authenticated requests.`,
									},
									"credentials": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"access_token",
											),
										},
										Description: `must be one of ["access_token"]`,
									},
								},
							},
							"source_zendesk_chat_update_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Access Token for making authenticated requests.`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of your OAuth application`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Client Secret of your OAuth application.`,
									},
									"credentials": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
										Description: `must be one of ["oauth2.0"]`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Refresh Token to obtain new Access Token, when it's expired.`,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"zendesk-chat",
							),
						},
						Description: `must be one of ["zendesk-chat"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `The date from which you'd like to replicate data for Zendesk Chat API, in the format YYYY-MM-DDT00:00:00Z.`,
					},
					"subdomain": schema.StringAttribute{
						Computed:    true,
						Description: `Required if you access Zendesk Chat from a Zendesk Support subdomain.`,
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

func (r *SourceZendeskChatDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceZendeskChatDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceZendeskChatDataSourceModel
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
	request := operations.GetSourceZendeskChatRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceZendeskChat(ctx, request)
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

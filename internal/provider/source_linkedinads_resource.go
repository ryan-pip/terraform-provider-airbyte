// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk"
	"context"
	"fmt"

	speakeasy_stringplanmodifier "github.com/ryan-pip/terraform-provider-airbyte/internal/planmodifiers/stringplanmodifier"
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/operations"
	"github.com/ryan-pip/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceLinkedinAdsResource{}
var _ resource.ResourceWithImportState = &SourceLinkedinAdsResource{}

func NewSourceLinkedinAdsResource() resource.Resource {
	return &SourceLinkedinAdsResource{}
}

// SourceLinkedinAdsResource defines the resource implementation.
type SourceLinkedinAdsResource struct {
	client *sdk.SDK
}

// SourceLinkedinAdsResourceModel describes the resource data model.
type SourceLinkedinAdsResourceModel struct {
	Configuration SourceLinkedinAds `tfsdk:"configuration"`
	Name          types.String      `tfsdk:"name"`
	SecretID      types.String      `tfsdk:"secret_id"`
	SourceID      types.String      `tfsdk:"source_id"`
	SourceType    types.String      `tfsdk:"source_type"`
	WorkspaceID   types.String      `tfsdk:"workspace_id"`
}

func (r *SourceLinkedinAdsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_linkedin_ads"
}

func (r *SourceLinkedinAdsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceLinkedinAds Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"account_ids": schema.ListAttribute{
						Optional:    true,
						ElementType: types.Int64Type,
						Description: `Specify the account IDs separated by a space, to pull the data from. Leave empty, if you want to pull the data from all associated accounts. See the <a href="https://www.linkedin.com/help/linkedin/answer/a424270/find-linkedin-ads-account-details?lang=en">LinkedIn Ads docs</a> for more info.`,
					},
					"ad_analytics_reports": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required:    true,
									Description: `The name for the report`,
								},
								"pivot_by": schema.StringAttribute{
									Required: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"COMPANY",
											"ACCOUNT",
											"SHARE",
											"CAMPAIGN",
											"CREATIVE",
											"CAMPAIGN_GROUP",
											"CONVERSION",
											"CONVERSATION_NODE",
											"CONVERSATION_NODE_OPTION_INDEX",
											"SERVING_LOCATION",
											"CARD_INDEX",
											"MEMBER_COMPANY_SIZE",
											"MEMBER_INDUSTRY",
											"MEMBER_SENIORITY",
											"MEMBER_JOB_TITLE ",
											"MEMBER_JOB_FUNCTION ",
											"MEMBER_COUNTRY_V2 ",
											"MEMBER_REGION_V2",
											"MEMBER_COMPANY",
											"PLACEMENT_NAME",
											"IMPRESSION_DEVICE_TYPE",
										),
									},
									MarkdownDescription: `must be one of ["COMPANY", "ACCOUNT", "SHARE", "CAMPAIGN", "CREATIVE", "CAMPAIGN_GROUP", "CONVERSION", "CONVERSATION_NODE", "CONVERSATION_NODE_OPTION_INDEX", "SERVING_LOCATION", "CARD_INDEX", "MEMBER_COMPANY_SIZE", "MEMBER_INDUSTRY", "MEMBER_SENIORITY", "MEMBER_JOB_TITLE ", "MEMBER_JOB_FUNCTION ", "MEMBER_COUNTRY_V2 ", "MEMBER_REGION_V2", "MEMBER_COMPANY", "PLACEMENT_NAME", "IMPRESSION_DEVICE_TYPE"]` + "\n" +
										`Select value from list to pivot by`,
								},
								"time_granularity": schema.StringAttribute{
									Required: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"ALL",
											"DAILY",
											"MONTHLY",
											"YEARLY",
										),
									},
									MarkdownDescription: `must be one of ["ALL", "DAILY", "MONTHLY", "YEARLY"]` + "\n" +
										`Set time granularity for report: ` + "\n" +
										`ALL - Results grouped into a single result across the entire time range of the report.` + "\n" +
										`DAILY - Results grouped by day.` + "\n" +
										`MONTHLY - Results grouped by month.` + "\n" +
										`YEARLY - Results grouped by year.`,
								},
							},
						},
					},
					"credentials": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_linkedin_ads_authentication_access_token": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Required:    true,
										Description: `The token value generated using the authentication code. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-ads#authentication">docs</a> to obtain yours.`,
									},
									"auth_method": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"access_token",
											),
										},
										Description: `must be one of ["access_token"]`,
									},
								},
							},
							"source_linkedin_ads_authentication_o_auth2_0": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"auth_method": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oAuth2.0",
											),
										},
										Description: `must be one of ["oAuth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Required:    true,
										Description: `The client ID of the LinkedIn Ads developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Required:    true,
										Description: `The client secret the LinkedIn Ads developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Required:    true,
										Description: `The key to refresh the expired access token.`,
									},
								},
							},
							"source_linkedin_ads_update_authentication_access_token": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Required:    true,
										Description: `The token value generated using the authentication code. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-ads#authentication">docs</a> to obtain yours.`,
									},
									"auth_method": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"access_token",
											),
										},
										Description: `must be one of ["access_token"]`,
									},
								},
							},
							"source_linkedin_ads_update_authentication_o_auth2_0": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"auth_method": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oAuth2.0",
											),
										},
										Description: `must be one of ["oAuth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Required:    true,
										Description: `The client ID of the LinkedIn Ads developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Required:    true,
										Description: `The client secret the LinkedIn Ads developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Required:    true,
										Description: `The key to refresh the expired access token.`,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"source_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"linkedin-ads",
							),
						},
						Description: `must be one of ["linkedin-ads"]`,
					},
					"start_date": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `UTC date in the format 2020-09-17. Any data before this date will not be replicated.`,
					},
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required: true,
			},
			"secret_id": schema.StringAttribute{
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"source_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required: true,
			},
		},
	}
}

func (r *SourceLinkedinAdsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceLinkedinAdsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceLinkedinAdsResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	request := *data.ToCreateSDKType()
	res, err := r.client.Sources.CreateSourceLinkedinAds(ctx, request)
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
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceLinkedinAdsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceLinkedinAdsResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.GetSourceLinkedinAdsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceLinkedinAds(ctx, request)
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

func (r *SourceLinkedinAdsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceLinkedinAdsResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceLinkedinAdsPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceLinkedinAdsRequest{
		SourceLinkedinAdsPutRequest: sourceLinkedinAdsPutRequest,
		SourceID:                    sourceID,
	}
	res, err := r.client.Sources.PutSourceLinkedinAds(ctx, request)
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
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	sourceId1 := data.SourceID.ValueString()
	getRequest := operations.GetSourceLinkedinAdsRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceLinkedinAds(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if getResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", getResponse))
		return
	}
	if getResponse.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", getResponse.StatusCode), debugResponse(getResponse.RawResponse))
		return
	}
	if getResponse.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceLinkedinAdsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceLinkedinAdsResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.DeleteSourceLinkedinAdsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceLinkedinAds(ctx, request)
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
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SourceLinkedinAdsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}

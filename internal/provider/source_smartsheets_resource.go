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
var _ resource.Resource = &SourceSmartsheetsResource{}
var _ resource.ResourceWithImportState = &SourceSmartsheetsResource{}

func NewSourceSmartsheetsResource() resource.Resource {
	return &SourceSmartsheetsResource{}
}

// SourceSmartsheetsResource defines the resource implementation.
type SourceSmartsheetsResource struct {
	client *sdk.SDK
}

// SourceSmartsheetsResourceModel describes the resource data model.
type SourceSmartsheetsResourceModel struct {
	Configuration SourceSmartsheets `tfsdk:"configuration"`
	Name          types.String      `tfsdk:"name"`
	SecretID      types.String      `tfsdk:"secret_id"`
	SourceID      types.String      `tfsdk:"source_id"`
	SourceType    types.String      `tfsdk:"source_type"`
	WorkspaceID   types.String      `tfsdk:"workspace_id"`
}

func (r *SourceSmartsheetsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_smartsheets"
}

func (r *SourceSmartsheetsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceSmartsheets Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"source_smartsheets_authorization_method_api_access_token": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Required:    true,
										Description: `The access token to use for accessing your data from Smartsheets. This access token must be generated by a user with at least read access to the data you'd like to replicate. Generate an access token in the Smartsheets main menu by clicking Account > Apps & Integrations > API Access. See the <a href="https://docs.airbyte.com/integrations/sources/smartsheets/#setup-guide">setup guide</a> for information on how to obtain this token.`,
									},
									"auth_type": schema.StringAttribute{
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
							"source_smartsheets_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Required:    true,
										Description: `Access Token for making authenticated requests.`,
									},
									"auth_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
										Description: `must be one of ["oauth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Required:    true,
										Description: `The API ID of the SmartSheets developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Required:    true,
										Description: `The API Secret the SmartSheets developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Required:    true,
										Description: `The key to refresh the expired access_token.`,
									},
									"token_expiry_date": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											validators.IsRFC3339(),
										},
										Description: `The date-time when the access token should be refreshed.`,
									},
								},
							},
							"source_smartsheets_update_authorization_method_api_access_token": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Required:    true,
										Description: `The access token to use for accessing your data from Smartsheets. This access token must be generated by a user with at least read access to the data you'd like to replicate. Generate an access token in the Smartsheets main menu by clicking Account > Apps & Integrations > API Access. See the <a href="https://docs.airbyte.com/integrations/sources/smartsheets/#setup-guide">setup guide</a> for information on how to obtain this token.`,
									},
									"auth_type": schema.StringAttribute{
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
							"source_smartsheets_update_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Required:    true,
										Description: `Access Token for making authenticated requests.`,
									},
									"auth_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
										Description: `must be one of ["oauth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Required:    true,
										Description: `The API ID of the SmartSheets developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Required:    true,
										Description: `The API Secret the SmartSheets developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Required:    true,
										Description: `The key to refresh the expired access_token.`,
									},
									"token_expiry_date": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											validators.IsRFC3339(),
										},
										Description: `The date-time when the access token should be refreshed.`,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"metadata_fields": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
						Description: `A List of available columns which metadata can be pulled from.`,
					},
					"source_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"smartsheets",
							),
						},
						Description: `must be one of ["smartsheets"]`,
					},
					"spreadsheet_id": schema.StringAttribute{
						Required:    true,
						Description: `The spreadsheet ID. Find it by opening the spreadsheet then navigating to File > Properties`,
					},
					"start_datetime": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `Only rows modified after this date/time will be replicated. This should be an ISO 8601 string, for instance: ` + "`" + `2000-01-01T13:00:00` + "`" + ``,
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

func (r *SourceSmartsheetsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceSmartsheetsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceSmartsheetsResourceModel
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
	res, err := r.client.Sources.CreateSourceSmartsheets(ctx, request)
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

func (r *SourceSmartsheetsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceSmartsheetsResourceModel
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
	request := operations.GetSourceSmartsheetsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceSmartsheets(ctx, request)
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

func (r *SourceSmartsheetsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceSmartsheetsResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceSmartsheetsPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceSmartsheetsRequest{
		SourceSmartsheetsPutRequest: sourceSmartsheetsPutRequest,
		SourceID:                    sourceID,
	}
	res, err := r.client.Sources.PutSourceSmartsheets(ctx, request)
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
	getRequest := operations.GetSourceSmartsheetsRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceSmartsheets(ctx, getRequest)
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

func (r *SourceSmartsheetsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceSmartsheetsResourceModel
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
	request := operations.DeleteSourceSmartsheetsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceSmartsheets(ctx, request)
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

func (r *SourceSmartsheetsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}

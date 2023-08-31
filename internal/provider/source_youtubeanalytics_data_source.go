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
var _ datasource.DataSource = &SourceYoutubeAnalyticsDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceYoutubeAnalyticsDataSource{}

func NewSourceYoutubeAnalyticsDataSource() datasource.DataSource {
	return &SourceYoutubeAnalyticsDataSource{}
}

// SourceYoutubeAnalyticsDataSource is the data source implementation.
type SourceYoutubeAnalyticsDataSource struct {
	client *sdk.SDK
}

// SourceYoutubeAnalyticsDataSourceModel describes the data model.
type SourceYoutubeAnalyticsDataSourceModel struct {
	Configuration SourceYoutubeAnalytics1 `tfsdk:"configuration"`
	Name          types.String            `tfsdk:"name"`
	SecretID      types.String            `tfsdk:"secret_id"`
	SourceID      types.String            `tfsdk:"source_id"`
	WorkspaceID   types.String            `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceYoutubeAnalyticsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_youtube_analytics"
}

// Schema defines the schema for the data source.
func (r *SourceYoutubeAnalyticsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceYoutubeAnalytics DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"client_id": schema.StringAttribute{
								Computed:    true,
								Description: `The Client ID of your developer application`,
							},
							"client_secret": schema.StringAttribute{
								Computed:    true,
								Description: `The client secret of your developer application`,
							},
							"refresh_token": schema.StringAttribute{
								Computed:    true,
								Description: `A refresh token generated using the above client ID and secret`,
							},
							"additional_properties": schema.StringAttribute{
								Optional: true,
								Validators: []validator.String{
									validators.IsValidJSON(),
								},
								Description: `Parsed as JSON.`,
							},
						},
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"youtube-analytics",
							),
						},
						Description: `must be one of ["youtube-analytics"]`,
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

func (r *SourceYoutubeAnalyticsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceYoutubeAnalyticsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceYoutubeAnalyticsDataSourceModel
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
	request := operations.GetSourceYoutubeAnalyticsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceYoutubeAnalytics(ctx, request)
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

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
var _ datasource.DataSource = &SourceSalesforceDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceSalesforceDataSource{}

func NewSourceSalesforceDataSource() datasource.DataSource {
	return &SourceSalesforceDataSource{}
}

// SourceSalesforceDataSource is the data source implementation.
type SourceSalesforceDataSource struct {
	client *sdk.SDK
}

// SourceSalesforceDataSourceModel describes the data model.
type SourceSalesforceDataSourceModel struct {
	Configuration SourceSalesforce `tfsdk:"configuration"`
	Name          types.String     `tfsdk:"name"`
	SecretID      types.String     `tfsdk:"secret_id"`
	SourceID      types.String     `tfsdk:"source_id"`
	WorkspaceID   types.String     `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceSalesforceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_salesforce"
}

// Schema defines the schema for the data source.
func (r *SourceSalesforceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceSalesforce DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"auth_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Client",
							),
						},
						Description: `must be one of ["Client"]`,
					},
					"client_id": schema.StringAttribute{
						Computed:    true,
						Description: `Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client ID</a>`,
					},
					"client_secret": schema.StringAttribute{
						Computed:    true,
						Description: `Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client secret</a>`,
					},
					"force_use_bulk_api": schema.BoolAttribute{
						Computed:    true,
						Description: `Toggle to use Bulk API (this might cause empty fields for some streams)`,
					},
					"is_sandbox": schema.BoolAttribute{
						Computed:    true,
						Description: `Toggle if you're using a <a href="https://help.salesforce.com/s/articleView?id=sf.deploy_sandboxes_parent.htm&type=5">Salesforce Sandbox</a>`,
					},
					"refresh_token": schema.StringAttribute{
						Computed:    true,
						Description: `Enter your application's <a href="https://developer.salesforce.com/docs/atlas.en-us.mobile_sdk.meta/mobile_sdk/oauth_refresh_token_flow.htm">Salesforce Refresh Token</a> used for Airbyte to access your Salesforce account.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"salesforce",
							),
						},
						Description: `must be one of ["salesforce"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `Enter the date (or date-time) in the YYYY-MM-DD or YYYY-MM-DDTHH:mm:ssZ format. Airbyte will replicate the data updated on and after this date. If this field is blank, Airbyte will replicate the data for last two years.`,
					},
					"streams_criteria": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"criteria": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"starts with",
											"ends with",
											"contains",
											"exacts",
											"starts not with",
											"ends not with",
											"not contains",
											"not exacts",
										),
									},
									Description: `must be one of ["starts with", "ends with", "contains", "exacts", "starts not with", "ends not with", "not contains", "not exacts"]`,
								},
								"value": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						Description: `Add filters to select only required stream based on ` + "`" + `SObject` + "`" + ` name. Use this field to filter which tables are displayed by this connector. This is useful if your Salesforce account has a large number of tables (>1000), in which case you may find it easier to navigate the UI and speed up the connector's performance if you restrict the tables displayed by this connector.`,
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

func (r *SourceSalesforceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceSalesforceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceSalesforceDataSourceModel
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
	request := operations.GetSourceSalesforceRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceSalesforce(ctx, request)
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

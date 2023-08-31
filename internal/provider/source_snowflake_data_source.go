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
var _ datasource.DataSource = &SourceSnowflakeDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceSnowflakeDataSource{}

func NewSourceSnowflakeDataSource() datasource.DataSource {
	return &SourceSnowflakeDataSource{}
}

// SourceSnowflakeDataSource is the data source implementation.
type SourceSnowflakeDataSource struct {
	client *sdk.SDK
}

// SourceSnowflakeDataSourceModel describes the data model.
type SourceSnowflakeDataSourceModel struct {
	Configuration SourceSnowflake `tfsdk:"configuration"`
	Name          types.String    `tfsdk:"name"`
	SecretID      types.String    `tfsdk:"secret_id"`
	SourceID      types.String    `tfsdk:"source_id"`
	WorkspaceID   types.String    `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceSnowflakeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_snowflake"
}

// Schema defines the schema for the data source.
func (r *SourceSnowflakeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceSnowflake DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_snowflake_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Access Token for making authenticated requests.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth",
											),
										},
										Description: `must be one of ["OAuth"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of your Snowflake developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Client Secret of your Snowflake developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Refresh Token for making authenticated requests.`,
									},
								},
							},
							"source_snowflake_authorization_method_username_and_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"username/password",
											),
										},
										Description: `must be one of ["username/password"]`,
									},
									"password": schema.StringAttribute{
										Computed:    true,
										Description: `The password associated with the username.`,
									},
									"username": schema.StringAttribute{
										Computed:    true,
										Description: `The username you created to allow Airbyte to access the database.`,
									},
								},
							},
							"source_snowflake_update_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Access Token for making authenticated requests.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth",
											),
										},
										Description: `must be one of ["OAuth"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of your Snowflake developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Client Secret of your Snowflake developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Refresh Token for making authenticated requests.`,
									},
								},
							},
							"source_snowflake_update_authorization_method_username_and_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"username/password",
											),
										},
										Description: `must be one of ["username/password"]`,
									},
									"password": schema.StringAttribute{
										Computed:    true,
										Description: `The password associated with the username.`,
									},
									"username": schema.StringAttribute{
										Computed:    true,
										Description: `The username you created to allow Airbyte to access the database.`,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"database": schema.StringAttribute{
						Computed:    true,
						Description: `The database you created for Airbyte to access data.`,
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Description: `The host domain of the snowflake instance (must include the account, region, cloud environment, and end with snowflakecomputing.com).`,
					},
					"jdbc_url_params": schema.StringAttribute{
						Computed:    true,
						Description: `Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).`,
					},
					"role": schema.StringAttribute{
						Computed:    true,
						Description: `The role you created for Airbyte to access Snowflake.`,
					},
					"schema": schema.StringAttribute{
						Computed:    true,
						Description: `The source Snowflake schema tables. Leave empty to access tables from multiple schemas.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"snowflake",
							),
						},
						Description: `must be one of ["snowflake"]`,
					},
					"warehouse": schema.StringAttribute{
						Computed:    true,
						Description: `The warehouse you created for Airbyte to access data.`,
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

func (r *SourceSnowflakeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceSnowflakeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceSnowflakeDataSourceModel
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
	request := operations.GetSourceSnowflakeRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceSnowflake(ctx, request)
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

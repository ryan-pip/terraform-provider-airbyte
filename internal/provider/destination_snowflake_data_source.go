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
var _ datasource.DataSource = &DestinationSnowflakeDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationSnowflakeDataSource{}

func NewDestinationSnowflakeDataSource() datasource.DataSource {
	return &DestinationSnowflakeDataSource{}
}

// DestinationSnowflakeDataSource is the data source implementation.
type DestinationSnowflakeDataSource struct {
	client *sdk.SDK
}

// DestinationSnowflakeDataSourceModel describes the data model.
type DestinationSnowflakeDataSourceModel struct {
	Configuration DestinationSnowflake `tfsdk:"configuration"`
	DestinationID types.String         `tfsdk:"destination_id"`
	Name          types.String         `tfsdk:"name"`
	WorkspaceID   types.String         `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationSnowflakeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_snowflake"
}

// Schema defines the schema for the data source.
func (r *DestinationSnowflakeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationSnowflake DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"destination_snowflake_authorization_method_key_pair_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Key Pair Authentication",
											),
										},
										Description: `must be one of ["Key Pair Authentication"]`,
									},
									"private_key": schema.StringAttribute{
										Computed:    true,
										Description: `RSA Private key to use for Snowflake connection. See the <a href="https://docs.airbyte.com/integrations/destinations/snowflake">docs</a> for more information on how to obtain this key.`,
									},
									"private_key_password": schema.StringAttribute{
										Computed:    true,
										Description: `Passphrase for private key`,
									},
								},
							},
							"destination_snowflake_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Enter you application's Access Token`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth2.0",
											),
										},
										Description: `must be one of ["OAuth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Client ID`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Client secret`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Refresh Token`,
									},
								},
							},
							"destination_snowflake_authorization_method_username_and_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Username and Password",
											),
										},
										Description: `must be one of ["Username and Password"]`,
									},
									"password": schema.StringAttribute{
										Computed:    true,
										Description: `Enter the password associated with the username.`,
									},
								},
							},
							"destination_snowflake_update_authorization_method_key_pair_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Key Pair Authentication",
											),
										},
										Description: `must be one of ["Key Pair Authentication"]`,
									},
									"private_key": schema.StringAttribute{
										Computed:    true,
										Description: `RSA Private key to use for Snowflake connection. See the <a href="https://docs.airbyte.com/integrations/destinations/snowflake">docs</a> for more information on how to obtain this key.`,
									},
									"private_key_password": schema.StringAttribute{
										Computed:    true,
										Description: `Passphrase for private key`,
									},
								},
							},
							"destination_snowflake_update_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Enter you application's Access Token`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth2.0",
											),
										},
										Description: `must be one of ["OAuth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Client ID`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Client secret`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Refresh Token`,
									},
								},
							},
							"destination_snowflake_update_authorization_method_username_and_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Username and Password",
											),
										},
										Description: `must be one of ["Username and Password"]`,
									},
									"password": schema.StringAttribute{
										Computed:    true,
										Description: `Enter the password associated with the username.`,
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
						Description: `Enter the name of the <a href="https://docs.snowflake.com/en/sql-reference/ddl-database.html#database-schema-share-ddl">database</a> you want to sync data into`,
					},
					"destination_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"snowflake",
							),
						},
						Description: `must be one of ["snowflake"]`,
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Description: `Enter your Snowflake account's <a href="https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#using-an-account-locator-as-an-identifier">locator</a> (in the format <account_locator>.<region>.<cloud>.snowflakecomputing.com)`,
					},
					"jdbc_url_params": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the additional properties to pass to the JDBC URL string when connecting to the database (formatted as key=value pairs separated by the symbol &). Example: key1=value1&key2=value2&key3=value3`,
					},
					"raw_data_schema": schema.StringAttribute{
						Computed:    true,
						Description: `(Beta) The schema to write raw tables into`,
					},
					"role": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the <a href="https://docs.snowflake.com/en/user-guide/security-access-control-overview.html#roles">role</a> that you want to use to access Snowflake`,
					},
					"schema": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the name of the default <a href="https://docs.snowflake.com/en/sql-reference/ddl-database.html#database-schema-share-ddl">schema</a>`,
					},
					"use_1s1t_format": schema.BoolAttribute{
						Computed:    true,
						Description: `(Beta) Use <a href="https://github.com/airbytehq/airbyte/issues/26028" target="_blank">Destinations V2</a>. Contact Airbyte Support to participate in the beta program.`,
					},
					"username": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the name of the user you want to use to access the database`,
					},
					"warehouse": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the name of the <a href="https://docs.snowflake.com/en/user-guide/warehouses-overview.html#overview-of-warehouses">warehouse</a> that you want to sync data into`,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *DestinationSnowflakeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DestinationSnowflakeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationSnowflakeDataSourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationSnowflakeRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationSnowflake(ctx, request)
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
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

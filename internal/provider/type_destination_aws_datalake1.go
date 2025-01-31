// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationAwsDatalake1 struct {
	AwsAccountID                          types.String                                `tfsdk:"aws_account_id"`
	BucketName                            types.String                                `tfsdk:"bucket_name"`
	BucketPrefix                          types.String                                `tfsdk:"bucket_prefix"`
	Credentials                           DestinationAwsDatalakeAuthenticationMode    `tfsdk:"credentials"`
	DestinationType                       types.String                                `tfsdk:"destination_type"`
	Format                                *DestinationAwsDatalakeOutputFormatWildcard `tfsdk:"format"`
	GlueCatalogFloatAsDecimal             types.Bool                                  `tfsdk:"glue_catalog_float_as_decimal"`
	LakeformationDatabaseDefaultTagKey    types.String                                `tfsdk:"lakeformation_database_default_tag_key"`
	LakeformationDatabaseDefaultTagValues types.String                                `tfsdk:"lakeformation_database_default_tag_values"`
	LakeformationDatabaseName             types.String                                `tfsdk:"lakeformation_database_name"`
	LakeformationGovernedTables           types.Bool                                  `tfsdk:"lakeformation_governed_tables"`
	Partitioning                          types.String                                `tfsdk:"partitioning"`
	Region                                types.String                                `tfsdk:"region"`
}

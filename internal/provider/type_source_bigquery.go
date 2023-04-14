// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceBigquery struct {
	CredentialsJSON types.String `tfsdk:"credentials_json"`
	DatasetID       types.String `tfsdk:"dataset_id"`
	ProjectID       types.String `tfsdk:"project_id"`
	SourceType      types.String `tfsdk:"source_type"`
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceS3 struct {
	Dataset     types.String                `tfsdk:"dataset"`
	Format      *SourceS3FileFormat         `tfsdk:"format"`
	PathPattern types.String                `tfsdk:"path_pattern"`
	Provider    SourceS3S3AmazonWebServices `tfsdk:"provider"`
	Schema      types.String                `tfsdk:"schema"`
	SourceType  types.String                `tfsdk:"source_type"`
}

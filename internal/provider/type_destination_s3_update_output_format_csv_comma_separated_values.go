// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationS3UpdateOutputFormatCSVCommaSeparatedValues struct {
	Compression *DestinationS3UpdateOutputFormatCSVCommaSeparatedValuesCompression `tfsdk:"compression"`
	Flattening  types.String                                                       `tfsdk:"flattening"`
	FormatType  types.String                                                       `tfsdk:"format_type"`
}

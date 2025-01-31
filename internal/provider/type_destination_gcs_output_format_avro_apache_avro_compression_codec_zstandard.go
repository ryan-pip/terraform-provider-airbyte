// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationGcsOutputFormatAvroApacheAvroCompressionCodecZstandard struct {
	Codec            types.String `tfsdk:"codec"`
	CompressionLevel types.Int64  `tfsdk:"compression_level"`
	IncludeChecksum  types.Bool   `tfsdk:"include_checksum"`
}

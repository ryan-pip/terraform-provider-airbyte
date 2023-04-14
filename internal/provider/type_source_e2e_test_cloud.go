// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceE2eTestCloud struct {
	MaxMessages       types.Int64                   `tfsdk:"max_messages"`
	MessageIntervalMs types.Int64                   `tfsdk:"message_interval_ms"`
	MockCatalog       SourceE2eTestCloudMockCatalog `tfsdk:"mock_catalog"`
	Seed              types.Int64                   `tfsdk:"seed"`
	SourceType        types.String                  `tfsdk:"source_type"`
	Type              types.String                  `tfsdk:"type"`
}

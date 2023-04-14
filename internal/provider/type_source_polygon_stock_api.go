// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourcePolygonStockAPI struct {
	Adjusted     types.String `tfsdk:"adjusted"`
	APIKey       types.String `tfsdk:"api_key"`
	EndDate      types.String `tfsdk:"end_date"`
	Limit        types.Int64  `tfsdk:"limit"`
	Multiplier   types.Int64  `tfsdk:"multiplier"`
	Sort         types.String `tfsdk:"sort"`
	SourceType   types.String `tfsdk:"source_type"`
	StartDate    types.String `tfsdk:"start_date"`
	StocksTicker types.String `tfsdk:"stocks_ticker"`
	Timespan     types.String `tfsdk:"timespan"`
}

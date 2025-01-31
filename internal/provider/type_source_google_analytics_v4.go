// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleAnalyticsV4 struct {
	Credentials   *SourceGoogleAnalyticsV4Credentials `tfsdk:"credentials"`
	CustomReports types.String                        `tfsdk:"custom_reports"`
	SourceType    types.String                        `tfsdk:"source_type"`
	StartDate     types.String                        `tfsdk:"start_date"`
	ViewID        types.String                        `tfsdk:"view_id"`
	WindowInDays  types.Int64                         `tfsdk:"window_in_days"`
}

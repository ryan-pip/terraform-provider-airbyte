// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMixpanel struct {
	AttributionWindow         types.Int64                           `tfsdk:"attribution_window"`
	Credentials               *SourceMixpanelAuthenticationWildcard `tfsdk:"credentials"`
	DateWindowSize            types.Int64                           `tfsdk:"date_window_size"`
	EndDate                   types.String                          `tfsdk:"end_date"`
	ProjectID                 types.Int64                           `tfsdk:"project_id"`
	ProjectTimezone           types.String                          `tfsdk:"project_timezone"`
	Region                    types.String                          `tfsdk:"region"`
	SelectPropertiesByDefault types.Bool                            `tfsdk:"select_properties_by_default"`
	SourceType                types.String                          `tfsdk:"source_type"`
	StartDate                 types.String                          `tfsdk:"start_date"`
}

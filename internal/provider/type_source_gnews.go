// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGnews struct {
	APIKey            types.String   `tfsdk:"api_key"`
	Country           types.String   `tfsdk:"country"`
	EndDate           types.String   `tfsdk:"end_date"`
	In                []types.String `tfsdk:"in"`
	Language          types.String   `tfsdk:"language"`
	Nullable          []types.String `tfsdk:"nullable"`
	Query             types.String   `tfsdk:"query"`
	Sortby            types.String   `tfsdk:"sortby"`
	SourceType        types.String   `tfsdk:"source_type"`
	StartDate         types.String   `tfsdk:"start_date"`
	TopHeadlinesQuery types.String   `tfsdk:"top_headlines_query"`
	TopHeadlinesTopic types.String   `tfsdk:"top_headlines_topic"`
}

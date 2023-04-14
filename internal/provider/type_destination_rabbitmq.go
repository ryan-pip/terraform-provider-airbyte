// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationRabbitmq struct {
	DestinationType types.String `tfsdk:"destination_type"`
	Exchange        types.String `tfsdk:"exchange"`
	Host            types.String `tfsdk:"host"`
	Password        types.String `tfsdk:"password"`
	Port            types.Int64  `tfsdk:"port"`
	RoutingKey      types.String `tfsdk:"routing_key"`
	Ssl             types.Bool   `tfsdk:"ssl"`
	Username        types.String `tfsdk:"username"`
	VirtualHost     types.String `tfsdk:"virtual_host"`
}

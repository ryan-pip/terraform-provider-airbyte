// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationMongodbMongoDbInstanceTypeReplicaSet struct {
	Instance        types.String `tfsdk:"instance"`
	ReplicaSet      types.String `tfsdk:"replica_set"`
	ServerAddresses types.String `tfsdk:"server_addresses"`
}

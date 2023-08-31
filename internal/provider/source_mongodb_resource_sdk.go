// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMongodbResourceModel) ToCreateSDKType() *shared.SourceMongodbCreateRequest {
	authSource := new(string)
	if !r.Configuration.AuthSource.IsUnknown() && !r.Configuration.AuthSource.IsNull() {
		*authSource = r.Configuration.AuthSource.ValueString()
	} else {
		authSource = nil
	}
	database := r.Configuration.Database.ValueString()
	var instanceType *shared.SourceMongodbMongoDbInstanceType
	if r.Configuration.InstanceType != nil {
		var sourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance *shared.SourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance
		if r.Configuration.InstanceType.SourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
			host := r.Configuration.InstanceType.SourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance.Host.ValueString()
			instance := shared.SourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstance(r.Configuration.InstanceType.SourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance.Instance.ValueString())
			port := r.Configuration.InstanceType.SourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance.Port.ValueInt64()
			sourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance = &shared.SourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance{
				Host:     host,
				Instance: instance,
				Port:     port,
			}
		}
		if sourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
			instanceType = &shared.SourceMongodbMongoDbInstanceType{
				SourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance: sourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance,
			}
		}
		var sourceMongodbMongoDbInstanceTypeReplicaSet *shared.SourceMongodbMongoDbInstanceTypeReplicaSet
		if r.Configuration.InstanceType.SourceMongodbMongoDbInstanceTypeReplicaSet != nil {
			instance1 := shared.SourceMongodbMongoDbInstanceTypeReplicaSetInstance(r.Configuration.InstanceType.SourceMongodbMongoDbInstanceTypeReplicaSet.Instance.ValueString())
			replicaSet := new(string)
			if !r.Configuration.InstanceType.SourceMongodbMongoDbInstanceTypeReplicaSet.ReplicaSet.IsUnknown() && !r.Configuration.InstanceType.SourceMongodbMongoDbInstanceTypeReplicaSet.ReplicaSet.IsNull() {
				*replicaSet = r.Configuration.InstanceType.SourceMongodbMongoDbInstanceTypeReplicaSet.ReplicaSet.ValueString()
			} else {
				replicaSet = nil
			}
			serverAddresses := r.Configuration.InstanceType.SourceMongodbMongoDbInstanceTypeReplicaSet.ServerAddresses.ValueString()
			sourceMongodbMongoDbInstanceTypeReplicaSet = &shared.SourceMongodbMongoDbInstanceTypeReplicaSet{
				Instance:        instance1,
				ReplicaSet:      replicaSet,
				ServerAddresses: serverAddresses,
			}
		}
		if sourceMongodbMongoDbInstanceTypeReplicaSet != nil {
			instanceType = &shared.SourceMongodbMongoDbInstanceType{
				SourceMongodbMongoDbInstanceTypeReplicaSet: sourceMongodbMongoDbInstanceTypeReplicaSet,
			}
		}
		var sourceMongodbMongoDBInstanceTypeMongoDBAtlas *shared.SourceMongodbMongoDBInstanceTypeMongoDBAtlas
		if r.Configuration.InstanceType.SourceMongodbMongoDBInstanceTypeMongoDBAtlas != nil {
			clusterURL := r.Configuration.InstanceType.SourceMongodbMongoDBInstanceTypeMongoDBAtlas.ClusterURL.ValueString()
			instance2 := shared.SourceMongodbMongoDBInstanceTypeMongoDBAtlasInstance(r.Configuration.InstanceType.SourceMongodbMongoDBInstanceTypeMongoDBAtlas.Instance.ValueString())
			var additionalProperties interface{}
			if !r.Configuration.InstanceType.SourceMongodbMongoDBInstanceTypeMongoDBAtlas.AdditionalProperties.IsUnknown() && !r.Configuration.InstanceType.SourceMongodbMongoDBInstanceTypeMongoDBAtlas.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.InstanceType.SourceMongodbMongoDBInstanceTypeMongoDBAtlas.AdditionalProperties.ValueString()), &additionalProperties)
			}
			sourceMongodbMongoDBInstanceTypeMongoDBAtlas = &shared.SourceMongodbMongoDBInstanceTypeMongoDBAtlas{
				ClusterURL:           clusterURL,
				Instance:             instance2,
				AdditionalProperties: additionalProperties,
			}
		}
		if sourceMongodbMongoDBInstanceTypeMongoDBAtlas != nil {
			instanceType = &shared.SourceMongodbMongoDbInstanceType{
				SourceMongodbMongoDBInstanceTypeMongoDBAtlas: sourceMongodbMongoDBInstanceTypeMongoDBAtlas,
			}
		}
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	sourceType := shared.SourceMongodbMongodb(r.Configuration.SourceType.ValueString())
	user := new(string)
	if !r.Configuration.User.IsUnknown() && !r.Configuration.User.IsNull() {
		*user = r.Configuration.User.ValueString()
	} else {
		user = nil
	}
	configuration := shared.SourceMongodb{
		AuthSource:   authSource,
		Database:     database,
		InstanceType: instanceType,
		Password:     password,
		SourceType:   sourceType,
		User:         user,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMongodbCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMongodbResourceModel) ToGetSDKType() *shared.SourceMongodbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMongodbResourceModel) ToUpdateSDKType() *shared.SourceMongodbPutRequest {
	authSource := new(string)
	if !r.Configuration.AuthSource.IsUnknown() && !r.Configuration.AuthSource.IsNull() {
		*authSource = r.Configuration.AuthSource.ValueString()
	} else {
		authSource = nil
	}
	database := r.Configuration.Database.ValueString()
	var instanceType *shared.SourceMongodbUpdateMongoDbInstanceType
	if r.Configuration.InstanceType != nil {
		var sourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance *shared.SourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance
		if r.Configuration.InstanceType.SourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
			host := r.Configuration.InstanceType.SourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance.Host.ValueString()
			instance := shared.SourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstance(r.Configuration.InstanceType.SourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance.Instance.ValueString())
			port := r.Configuration.InstanceType.SourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance.Port.ValueInt64()
			sourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance = &shared.SourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance{
				Host:     host,
				Instance: instance,
				Port:     port,
			}
		}
		if sourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
			instanceType = &shared.SourceMongodbUpdateMongoDbInstanceType{
				SourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance: sourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance,
			}
		}
		var sourceMongodbUpdateMongoDbInstanceTypeReplicaSet *shared.SourceMongodbUpdateMongoDbInstanceTypeReplicaSet
		if r.Configuration.InstanceType.SourceMongodbUpdateMongoDbInstanceTypeReplicaSet != nil {
			instance1 := shared.SourceMongodbUpdateMongoDbInstanceTypeReplicaSetInstance(r.Configuration.InstanceType.SourceMongodbUpdateMongoDbInstanceTypeReplicaSet.Instance.ValueString())
			replicaSet := new(string)
			if !r.Configuration.InstanceType.SourceMongodbUpdateMongoDbInstanceTypeReplicaSet.ReplicaSet.IsUnknown() && !r.Configuration.InstanceType.SourceMongodbUpdateMongoDbInstanceTypeReplicaSet.ReplicaSet.IsNull() {
				*replicaSet = r.Configuration.InstanceType.SourceMongodbUpdateMongoDbInstanceTypeReplicaSet.ReplicaSet.ValueString()
			} else {
				replicaSet = nil
			}
			serverAddresses := r.Configuration.InstanceType.SourceMongodbUpdateMongoDbInstanceTypeReplicaSet.ServerAddresses.ValueString()
			sourceMongodbUpdateMongoDbInstanceTypeReplicaSet = &shared.SourceMongodbUpdateMongoDbInstanceTypeReplicaSet{
				Instance:        instance1,
				ReplicaSet:      replicaSet,
				ServerAddresses: serverAddresses,
			}
		}
		if sourceMongodbUpdateMongoDbInstanceTypeReplicaSet != nil {
			instanceType = &shared.SourceMongodbUpdateMongoDbInstanceType{
				SourceMongodbUpdateMongoDbInstanceTypeReplicaSet: sourceMongodbUpdateMongoDbInstanceTypeReplicaSet,
			}
		}
		var sourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas *shared.SourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas
		if r.Configuration.InstanceType.SourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas != nil {
			clusterURL := r.Configuration.InstanceType.SourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas.ClusterURL.ValueString()
			instance2 := shared.SourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstance(r.Configuration.InstanceType.SourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas.Instance.ValueString())
			var additionalProperties interface{}
			if !r.Configuration.InstanceType.SourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas.AdditionalProperties.IsUnknown() && !r.Configuration.InstanceType.SourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.InstanceType.SourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas.AdditionalProperties.ValueString()), &additionalProperties)
			}
			sourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas = &shared.SourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas{
				ClusterURL:           clusterURL,
				Instance:             instance2,
				AdditionalProperties: additionalProperties,
			}
		}
		if sourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas != nil {
			instanceType = &shared.SourceMongodbUpdateMongoDbInstanceType{
				SourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas: sourceMongodbUpdateMongoDBInstanceTypeMongoDBAtlas,
			}
		}
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	user := new(string)
	if !r.Configuration.User.IsUnknown() && !r.Configuration.User.IsNull() {
		*user = r.Configuration.User.ValueString()
	} else {
		user = nil
	}
	configuration := shared.SourceMongodbUpdate{
		AuthSource:   authSource,
		Database:     database,
		InstanceType: instanceType,
		Password:     password,
		User:         user,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMongodbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMongodbResourceModel) ToDeleteSDKType() *shared.SourceMongodbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMongodbResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMongodbResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}

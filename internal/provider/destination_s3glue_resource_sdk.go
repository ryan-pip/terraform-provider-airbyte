// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationS3GlueResourceModel) ToCreateSDKType() *shared.DestinationS3GlueCreateRequest {
	accessKeyID := new(string)
	if !r.Configuration.AccessKeyID.IsUnknown() && !r.Configuration.AccessKeyID.IsNull() {
		*accessKeyID = r.Configuration.AccessKeyID.ValueString()
	} else {
		accessKeyID = nil
	}
	destinationType := shared.DestinationS3GlueS3Glue(r.Configuration.DestinationType.ValueString())
	fileNamePattern := new(string)
	if !r.Configuration.FileNamePattern.IsUnknown() && !r.Configuration.FileNamePattern.IsNull() {
		*fileNamePattern = r.Configuration.FileNamePattern.ValueString()
	} else {
		fileNamePattern = nil
	}
	var format shared.DestinationS3GlueOutputFormat
	var destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON *shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON
	if r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		var compression *shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression
		if r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Compression != nil {
			var destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression *shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression
			if r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression != nil {
				compressionType := new(shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionType)
				if !r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression.CompressionType.IsUnknown() && !r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression.CompressionType.IsNull() {
					*compressionType = shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionType(r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression.CompressionType.ValueString())
				} else {
					compressionType = nil
				}
				destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression = &shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression{
					CompressionType: compressionType,
				}
			}
			if destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression != nil {
				compression = &shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression{
					DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression: destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression,
				}
			}
			var destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP *shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP
			if r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP != nil {
				compressionType1 := new(shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionType)
				if !r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP.CompressionType.IsUnknown() && !r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP.CompressionType.IsNull() {
					*compressionType1 = shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionType(r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP.CompressionType.ValueString())
				} else {
					compressionType1 = nil
				}
				destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP = &shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP{
					CompressionType: compressionType1,
				}
			}
			if destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP != nil {
				compression = &shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression{
					DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP: destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP,
				}
			}
		}
		flattening := new(shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlattening)
		if !r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Flattening.IsUnknown() && !r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Flattening.IsNull() {
			*flattening = shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlattening(r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.Flattening.ValueString())
		} else {
			flattening = nil
		}
		formatType := shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFormatType(r.Configuration.Format.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON.FormatType.ValueString())
		destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON = &shared.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON{
			Compression: compression,
			Flattening:  flattening,
			FormatType:  formatType,
		}
	}
	if destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		format = shared.DestinationS3GlueOutputFormat{
			DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON: destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON,
		}
	}
	glueDatabase := r.Configuration.GlueDatabase.ValueString()
	glueSerializationLibrary := shared.DestinationS3GlueSerializationLibrary(r.Configuration.GlueSerializationLibrary.ValueString())
	s3BucketName := r.Configuration.S3BucketName.ValueString()
	s3BucketPath := r.Configuration.S3BucketPath.ValueString()
	s3BucketRegion := shared.DestinationS3GlueS3BucketRegion(r.Configuration.S3BucketRegion.ValueString())
	s3Endpoint := new(string)
	if !r.Configuration.S3Endpoint.IsUnknown() && !r.Configuration.S3Endpoint.IsNull() {
		*s3Endpoint = r.Configuration.S3Endpoint.ValueString()
	} else {
		s3Endpoint = nil
	}
	s3PathFormat := new(string)
	if !r.Configuration.S3PathFormat.IsUnknown() && !r.Configuration.S3PathFormat.IsNull() {
		*s3PathFormat = r.Configuration.S3PathFormat.ValueString()
	} else {
		s3PathFormat = nil
	}
	secretAccessKey := new(string)
	if !r.Configuration.SecretAccessKey.IsUnknown() && !r.Configuration.SecretAccessKey.IsNull() {
		*secretAccessKey = r.Configuration.SecretAccessKey.ValueString()
	} else {
		secretAccessKey = nil
	}
	configuration := shared.DestinationS3Glue{
		AccessKeyID:              accessKeyID,
		DestinationType:          destinationType,
		FileNamePattern:          fileNamePattern,
		Format:                   format,
		GlueDatabase:             glueDatabase,
		GlueSerializationLibrary: glueSerializationLibrary,
		S3BucketName:             s3BucketName,
		S3BucketPath:             s3BucketPath,
		S3BucketRegion:           s3BucketRegion,
		S3Endpoint:               s3Endpoint,
		S3PathFormat:             s3PathFormat,
		SecretAccessKey:          secretAccessKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationS3GlueCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationS3GlueResourceModel) ToGetSDKType() *shared.DestinationS3GlueCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationS3GlueResourceModel) ToUpdateSDKType() *shared.DestinationS3GluePutRequest {
	accessKeyID := new(string)
	if !r.Configuration.AccessKeyID.IsUnknown() && !r.Configuration.AccessKeyID.IsNull() {
		*accessKeyID = r.Configuration.AccessKeyID.ValueString()
	} else {
		accessKeyID = nil
	}
	fileNamePattern := new(string)
	if !r.Configuration.FileNamePattern.IsUnknown() && !r.Configuration.FileNamePattern.IsNull() {
		*fileNamePattern = r.Configuration.FileNamePattern.ValueString()
	} else {
		fileNamePattern = nil
	}
	var format shared.DestinationS3GlueUpdateOutputFormat
	var destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON *shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON
	if r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		var compression *shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompression
		if r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Compression != nil {
			var destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression *shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression
			if r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression != nil {
				compressionType := new(shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionType)
				if !r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression.CompressionType.IsUnknown() && !r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression.CompressionType.IsNull() {
					*compressionType = shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionType(r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression.CompressionType.ValueString())
				} else {
					compressionType = nil
				}
				destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression = &shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression{
					CompressionType: compressionType,
				}
			}
			if destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression != nil {
				compression = &shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompression{
					DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression: destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression,
				}
			}
			var destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP *shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP
			if r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP != nil {
				compressionType1 := new(shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionType)
				if !r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP.CompressionType.IsUnknown() && !r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP.CompressionType.IsNull() {
					*compressionType1 = shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionType(r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP.CompressionType.ValueString())
				} else {
					compressionType1 = nil
				}
				destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP = &shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP{
					CompressionType: compressionType1,
				}
			}
			if destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP != nil {
				compression = &shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompression{
					DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP: destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP,
				}
			}
		}
		flattening := new(shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONFlattening)
		if !r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Flattening.IsUnknown() && !r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Flattening.IsNull() {
			*flattening = shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONFlattening(r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.Flattening.ValueString())
		} else {
			flattening = nil
		}
		formatType := shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType(r.Configuration.Format.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON.FormatType.ValueString())
		destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON = &shared.DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON{
			Compression: compression,
			Flattening:  flattening,
			FormatType:  formatType,
		}
	}
	if destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		format = shared.DestinationS3GlueUpdateOutputFormat{
			DestinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON: destinationS3GlueUpdateOutputFormatJSONLinesNewlineDelimitedJSON,
		}
	}
	glueDatabase := r.Configuration.GlueDatabase.ValueString()
	glueSerializationLibrary := shared.DestinationS3GlueUpdateSerializationLibrary(r.Configuration.GlueSerializationLibrary.ValueString())
	s3BucketName := r.Configuration.S3BucketName.ValueString()
	s3BucketPath := r.Configuration.S3BucketPath.ValueString()
	s3BucketRegion := shared.DestinationS3GlueUpdateS3BucketRegion(r.Configuration.S3BucketRegion.ValueString())
	s3Endpoint := new(string)
	if !r.Configuration.S3Endpoint.IsUnknown() && !r.Configuration.S3Endpoint.IsNull() {
		*s3Endpoint = r.Configuration.S3Endpoint.ValueString()
	} else {
		s3Endpoint = nil
	}
	s3PathFormat := new(string)
	if !r.Configuration.S3PathFormat.IsUnknown() && !r.Configuration.S3PathFormat.IsNull() {
		*s3PathFormat = r.Configuration.S3PathFormat.ValueString()
	} else {
		s3PathFormat = nil
	}
	secretAccessKey := new(string)
	if !r.Configuration.SecretAccessKey.IsUnknown() && !r.Configuration.SecretAccessKey.IsNull() {
		*secretAccessKey = r.Configuration.SecretAccessKey.ValueString()
	} else {
		secretAccessKey = nil
	}
	configuration := shared.DestinationS3GlueUpdate{
		AccessKeyID:              accessKeyID,
		FileNamePattern:          fileNamePattern,
		Format:                   format,
		GlueDatabase:             glueDatabase,
		GlueSerializationLibrary: glueSerializationLibrary,
		S3BucketName:             s3BucketName,
		S3BucketPath:             s3BucketPath,
		S3BucketRegion:           s3BucketRegion,
		S3Endpoint:               s3Endpoint,
		S3PathFormat:             s3PathFormat,
		SecretAccessKey:          secretAccessKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationS3GluePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationS3GlueResourceModel) ToDeleteSDKType() *shared.DestinationS3GlueCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationS3GlueResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationS3GlueResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}

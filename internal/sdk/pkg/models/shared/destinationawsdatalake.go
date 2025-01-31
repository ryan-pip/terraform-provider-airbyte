// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitle - Name of the credentials
type DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitle string

const (
	DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitleIamUser DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitle = "IAM User"
)

func (e DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitle) ToPointer() *DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitle {
	return &e
}

func (e *DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IAM User":
		*e = DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitle: %v", v)
	}
}

// DestinationAwsDatalakeAuthenticationModeIAMUser - Choose How to Authenticate to AWS.
type DestinationAwsDatalakeAuthenticationModeIAMUser struct {
	// AWS User Access Key Id
	AwsAccessKeyID string `json:"aws_access_key_id"`
	// Secret Access Key
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	// Name of the credentials
	CredentialsTitle DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitle `json:"credentials_title"`
}

// DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitle - Name of the credentials
type DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitle string

const (
	DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitleIamRole DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitle = "IAM Role"
)

func (e DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitle) ToPointer() *DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitle {
	return &e
}

func (e *DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IAM Role":
		*e = DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitle: %v", v)
	}
}

// DestinationAwsDatalakeAuthenticationModeIAMRole - Choose How to Authenticate to AWS.
type DestinationAwsDatalakeAuthenticationModeIAMRole struct {
	// Name of the credentials
	CredentialsTitle DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitle `json:"credentials_title"`
	// Will assume this role to write data to s3
	RoleArn string `json:"role_arn"`
}

type DestinationAwsDatalakeAuthenticationModeType string

const (
	DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMRole DestinationAwsDatalakeAuthenticationModeType = "destination-aws-datalake_Authentication mode_IAM Role"
	DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMUser DestinationAwsDatalakeAuthenticationModeType = "destination-aws-datalake_Authentication mode_IAM User"
)

type DestinationAwsDatalakeAuthenticationMode struct {
	DestinationAwsDatalakeAuthenticationModeIAMRole *DestinationAwsDatalakeAuthenticationModeIAMRole
	DestinationAwsDatalakeAuthenticationModeIAMUser *DestinationAwsDatalakeAuthenticationModeIAMUser

	Type DestinationAwsDatalakeAuthenticationModeType
}

func CreateDestinationAwsDatalakeAuthenticationModeDestinationAwsDatalakeAuthenticationModeIAMRole(destinationAwsDatalakeAuthenticationModeIAMRole DestinationAwsDatalakeAuthenticationModeIAMRole) DestinationAwsDatalakeAuthenticationMode {
	typ := DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMRole

	return DestinationAwsDatalakeAuthenticationMode{
		DestinationAwsDatalakeAuthenticationModeIAMRole: &destinationAwsDatalakeAuthenticationModeIAMRole,
		Type: typ,
	}
}

func CreateDestinationAwsDatalakeAuthenticationModeDestinationAwsDatalakeAuthenticationModeIAMUser(destinationAwsDatalakeAuthenticationModeIAMUser DestinationAwsDatalakeAuthenticationModeIAMUser) DestinationAwsDatalakeAuthenticationMode {
	typ := DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMUser

	return DestinationAwsDatalakeAuthenticationMode{
		DestinationAwsDatalakeAuthenticationModeIAMUser: &destinationAwsDatalakeAuthenticationModeIAMUser,
		Type: typ,
	}
}

func (u *DestinationAwsDatalakeAuthenticationMode) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationAwsDatalakeAuthenticationModeIAMRole := new(DestinationAwsDatalakeAuthenticationModeIAMRole)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationAwsDatalakeAuthenticationModeIAMRole); err == nil {
		u.DestinationAwsDatalakeAuthenticationModeIAMRole = destinationAwsDatalakeAuthenticationModeIAMRole
		u.Type = DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMRole
		return nil
	}

	destinationAwsDatalakeAuthenticationModeIAMUser := new(DestinationAwsDatalakeAuthenticationModeIAMUser)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationAwsDatalakeAuthenticationModeIAMUser); err == nil {
		u.DestinationAwsDatalakeAuthenticationModeIAMUser = destinationAwsDatalakeAuthenticationModeIAMUser
		u.Type = DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMUser
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationAwsDatalakeAuthenticationMode) MarshalJSON() ([]byte, error) {
	if u.DestinationAwsDatalakeAuthenticationModeIAMRole != nil {
		return json.Marshal(u.DestinationAwsDatalakeAuthenticationModeIAMRole)
	}

	if u.DestinationAwsDatalakeAuthenticationModeIAMUser != nil {
		return json.Marshal(u.DestinationAwsDatalakeAuthenticationModeIAMUser)
	}

	return nil, nil
}

type DestinationAwsDatalakeAwsDatalake string

const (
	DestinationAwsDatalakeAwsDatalakeAwsDatalake DestinationAwsDatalakeAwsDatalake = "aws-datalake"
)

func (e DestinationAwsDatalakeAwsDatalake) ToPointer() *DestinationAwsDatalakeAwsDatalake {
	return &e
}

func (e *DestinationAwsDatalakeAwsDatalake) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aws-datalake":
		*e = DestinationAwsDatalakeAwsDatalake(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeAwsDatalake: %v", v)
	}
}

// DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional - The compression algorithm used to compress data.
type DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional string

const (
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalUncompressed DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional = "UNCOMPRESSED"
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalSnappy       DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional = "SNAPPY"
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalGzip         DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional = "GZIP"
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalZstd         DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional = "ZSTD"
)

func (e DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional) ToPointer() *DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional {
	return &e
}

func (e *DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNCOMPRESSED":
		fallthrough
	case "SNAPPY":
		fallthrough
	case "GZIP":
		fallthrough
	case "ZSTD":
		*e = DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional: %v", v)
	}
}

type DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcard string

const (
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcardParquet DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcard = "Parquet"
)

func (e DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcard) ToPointer() *DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcard {
	return &e
}

func (e *DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcard) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Parquet":
		*e = DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcard(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcard: %v", v)
	}
}

// DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage - Format of the data output.
type DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage struct {
	// The compression algorithm used to compress data.
	CompressionCodec *DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptional `json:"compression_codec,omitempty"`
	FormatType       DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcard        `json:"format_type"`
}

// DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptional - The compression algorithm used to compress data.
type DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptional string

const (
	DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalUncompressed DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptional = "UNCOMPRESSED"
	DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalGzip         DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptional = "GZIP"
)

func (e DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptional) ToPointer() *DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptional {
	return &e
}

func (e *DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptional) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNCOMPRESSED":
		fallthrough
	case "GZIP":
		*e = DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptional(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptional: %v", v)
	}
}

type DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcard string

const (
	DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcardJsonl DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcard = "JSONL"
)

func (e DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcard) ToPointer() *DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcard {
	return &e
}

func (e *DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcard) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSONL":
		*e = DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcard(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcard: %v", v)
	}
}

// DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON - Format of the data output.
type DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON struct {
	// The compression algorithm used to compress data.
	CompressionCodec *DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptional `json:"compression_codec,omitempty"`
	FormatType       DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcard        `json:"format_type"`
}

type DestinationAwsDatalakeOutputFormatWildcardType string

const (
	DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON DestinationAwsDatalakeOutputFormatWildcardType = "destination-aws-datalake_Output Format *_JSON Lines: Newline-delimited JSON"
	DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage        DestinationAwsDatalakeOutputFormatWildcardType = "destination-aws-datalake_Output Format *_Parquet: Columnar Storage"
)

type DestinationAwsDatalakeOutputFormatWildcard struct {
	DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON *DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage        *DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage

	Type DestinationAwsDatalakeOutputFormatWildcardType
}

func CreateDestinationAwsDatalakeOutputFormatWildcardDestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON(destinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON) DestinationAwsDatalakeOutputFormatWildcard {
	typ := DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON

	return DestinationAwsDatalakeOutputFormatWildcard{
		DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON: &destinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON,
		Type: typ,
	}
}

func CreateDestinationAwsDatalakeOutputFormatWildcardDestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage(destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage) DestinationAwsDatalakeOutputFormatWildcard {
	typ := DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage

	return DestinationAwsDatalakeOutputFormatWildcard{
		DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage: &destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage,
		Type: typ,
	}
}

func (u *DestinationAwsDatalakeOutputFormatWildcard) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON := new(DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON); err == nil {
		u.DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON = destinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON
		u.Type = DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON
		return nil
	}

	destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage := new(DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage); err == nil {
		u.DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage = destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage
		u.Type = DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationAwsDatalakeOutputFormatWildcard) MarshalJSON() ([]byte, error) {
	if u.DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON != nil {
		return json.Marshal(u.DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON)
	}

	if u.DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage != nil {
		return json.Marshal(u.DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage)
	}

	return nil, nil
}

// DestinationAwsDatalakeChooseHowToPartitionData - Partition data by cursor fields when a cursor field is a date
type DestinationAwsDatalakeChooseHowToPartitionData string

const (
	DestinationAwsDatalakeChooseHowToPartitionDataNoPartitioning DestinationAwsDatalakeChooseHowToPartitionData = "NO PARTITIONING"
	DestinationAwsDatalakeChooseHowToPartitionDataDate           DestinationAwsDatalakeChooseHowToPartitionData = "DATE"
	DestinationAwsDatalakeChooseHowToPartitionDataYear           DestinationAwsDatalakeChooseHowToPartitionData = "YEAR"
	DestinationAwsDatalakeChooseHowToPartitionDataMonth          DestinationAwsDatalakeChooseHowToPartitionData = "MONTH"
	DestinationAwsDatalakeChooseHowToPartitionDataDay            DestinationAwsDatalakeChooseHowToPartitionData = "DAY"
	DestinationAwsDatalakeChooseHowToPartitionDataYearMonth      DestinationAwsDatalakeChooseHowToPartitionData = "YEAR/MONTH"
	DestinationAwsDatalakeChooseHowToPartitionDataYearMonthDay   DestinationAwsDatalakeChooseHowToPartitionData = "YEAR/MONTH/DAY"
)

func (e DestinationAwsDatalakeChooseHowToPartitionData) ToPointer() *DestinationAwsDatalakeChooseHowToPartitionData {
	return &e
}

func (e *DestinationAwsDatalakeChooseHowToPartitionData) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO PARTITIONING":
		fallthrough
	case "DATE":
		fallthrough
	case "YEAR":
		fallthrough
	case "MONTH":
		fallthrough
	case "DAY":
		fallthrough
	case "YEAR/MONTH":
		fallthrough
	case "YEAR/MONTH/DAY":
		*e = DestinationAwsDatalakeChooseHowToPartitionData(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeChooseHowToPartitionData: %v", v)
	}
}

// DestinationAwsDatalakeS3BucketRegion - The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
type DestinationAwsDatalakeS3BucketRegion string

const (
	DestinationAwsDatalakeS3BucketRegionUnknown      DestinationAwsDatalakeS3BucketRegion = ""
	DestinationAwsDatalakeS3BucketRegionUsEast1      DestinationAwsDatalakeS3BucketRegion = "us-east-1"
	DestinationAwsDatalakeS3BucketRegionUsEast2      DestinationAwsDatalakeS3BucketRegion = "us-east-2"
	DestinationAwsDatalakeS3BucketRegionUsWest1      DestinationAwsDatalakeS3BucketRegion = "us-west-1"
	DestinationAwsDatalakeS3BucketRegionUsWest2      DestinationAwsDatalakeS3BucketRegion = "us-west-2"
	DestinationAwsDatalakeS3BucketRegionAfSouth1     DestinationAwsDatalakeS3BucketRegion = "af-south-1"
	DestinationAwsDatalakeS3BucketRegionApEast1      DestinationAwsDatalakeS3BucketRegion = "ap-east-1"
	DestinationAwsDatalakeS3BucketRegionApSouth1     DestinationAwsDatalakeS3BucketRegion = "ap-south-1"
	DestinationAwsDatalakeS3BucketRegionApNortheast1 DestinationAwsDatalakeS3BucketRegion = "ap-northeast-1"
	DestinationAwsDatalakeS3BucketRegionApNortheast2 DestinationAwsDatalakeS3BucketRegion = "ap-northeast-2"
	DestinationAwsDatalakeS3BucketRegionApNortheast3 DestinationAwsDatalakeS3BucketRegion = "ap-northeast-3"
	DestinationAwsDatalakeS3BucketRegionApSoutheast1 DestinationAwsDatalakeS3BucketRegion = "ap-southeast-1"
	DestinationAwsDatalakeS3BucketRegionApSoutheast2 DestinationAwsDatalakeS3BucketRegion = "ap-southeast-2"
	DestinationAwsDatalakeS3BucketRegionCaCentral1   DestinationAwsDatalakeS3BucketRegion = "ca-central-1"
	DestinationAwsDatalakeS3BucketRegionCnNorth1     DestinationAwsDatalakeS3BucketRegion = "cn-north-1"
	DestinationAwsDatalakeS3BucketRegionCnNorthwest1 DestinationAwsDatalakeS3BucketRegion = "cn-northwest-1"
	DestinationAwsDatalakeS3BucketRegionEuCentral1   DestinationAwsDatalakeS3BucketRegion = "eu-central-1"
	DestinationAwsDatalakeS3BucketRegionEuNorth1     DestinationAwsDatalakeS3BucketRegion = "eu-north-1"
	DestinationAwsDatalakeS3BucketRegionEuSouth1     DestinationAwsDatalakeS3BucketRegion = "eu-south-1"
	DestinationAwsDatalakeS3BucketRegionEuWest1      DestinationAwsDatalakeS3BucketRegion = "eu-west-1"
	DestinationAwsDatalakeS3BucketRegionEuWest2      DestinationAwsDatalakeS3BucketRegion = "eu-west-2"
	DestinationAwsDatalakeS3BucketRegionEuWest3      DestinationAwsDatalakeS3BucketRegion = "eu-west-3"
	DestinationAwsDatalakeS3BucketRegionSaEast1      DestinationAwsDatalakeS3BucketRegion = "sa-east-1"
	DestinationAwsDatalakeS3BucketRegionMeSouth1     DestinationAwsDatalakeS3BucketRegion = "me-south-1"
	DestinationAwsDatalakeS3BucketRegionUsGovEast1   DestinationAwsDatalakeS3BucketRegion = "us-gov-east-1"
	DestinationAwsDatalakeS3BucketRegionUsGovWest1   DestinationAwsDatalakeS3BucketRegion = "us-gov-west-1"
)

func (e DestinationAwsDatalakeS3BucketRegion) ToPointer() *DestinationAwsDatalakeS3BucketRegion {
	return &e
}

func (e *DestinationAwsDatalakeS3BucketRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		*e = DestinationAwsDatalakeS3BucketRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeS3BucketRegion: %v", v)
	}
}

type DestinationAwsDatalake struct {
	// target aws account id
	AwsAccountID *string `json:"aws_account_id,omitempty"`
	// The name of the S3 bucket. Read more <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html">here</a>.
	BucketName string `json:"bucket_name"`
	// S3 prefix
	BucketPrefix *string `json:"bucket_prefix,omitempty"`
	// Choose How to Authenticate to AWS.
	Credentials     DestinationAwsDatalakeAuthenticationMode `json:"credentials"`
	DestinationType DestinationAwsDatalakeAwsDatalake        `json:"destinationType"`
	// Format of the data output.
	Format *DestinationAwsDatalakeOutputFormatWildcard `json:"format,omitempty"`
	// Cast float/double as decimal(38,18). This can help achieve higher accuracy and represent numbers correctly as received from the source.
	GlueCatalogFloatAsDecimal *bool `json:"glue_catalog_float_as_decimal,omitempty"`
	// Add a default tag key to databases created by this destination
	LakeformationDatabaseDefaultTagKey *string `json:"lakeformation_database_default_tag_key,omitempty"`
	// Add default values for the `Tag Key` to databases created by this destination. Comma separate for multiple values.
	LakeformationDatabaseDefaultTagValues *string `json:"lakeformation_database_default_tag_values,omitempty"`
	// The default database this destination will use to create tables in per stream. Can be changed per connection by customizing the namespace.
	LakeformationDatabaseName string `json:"lakeformation_database_name"`
	// Whether to create tables as LF governed tables.
	LakeformationGovernedTables *bool `json:"lakeformation_governed_tables,omitempty"`
	// Partition data by cursor fields when a cursor field is a date
	Partitioning *DestinationAwsDatalakeChooseHowToPartitionData `json:"partitioning,omitempty"`
	// The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
	Region DestinationAwsDatalakeS3BucketRegion `json:"region"`
}

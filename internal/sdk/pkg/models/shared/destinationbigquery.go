// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// DestinationBigqueryDatasetLocation - The location of the dataset. Warning: Changes made after creation will not be applied. Read more <a href="https://cloud.google.com/bigquery/docs/locations">here</a>.
type DestinationBigqueryDatasetLocation string

const (
	DestinationBigqueryDatasetLocationUs                     DestinationBigqueryDatasetLocation = "US"
	DestinationBigqueryDatasetLocationEu                     DestinationBigqueryDatasetLocation = "EU"
	DestinationBigqueryDatasetLocationAsiaEast1              DestinationBigqueryDatasetLocation = "asia-east1"
	DestinationBigqueryDatasetLocationAsiaEast2              DestinationBigqueryDatasetLocation = "asia-east2"
	DestinationBigqueryDatasetLocationAsiaNortheast1         DestinationBigqueryDatasetLocation = "asia-northeast1"
	DestinationBigqueryDatasetLocationAsiaNortheast2         DestinationBigqueryDatasetLocation = "asia-northeast2"
	DestinationBigqueryDatasetLocationAsiaNortheast3         DestinationBigqueryDatasetLocation = "asia-northeast3"
	DestinationBigqueryDatasetLocationAsiaSouth1             DestinationBigqueryDatasetLocation = "asia-south1"
	DestinationBigqueryDatasetLocationAsiaSouth2             DestinationBigqueryDatasetLocation = "asia-south2"
	DestinationBigqueryDatasetLocationAsiaSoutheast1         DestinationBigqueryDatasetLocation = "asia-southeast1"
	DestinationBigqueryDatasetLocationAsiaSoutheast2         DestinationBigqueryDatasetLocation = "asia-southeast2"
	DestinationBigqueryDatasetLocationAustraliaSoutheast1    DestinationBigqueryDatasetLocation = "australia-southeast1"
	DestinationBigqueryDatasetLocationAustraliaSoutheast2    DestinationBigqueryDatasetLocation = "australia-southeast2"
	DestinationBigqueryDatasetLocationEuropeCentral1         DestinationBigqueryDatasetLocation = "europe-central1"
	DestinationBigqueryDatasetLocationEuropeCentral2         DestinationBigqueryDatasetLocation = "europe-central2"
	DestinationBigqueryDatasetLocationEuropeNorth1           DestinationBigqueryDatasetLocation = "europe-north1"
	DestinationBigqueryDatasetLocationEuropeSouthwest1       DestinationBigqueryDatasetLocation = "europe-southwest1"
	DestinationBigqueryDatasetLocationEuropeWest1            DestinationBigqueryDatasetLocation = "europe-west1"
	DestinationBigqueryDatasetLocationEuropeWest2            DestinationBigqueryDatasetLocation = "europe-west2"
	DestinationBigqueryDatasetLocationEuropeWest3            DestinationBigqueryDatasetLocation = "europe-west3"
	DestinationBigqueryDatasetLocationEuropeWest4            DestinationBigqueryDatasetLocation = "europe-west4"
	DestinationBigqueryDatasetLocationEuropeWest6            DestinationBigqueryDatasetLocation = "europe-west6"
	DestinationBigqueryDatasetLocationEuropeWest7            DestinationBigqueryDatasetLocation = "europe-west7"
	DestinationBigqueryDatasetLocationEuropeWest8            DestinationBigqueryDatasetLocation = "europe-west8"
	DestinationBigqueryDatasetLocationEuropeWest9            DestinationBigqueryDatasetLocation = "europe-west9"
	DestinationBigqueryDatasetLocationMeWest1                DestinationBigqueryDatasetLocation = "me-west1"
	DestinationBigqueryDatasetLocationNorthamericaNortheast1 DestinationBigqueryDatasetLocation = "northamerica-northeast1"
	DestinationBigqueryDatasetLocationNorthamericaNortheast2 DestinationBigqueryDatasetLocation = "northamerica-northeast2"
	DestinationBigqueryDatasetLocationSouthamericaEast1      DestinationBigqueryDatasetLocation = "southamerica-east1"
	DestinationBigqueryDatasetLocationSouthamericaWest1      DestinationBigqueryDatasetLocation = "southamerica-west1"
	DestinationBigqueryDatasetLocationUsCentral1             DestinationBigqueryDatasetLocation = "us-central1"
	DestinationBigqueryDatasetLocationUsEast1                DestinationBigqueryDatasetLocation = "us-east1"
	DestinationBigqueryDatasetLocationUsEast2                DestinationBigqueryDatasetLocation = "us-east2"
	DestinationBigqueryDatasetLocationUsEast3                DestinationBigqueryDatasetLocation = "us-east3"
	DestinationBigqueryDatasetLocationUsEast4                DestinationBigqueryDatasetLocation = "us-east4"
	DestinationBigqueryDatasetLocationUsEast5                DestinationBigqueryDatasetLocation = "us-east5"
	DestinationBigqueryDatasetLocationUsWest1                DestinationBigqueryDatasetLocation = "us-west1"
	DestinationBigqueryDatasetLocationUsWest2                DestinationBigqueryDatasetLocation = "us-west2"
	DestinationBigqueryDatasetLocationUsWest3                DestinationBigqueryDatasetLocation = "us-west3"
	DestinationBigqueryDatasetLocationUsWest4                DestinationBigqueryDatasetLocation = "us-west4"
)

func (e DestinationBigqueryDatasetLocation) ToPointer() *DestinationBigqueryDatasetLocation {
	return &e
}

func (e *DestinationBigqueryDatasetLocation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		fallthrough
	case "asia-east1":
		fallthrough
	case "asia-east2":
		fallthrough
	case "asia-northeast1":
		fallthrough
	case "asia-northeast2":
		fallthrough
	case "asia-northeast3":
		fallthrough
	case "asia-south1":
		fallthrough
	case "asia-south2":
		fallthrough
	case "asia-southeast1":
		fallthrough
	case "asia-southeast2":
		fallthrough
	case "australia-southeast1":
		fallthrough
	case "australia-southeast2":
		fallthrough
	case "europe-central1":
		fallthrough
	case "europe-central2":
		fallthrough
	case "europe-north1":
		fallthrough
	case "europe-southwest1":
		fallthrough
	case "europe-west1":
		fallthrough
	case "europe-west2":
		fallthrough
	case "europe-west3":
		fallthrough
	case "europe-west4":
		fallthrough
	case "europe-west6":
		fallthrough
	case "europe-west7":
		fallthrough
	case "europe-west8":
		fallthrough
	case "europe-west9":
		fallthrough
	case "me-west1":
		fallthrough
	case "northamerica-northeast1":
		fallthrough
	case "northamerica-northeast2":
		fallthrough
	case "southamerica-east1":
		fallthrough
	case "southamerica-west1":
		fallthrough
	case "us-central1":
		fallthrough
	case "us-east1":
		fallthrough
	case "us-east2":
		fallthrough
	case "us-east3":
		fallthrough
	case "us-east4":
		fallthrough
	case "us-east5":
		fallthrough
	case "us-west1":
		fallthrough
	case "us-west2":
		fallthrough
	case "us-west3":
		fallthrough
	case "us-west4":
		*e = DestinationBigqueryDatasetLocation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryDatasetLocation: %v", v)
	}
}

type DestinationBigqueryBigquery string

const (
	DestinationBigqueryBigqueryBigquery DestinationBigqueryBigquery = "bigquery"
)

func (e DestinationBigqueryBigquery) ToPointer() *DestinationBigqueryBigquery {
	return &e
}

func (e *DestinationBigqueryBigquery) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bigquery":
		*e = DestinationBigqueryBigquery(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryBigquery: %v", v)
	}
}

type DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKeyCredentialType string

const (
	DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKeyCredentialTypeHmacKey DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKeyCredentialType = "HMAC_KEY"
)

func (e DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKeyCredentialType) ToPointer() *DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKeyCredentialType {
	return &e
}

func (e *DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKeyCredentialType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HMAC_KEY":
		*e = DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKeyCredentialType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKeyCredentialType: %v", v)
	}
}

// DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey - An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.
type DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey struct {
	CredentialType DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKeyCredentialType `json:"credential_type"`
	// HMAC key access ID. When linked to a service account, this ID is 61 characters long; when linked to a user account, it is 24 characters long.
	HmacKeyAccessID string `json:"hmac_key_access_id"`
	// The corresponding secret for the access ID. It is a 40-character base-64 encoded string.
	HmacKeySecret string `json:"hmac_key_secret"`
}

type DestinationBigqueryLoadingMethodGCSStagingCredentialType string

const (
	DestinationBigqueryLoadingMethodGCSStagingCredentialTypeDestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey DestinationBigqueryLoadingMethodGCSStagingCredentialType = "destination-bigquery_Loading Method_GCS Staging_Credential_HMAC key"
)

type DestinationBigqueryLoadingMethodGCSStagingCredential struct {
	DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey *DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey

	Type DestinationBigqueryLoadingMethodGCSStagingCredentialType
}

func CreateDestinationBigqueryLoadingMethodGCSStagingCredentialDestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey(destinationBigqueryLoadingMethodGCSStagingCredentialHMACKey DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey) DestinationBigqueryLoadingMethodGCSStagingCredential {
	typ := DestinationBigqueryLoadingMethodGCSStagingCredentialTypeDestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey

	return DestinationBigqueryLoadingMethodGCSStagingCredential{
		DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey: &destinationBigqueryLoadingMethodGCSStagingCredentialHMACKey,
		Type: typ,
	}
}

func (u *DestinationBigqueryLoadingMethodGCSStagingCredential) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationBigqueryLoadingMethodGCSStagingCredentialHMACKey := new(DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationBigqueryLoadingMethodGCSStagingCredentialHMACKey); err == nil {
		u.DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey = destinationBigqueryLoadingMethodGCSStagingCredentialHMACKey
		u.Type = DestinationBigqueryLoadingMethodGCSStagingCredentialTypeDestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationBigqueryLoadingMethodGCSStagingCredential) MarshalJSON() ([]byte, error) {
	if u.DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey != nil {
		return json.Marshal(u.DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey)
	}

	return nil, nil
}

// DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing - This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.
type DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing string

const (
	DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessingDeleteAllTmpFilesFromGcs DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing = "Delete all tmp files from GCS"
	DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessingKeepAllTmpFilesInGcs     DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing = "Keep all tmp files in GCS"
)

func (e DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing) ToPointer() *DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing {
	return &e
}

func (e *DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Delete all tmp files from GCS":
		fallthrough
	case "Keep all tmp files in GCS":
		*e = DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing: %v", v)
	}
}

type DestinationBigqueryLoadingMethodGCSStagingMethod string

const (
	DestinationBigqueryLoadingMethodGCSStagingMethodGcsStaging DestinationBigqueryLoadingMethodGCSStagingMethod = "GCS Staging"
)

func (e DestinationBigqueryLoadingMethodGCSStagingMethod) ToPointer() *DestinationBigqueryLoadingMethodGCSStagingMethod {
	return &e
}

func (e *DestinationBigqueryLoadingMethodGCSStagingMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GCS Staging":
		*e = DestinationBigqueryLoadingMethodGCSStagingMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryLoadingMethodGCSStagingMethod: %v", v)
	}
}

// DestinationBigqueryLoadingMethodGCSStaging - Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.
type DestinationBigqueryLoadingMethodGCSStaging struct {
	// An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.
	Credential DestinationBigqueryLoadingMethodGCSStagingCredential `json:"credential"`
	// Number of file buffers allocated for writing data. Increasing this number is beneficial for connections using Change Data Capture (CDC) and up to the number of streams within a connection. Increasing the number of file buffers past the maximum number of streams has deteriorating effects
	FileBufferCount *int64 `json:"file_buffer_count,omitempty"`
	// The name of the GCS bucket. Read more <a href="https://cloud.google.com/storage/docs/naming-buckets">here</a>.
	GcsBucketName string `json:"gcs_bucket_name"`
	// Directory under the GCS bucket where data will be written.
	GcsBucketPath string `json:"gcs_bucket_path"`
	// This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.
	KeepFilesInGcsBucket *DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing `json:"keep_files_in_gcs-bucket,omitempty"`
	Method               DestinationBigqueryLoadingMethodGCSStagingMethod                          `json:"method"`
}

type DestinationBigqueryLoadingMethodStandardInsertsMethod string

const (
	DestinationBigqueryLoadingMethodStandardInsertsMethodStandard DestinationBigqueryLoadingMethodStandardInsertsMethod = "Standard"
)

func (e DestinationBigqueryLoadingMethodStandardInsertsMethod) ToPointer() *DestinationBigqueryLoadingMethodStandardInsertsMethod {
	return &e
}

func (e *DestinationBigqueryLoadingMethodStandardInsertsMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Standard":
		*e = DestinationBigqueryLoadingMethodStandardInsertsMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryLoadingMethodStandardInsertsMethod: %v", v)
	}
}

// DestinationBigqueryLoadingMethodStandardInserts - Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.
type DestinationBigqueryLoadingMethodStandardInserts struct {
	Method DestinationBigqueryLoadingMethodStandardInsertsMethod `json:"method"`
}

type DestinationBigqueryLoadingMethodType string

const (
	DestinationBigqueryLoadingMethodTypeDestinationBigqueryLoadingMethodStandardInserts DestinationBigqueryLoadingMethodType = "destination-bigquery_Loading Method_Standard Inserts"
	DestinationBigqueryLoadingMethodTypeDestinationBigqueryLoadingMethodGCSStaging      DestinationBigqueryLoadingMethodType = "destination-bigquery_Loading Method_GCS Staging"
)

type DestinationBigqueryLoadingMethod struct {
	DestinationBigqueryLoadingMethodStandardInserts *DestinationBigqueryLoadingMethodStandardInserts
	DestinationBigqueryLoadingMethodGCSStaging      *DestinationBigqueryLoadingMethodGCSStaging

	Type DestinationBigqueryLoadingMethodType
}

func CreateDestinationBigqueryLoadingMethodDestinationBigqueryLoadingMethodStandardInserts(destinationBigqueryLoadingMethodStandardInserts DestinationBigqueryLoadingMethodStandardInserts) DestinationBigqueryLoadingMethod {
	typ := DestinationBigqueryLoadingMethodTypeDestinationBigqueryLoadingMethodStandardInserts

	return DestinationBigqueryLoadingMethod{
		DestinationBigqueryLoadingMethodStandardInserts: &destinationBigqueryLoadingMethodStandardInserts,
		Type: typ,
	}
}

func CreateDestinationBigqueryLoadingMethodDestinationBigqueryLoadingMethodGCSStaging(destinationBigqueryLoadingMethodGCSStaging DestinationBigqueryLoadingMethodGCSStaging) DestinationBigqueryLoadingMethod {
	typ := DestinationBigqueryLoadingMethodTypeDestinationBigqueryLoadingMethodGCSStaging

	return DestinationBigqueryLoadingMethod{
		DestinationBigqueryLoadingMethodGCSStaging: &destinationBigqueryLoadingMethodGCSStaging,
		Type: typ,
	}
}

func (u *DestinationBigqueryLoadingMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationBigqueryLoadingMethodStandardInserts := new(DestinationBigqueryLoadingMethodStandardInserts)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationBigqueryLoadingMethodStandardInserts); err == nil {
		u.DestinationBigqueryLoadingMethodStandardInserts = destinationBigqueryLoadingMethodStandardInserts
		u.Type = DestinationBigqueryLoadingMethodTypeDestinationBigqueryLoadingMethodStandardInserts
		return nil
	}

	destinationBigqueryLoadingMethodGCSStaging := new(DestinationBigqueryLoadingMethodGCSStaging)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationBigqueryLoadingMethodGCSStaging); err == nil {
		u.DestinationBigqueryLoadingMethodGCSStaging = destinationBigqueryLoadingMethodGCSStaging
		u.Type = DestinationBigqueryLoadingMethodTypeDestinationBigqueryLoadingMethodGCSStaging
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationBigqueryLoadingMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationBigqueryLoadingMethodStandardInserts != nil {
		return json.Marshal(u.DestinationBigqueryLoadingMethodStandardInserts)
	}

	if u.DestinationBigqueryLoadingMethodGCSStaging != nil {
		return json.Marshal(u.DestinationBigqueryLoadingMethodGCSStaging)
	}

	return nil, nil
}

// DestinationBigqueryTransformationQueryRunType - Interactive run type means that the query is executed as soon as possible, and these queries count towards concurrent rate limit and daily limit. Read more about interactive run type <a href="https://cloud.google.com/bigquery/docs/running-queries#queries">here</a>. Batch queries are queued and started as soon as idle resources are available in the BigQuery shared resource pool, which usually occurs within a few minutes. Batch queries don’t count towards your concurrent rate limit. Read more about batch queries <a href="https://cloud.google.com/bigquery/docs/running-queries#batch">here</a>. The default "interactive" value is used if not set explicitly.
type DestinationBigqueryTransformationQueryRunType string

const (
	DestinationBigqueryTransformationQueryRunTypeInteractive DestinationBigqueryTransformationQueryRunType = "interactive"
	DestinationBigqueryTransformationQueryRunTypeBatch       DestinationBigqueryTransformationQueryRunType = "batch"
)

func (e DestinationBigqueryTransformationQueryRunType) ToPointer() *DestinationBigqueryTransformationQueryRunType {
	return &e
}

func (e *DestinationBigqueryTransformationQueryRunType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "interactive":
		fallthrough
	case "batch":
		*e = DestinationBigqueryTransformationQueryRunType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryTransformationQueryRunType: %v", v)
	}
}

type DestinationBigquery struct {
	// Google BigQuery client's chunk (buffer) size (MIN=1, MAX = 15) for each table. The size that will be written by a single RPC. Written data will be buffered and only flushed upon reaching this size or closing the channel. The default 15MB value is used if not set explicitly. Read more <a href="https://googleapis.dev/python/bigquery/latest/generated/google.cloud.bigquery.client.Client.html">here</a>.
	BigQueryClientBufferSizeMb *int64 `json:"big_query_client_buffer_size_mb,omitempty"`
	// The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.com/integrations/destinations/bigquery#service-account-key">docs</a> if you need help generating this key. Default credentials will be used if this field is left empty.
	CredentialsJSON *string `json:"credentials_json,omitempty"`
	// The default BigQuery Dataset ID that tables are replicated to if the source does not specify a namespace. Read more <a href="https://cloud.google.com/bigquery/docs/datasets#create-dataset">here</a>.
	DatasetID string `json:"dataset_id"`
	// The location of the dataset. Warning: Changes made after creation will not be applied. Read more <a href="https://cloud.google.com/bigquery/docs/locations">here</a>.
	DatasetLocation DestinationBigqueryDatasetLocation `json:"dataset_location"`
	DestinationType DestinationBigqueryBigquery        `json:"destinationType"`
	// Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.
	LoadingMethod *DestinationBigqueryLoadingMethod `json:"loading_method,omitempty"`
	// The GCP project ID for the project containing the target BigQuery dataset. Read more <a href="https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects">here</a>.
	ProjectID string `json:"project_id"`
	// (Early Access) The dataset to write raw tables into
	RawDataDataset *string `json:"raw_data_dataset,omitempty"`
	// Interactive run type means that the query is executed as soon as possible, and these queries count towards concurrent rate limit and daily limit. Read more about interactive run type <a href="https://cloud.google.com/bigquery/docs/running-queries#queries">here</a>. Batch queries are queued and started as soon as idle resources are available in the BigQuery shared resource pool, which usually occurs within a few minutes. Batch queries don’t count towards your concurrent rate limit. Read more about batch queries <a href="https://cloud.google.com/bigquery/docs/running-queries#batch">here</a>. The default "interactive" value is used if not set explicitly.
	TransformationPriority *DestinationBigqueryTransformationQueryRunType `json:"transformation_priority,omitempty"`
	// (Early Access) Use <a href="https://docs.airbyte.com/understanding-airbyte/typing-deduping" target="_blank">Destinations V2</a>.
	Use1s1tFormat *bool `json:"use_1s1t_format,omitempty"`
}

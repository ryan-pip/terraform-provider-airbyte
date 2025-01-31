// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceFileSecureUpdateFileFormat - The Format of the file which should be replicated (Warning: some formats may be experimental, please refer to the docs).
type SourceFileSecureUpdateFileFormat string

const (
	SourceFileSecureUpdateFileFormatCsv         SourceFileSecureUpdateFileFormat = "csv"
	SourceFileSecureUpdateFileFormatJSON        SourceFileSecureUpdateFileFormat = "json"
	SourceFileSecureUpdateFileFormatJsonl       SourceFileSecureUpdateFileFormat = "jsonl"
	SourceFileSecureUpdateFileFormatExcel       SourceFileSecureUpdateFileFormat = "excel"
	SourceFileSecureUpdateFileFormatExcelBinary SourceFileSecureUpdateFileFormat = "excel_binary"
	SourceFileSecureUpdateFileFormatFeather     SourceFileSecureUpdateFileFormat = "feather"
	SourceFileSecureUpdateFileFormatParquet     SourceFileSecureUpdateFileFormat = "parquet"
	SourceFileSecureUpdateFileFormatYaml        SourceFileSecureUpdateFileFormat = "yaml"
)

func (e SourceFileSecureUpdateFileFormat) ToPointer() *SourceFileSecureUpdateFileFormat {
	return &e
}

func (e *SourceFileSecureUpdateFileFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "csv":
		fallthrough
	case "json":
		fallthrough
	case "jsonl":
		fallthrough
	case "excel":
		fallthrough
	case "excel_binary":
		fallthrough
	case "feather":
		fallthrough
	case "parquet":
		fallthrough
	case "yaml":
		*e = SourceFileSecureUpdateFileFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureUpdateFileFormat: %v", v)
	}
}

type SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocolStorage string

const (
	SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocolStorageSftp SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocolStorage = "SFTP"
)

func (e SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocolStorage) ToPointer() *SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocolStorage {
	return &e
}

func (e *SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocolStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SFTP":
		*e = SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocolStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocolStorage: %v", v)
	}
}

// SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol struct {
	Host     string                                                                     `json:"host"`
	Password *string                                                                    `json:"password,omitempty"`
	Port     *string                                                                    `json:"port,omitempty"`
	Storage  SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocolStorage `json:"storage"`
	User     string                                                                     `json:"user"`
}

type SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocolStorage string

const (
	SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocolStorageScp SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocolStorage = "SCP"
)

func (e SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocolStorage) ToPointer() *SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocolStorage {
	return &e
}

func (e *SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocolStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SCP":
		*e = SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocolStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocolStorage: %v", v)
	}
}

// SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol struct {
	Host     string                                                            `json:"host"`
	Password *string                                                           `json:"password,omitempty"`
	Port     *string                                                           `json:"port,omitempty"`
	Storage  SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocolStorage `json:"storage"`
	User     string                                                            `json:"user"`
}

type SourceFileSecureUpdateStorageProviderSSHSecureShellStorage string

const (
	SourceFileSecureUpdateStorageProviderSSHSecureShellStorageSSH SourceFileSecureUpdateStorageProviderSSHSecureShellStorage = "SSH"
)

func (e SourceFileSecureUpdateStorageProviderSSHSecureShellStorage) ToPointer() *SourceFileSecureUpdateStorageProviderSSHSecureShellStorage {
	return &e
}

func (e *SourceFileSecureUpdateStorageProviderSSHSecureShellStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH":
		*e = SourceFileSecureUpdateStorageProviderSSHSecureShellStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureUpdateStorageProviderSSHSecureShellStorage: %v", v)
	}
}

// SourceFileSecureUpdateStorageProviderSSHSecureShell - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureUpdateStorageProviderSSHSecureShell struct {
	Host     string                                                     `json:"host"`
	Password *string                                                    `json:"password,omitempty"`
	Port     *string                                                    `json:"port,omitempty"`
	Storage  SourceFileSecureUpdateStorageProviderSSHSecureShellStorage `json:"storage"`
	User     string                                                     `json:"user"`
}

type SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorageStorage string

const (
	SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorageStorageAzBlob SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorageStorage = "AzBlob"
)

func (e SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorageStorage) ToPointer() *SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorageStorage {
	return &e
}

func (e *SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorageStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AzBlob":
		*e = SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorageStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorageStorage: %v", v)
	}
}

// SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage struct {
	// To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a SAS (Shared Access Signature) token. If accessing publicly available data, this field is not necessary.
	SasToken *string `json:"sas_token,omitempty"`
	// To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a storage account shared key (aka account key or access key). If accessing publicly available data, this field is not necessary.
	SharedKey *string                                                            `json:"shared_key,omitempty"`
	Storage   SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorageStorage `json:"storage"`
	// The globally unique name of the storage account that the desired blob sits within. See <a href="https://docs.microsoft.com/en-us/azure/storage/common/storage-account-overview" target="_blank">here</a> for more details.
	StorageAccount string `json:"storage_account"`
}

type SourceFileSecureUpdateStorageProviderS3AmazonWebServicesStorage string

const (
	SourceFileSecureUpdateStorageProviderS3AmazonWebServicesStorageS3 SourceFileSecureUpdateStorageProviderS3AmazonWebServicesStorage = "S3"
)

func (e SourceFileSecureUpdateStorageProviderS3AmazonWebServicesStorage) ToPointer() *SourceFileSecureUpdateStorageProviderS3AmazonWebServicesStorage {
	return &e
}

func (e *SourceFileSecureUpdateStorageProviderS3AmazonWebServicesStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3":
		*e = SourceFileSecureUpdateStorageProviderS3AmazonWebServicesStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureUpdateStorageProviderS3AmazonWebServicesStorage: %v", v)
	}
}

// SourceFileSecureUpdateStorageProviderS3AmazonWebServices - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureUpdateStorageProviderS3AmazonWebServices struct {
	// In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`
	// In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
	AwsSecretAccessKey *string                                                         `json:"aws_secret_access_key,omitempty"`
	Storage            SourceFileSecureUpdateStorageProviderS3AmazonWebServicesStorage `json:"storage"`
}

type SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorageStorage string

const (
	SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorageStorageGcs SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorageStorage = "GCS"
)

func (e SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorageStorage) ToPointer() *SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorageStorage {
	return &e
}

func (e *SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorageStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GCS":
		*e = SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorageStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorageStorage: %v", v)
	}
}

// SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage struct {
	// In order to access private Buckets stored on Google Cloud, this connector would need a service account json credentials with the proper permissions as described <a href="https://cloud.google.com/iam/docs/service-accounts" target="_blank">here</a>. Please generate the credentials.json file and copy/paste its content to this field (expecting JSON formats). If accessing publicly available data, this field is not necessary.
	ServiceAccountJSON *string                                                           `json:"service_account_json,omitempty"`
	Storage            SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorageStorage `json:"storage"`
}

type SourceFileSecureUpdateStorageProviderHTTPSPublicWebStorage string

const (
	SourceFileSecureUpdateStorageProviderHTTPSPublicWebStorageHTTPS SourceFileSecureUpdateStorageProviderHTTPSPublicWebStorage = "HTTPS"
)

func (e SourceFileSecureUpdateStorageProviderHTTPSPublicWebStorage) ToPointer() *SourceFileSecureUpdateStorageProviderHTTPSPublicWebStorage {
	return &e
}

func (e *SourceFileSecureUpdateStorageProviderHTTPSPublicWebStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HTTPS":
		*e = SourceFileSecureUpdateStorageProviderHTTPSPublicWebStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureUpdateStorageProviderHTTPSPublicWebStorage: %v", v)
	}
}

// SourceFileSecureUpdateStorageProviderHTTPSPublicWeb - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureUpdateStorageProviderHTTPSPublicWeb struct {
	Storage SourceFileSecureUpdateStorageProviderHTTPSPublicWebStorage `json:"storage"`
	// Add User-Agent to request
	UserAgent *bool `json:"user_agent,omitempty"`
}

type SourceFileSecureUpdateStorageProviderType string

const (
	SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderHTTPSPublicWeb                 SourceFileSecureUpdateStorageProviderType = "source-file-secure-update_Storage Provider_HTTPS: Public Web"
	SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage          SourceFileSecureUpdateStorageProviderType = "source-file-secure-update_Storage Provider_GCS: Google Cloud Storage"
	SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderS3AmazonWebServices            SourceFileSecureUpdateStorageProviderType = "source-file-secure-update_Storage Provider_S3: Amazon Web Services"
	SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage         SourceFileSecureUpdateStorageProviderType = "source-file-secure-update_Storage Provider_AzBlob: Azure Blob Storage"
	SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderSSHSecureShell                 SourceFileSecureUpdateStorageProviderType = "source-file-secure-update_Storage Provider_SSH: Secure Shell"
	SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol          SourceFileSecureUpdateStorageProviderType = "source-file-secure-update_Storage Provider_SCP: Secure copy protocol"
	SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol SourceFileSecureUpdateStorageProviderType = "source-file-secure-update_Storage Provider_SFTP: Secure File Transfer Protocol"
)

type SourceFileSecureUpdateStorageProvider struct {
	SourceFileSecureUpdateStorageProviderHTTPSPublicWeb                 *SourceFileSecureUpdateStorageProviderHTTPSPublicWeb
	SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage          *SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage
	SourceFileSecureUpdateStorageProviderS3AmazonWebServices            *SourceFileSecureUpdateStorageProviderS3AmazonWebServices
	SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage         *SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage
	SourceFileSecureUpdateStorageProviderSSHSecureShell                 *SourceFileSecureUpdateStorageProviderSSHSecureShell
	SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol          *SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol
	SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol *SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol

	Type SourceFileSecureUpdateStorageProviderType
}

func CreateSourceFileSecureUpdateStorageProviderSourceFileSecureUpdateStorageProviderHTTPSPublicWeb(sourceFileSecureUpdateStorageProviderHTTPSPublicWeb SourceFileSecureUpdateStorageProviderHTTPSPublicWeb) SourceFileSecureUpdateStorageProvider {
	typ := SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderHTTPSPublicWeb

	return SourceFileSecureUpdateStorageProvider{
		SourceFileSecureUpdateStorageProviderHTTPSPublicWeb: &sourceFileSecureUpdateStorageProviderHTTPSPublicWeb,
		Type: typ,
	}
}

func CreateSourceFileSecureUpdateStorageProviderSourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage(sourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage) SourceFileSecureUpdateStorageProvider {
	typ := SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage

	return SourceFileSecureUpdateStorageProvider{
		SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage: &sourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage,
		Type: typ,
	}
}

func CreateSourceFileSecureUpdateStorageProviderSourceFileSecureUpdateStorageProviderS3AmazonWebServices(sourceFileSecureUpdateStorageProviderS3AmazonWebServices SourceFileSecureUpdateStorageProviderS3AmazonWebServices) SourceFileSecureUpdateStorageProvider {
	typ := SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderS3AmazonWebServices

	return SourceFileSecureUpdateStorageProvider{
		SourceFileSecureUpdateStorageProviderS3AmazonWebServices: &sourceFileSecureUpdateStorageProviderS3AmazonWebServices,
		Type: typ,
	}
}

func CreateSourceFileSecureUpdateStorageProviderSourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage(sourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage) SourceFileSecureUpdateStorageProvider {
	typ := SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage

	return SourceFileSecureUpdateStorageProvider{
		SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage: &sourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage,
		Type: typ,
	}
}

func CreateSourceFileSecureUpdateStorageProviderSourceFileSecureUpdateStorageProviderSSHSecureShell(sourceFileSecureUpdateStorageProviderSSHSecureShell SourceFileSecureUpdateStorageProviderSSHSecureShell) SourceFileSecureUpdateStorageProvider {
	typ := SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderSSHSecureShell

	return SourceFileSecureUpdateStorageProvider{
		SourceFileSecureUpdateStorageProviderSSHSecureShell: &sourceFileSecureUpdateStorageProviderSSHSecureShell,
		Type: typ,
	}
}

func CreateSourceFileSecureUpdateStorageProviderSourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol(sourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol) SourceFileSecureUpdateStorageProvider {
	typ := SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol

	return SourceFileSecureUpdateStorageProvider{
		SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol: &sourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol,
		Type: typ,
	}
}

func CreateSourceFileSecureUpdateStorageProviderSourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol(sourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol) SourceFileSecureUpdateStorageProvider {
	typ := SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol

	return SourceFileSecureUpdateStorageProvider{
		SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol: &sourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol,
		Type: typ,
	}
}

func (u *SourceFileSecureUpdateStorageProvider) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceFileSecureUpdateStorageProviderHTTPSPublicWeb := new(SourceFileSecureUpdateStorageProviderHTTPSPublicWeb)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureUpdateStorageProviderHTTPSPublicWeb); err == nil {
		u.SourceFileSecureUpdateStorageProviderHTTPSPublicWeb = sourceFileSecureUpdateStorageProviderHTTPSPublicWeb
		u.Type = SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderHTTPSPublicWeb
		return nil
	}

	sourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage := new(SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage); err == nil {
		u.SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage = sourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage
		u.Type = SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage
		return nil
	}

	sourceFileSecureUpdateStorageProviderS3AmazonWebServices := new(SourceFileSecureUpdateStorageProviderS3AmazonWebServices)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureUpdateStorageProviderS3AmazonWebServices); err == nil {
		u.SourceFileSecureUpdateStorageProviderS3AmazonWebServices = sourceFileSecureUpdateStorageProviderS3AmazonWebServices
		u.Type = SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderS3AmazonWebServices
		return nil
	}

	sourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage := new(SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage); err == nil {
		u.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage = sourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage
		u.Type = SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage
		return nil
	}

	sourceFileSecureUpdateStorageProviderSSHSecureShell := new(SourceFileSecureUpdateStorageProviderSSHSecureShell)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureUpdateStorageProviderSSHSecureShell); err == nil {
		u.SourceFileSecureUpdateStorageProviderSSHSecureShell = sourceFileSecureUpdateStorageProviderSSHSecureShell
		u.Type = SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderSSHSecureShell
		return nil
	}

	sourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol := new(SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol); err == nil {
		u.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol = sourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol
		u.Type = SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol
		return nil
	}

	sourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol := new(SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol); err == nil {
		u.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol = sourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol
		u.Type = SourceFileSecureUpdateStorageProviderTypeSourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceFileSecureUpdateStorageProvider) MarshalJSON() ([]byte, error) {
	if u.SourceFileSecureUpdateStorageProviderHTTPSPublicWeb != nil {
		return json.Marshal(u.SourceFileSecureUpdateStorageProviderHTTPSPublicWeb)
	}

	if u.SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage != nil {
		return json.Marshal(u.SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage)
	}

	if u.SourceFileSecureUpdateStorageProviderS3AmazonWebServices != nil {
		return json.Marshal(u.SourceFileSecureUpdateStorageProviderS3AmazonWebServices)
	}

	if u.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage != nil {
		return json.Marshal(u.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage)
	}

	if u.SourceFileSecureUpdateStorageProviderSSHSecureShell != nil {
		return json.Marshal(u.SourceFileSecureUpdateStorageProviderSSHSecureShell)
	}

	if u.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol != nil {
		return json.Marshal(u.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol)
	}

	if u.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol != nil {
		return json.Marshal(u.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol)
	}

	return nil, nil
}

type SourceFileSecureUpdate struct {
	// The Name of the final table to replicate this file into (should include letters, numbers dash and underscores only).
	DatasetName string `json:"dataset_name"`
	// The Format of the file which should be replicated (Warning: some formats may be experimental, please refer to the docs).
	Format SourceFileSecureUpdateFileFormat `json:"format"`
	// The storage Provider or Location of the file(s) which should be replicated.
	Provider SourceFileSecureUpdateStorageProvider `json:"provider"`
	// This should be a string in JSON format. It depends on the chosen file format to provide additional options and tune its behavior.
	ReaderOptions *string `json:"reader_options,omitempty"`
	// The URL path to access the file which should be replicated.
	URL string `json:"url"`
}

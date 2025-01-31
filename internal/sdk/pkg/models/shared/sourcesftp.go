// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethod - Connect through ssh key
type SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethod string

const (
	SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethodSSHKeyAuth SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethod = "SSH_KEY_AUTH"
)

func (e SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethod) ToPointer() *SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethod {
	return &e
}

func (e *SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethod: %v", v)
	}
}

// SourceSftpAuthenticationWildcardSSHKeyAuthentication - The server authentication method
type SourceSftpAuthenticationWildcardSSHKeyAuthentication struct {
	// Connect through ssh key
	AuthMethod SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethod `json:"auth_method"`
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	AuthSSHKey string `json:"auth_ssh_key"`
}

// SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethod - Connect through password authentication
type SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethod string

const (
	SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethodSSHPasswordAuth SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethod = "SSH_PASSWORD_AUTH"
)

func (e SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethod) ToPointer() *SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethod {
	return &e
}

func (e *SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethod: %v", v)
	}
}

// SourceSftpAuthenticationWildcardPasswordAuthentication - The server authentication method
type SourceSftpAuthenticationWildcardPasswordAuthentication struct {
	// Connect through password authentication
	AuthMethod SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethod `json:"auth_method"`
	// OS-level password for logging into the jump server host
	AuthUserPassword string `json:"auth_user_password"`
}

type SourceSftpAuthenticationWildcardType string

const (
	SourceSftpAuthenticationWildcardTypeSourceSftpAuthenticationWildcardPasswordAuthentication SourceSftpAuthenticationWildcardType = "source-sftp_Authentication *_Password Authentication"
	SourceSftpAuthenticationWildcardTypeSourceSftpAuthenticationWildcardSSHKeyAuthentication   SourceSftpAuthenticationWildcardType = "source-sftp_Authentication *_SSH Key Authentication"
)

type SourceSftpAuthenticationWildcard struct {
	SourceSftpAuthenticationWildcardPasswordAuthentication *SourceSftpAuthenticationWildcardPasswordAuthentication
	SourceSftpAuthenticationWildcardSSHKeyAuthentication   *SourceSftpAuthenticationWildcardSSHKeyAuthentication

	Type SourceSftpAuthenticationWildcardType
}

func CreateSourceSftpAuthenticationWildcardSourceSftpAuthenticationWildcardPasswordAuthentication(sourceSftpAuthenticationWildcardPasswordAuthentication SourceSftpAuthenticationWildcardPasswordAuthentication) SourceSftpAuthenticationWildcard {
	typ := SourceSftpAuthenticationWildcardTypeSourceSftpAuthenticationWildcardPasswordAuthentication

	return SourceSftpAuthenticationWildcard{
		SourceSftpAuthenticationWildcardPasswordAuthentication: &sourceSftpAuthenticationWildcardPasswordAuthentication,
		Type: typ,
	}
}

func CreateSourceSftpAuthenticationWildcardSourceSftpAuthenticationWildcardSSHKeyAuthentication(sourceSftpAuthenticationWildcardSSHKeyAuthentication SourceSftpAuthenticationWildcardSSHKeyAuthentication) SourceSftpAuthenticationWildcard {
	typ := SourceSftpAuthenticationWildcardTypeSourceSftpAuthenticationWildcardSSHKeyAuthentication

	return SourceSftpAuthenticationWildcard{
		SourceSftpAuthenticationWildcardSSHKeyAuthentication: &sourceSftpAuthenticationWildcardSSHKeyAuthentication,
		Type: typ,
	}
}

func (u *SourceSftpAuthenticationWildcard) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceSftpAuthenticationWildcardPasswordAuthentication := new(SourceSftpAuthenticationWildcardPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSftpAuthenticationWildcardPasswordAuthentication); err == nil {
		u.SourceSftpAuthenticationWildcardPasswordAuthentication = sourceSftpAuthenticationWildcardPasswordAuthentication
		u.Type = SourceSftpAuthenticationWildcardTypeSourceSftpAuthenticationWildcardPasswordAuthentication
		return nil
	}

	sourceSftpAuthenticationWildcardSSHKeyAuthentication := new(SourceSftpAuthenticationWildcardSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSftpAuthenticationWildcardSSHKeyAuthentication); err == nil {
		u.SourceSftpAuthenticationWildcardSSHKeyAuthentication = sourceSftpAuthenticationWildcardSSHKeyAuthentication
		u.Type = SourceSftpAuthenticationWildcardTypeSourceSftpAuthenticationWildcardSSHKeyAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSftpAuthenticationWildcard) MarshalJSON() ([]byte, error) {
	if u.SourceSftpAuthenticationWildcardPasswordAuthentication != nil {
		return json.Marshal(u.SourceSftpAuthenticationWildcardPasswordAuthentication)
	}

	if u.SourceSftpAuthenticationWildcardSSHKeyAuthentication != nil {
		return json.Marshal(u.SourceSftpAuthenticationWildcardSSHKeyAuthentication)
	}

	return nil, nil
}

type SourceSftpSftp string

const (
	SourceSftpSftpSftp SourceSftpSftp = "sftp"
)

func (e SourceSftpSftp) ToPointer() *SourceSftpSftp {
	return &e
}

func (e *SourceSftpSftp) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sftp":
		*e = SourceSftpSftp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSftpSftp: %v", v)
	}
}

type SourceSftp struct {
	// The server authentication method
	Credentials *SourceSftpAuthenticationWildcard `json:"credentials,omitempty"`
	// The regular expression to specify files for sync in a chosen Folder Path
	FilePattern *string `json:"file_pattern,omitempty"`
	// Coma separated file types. Currently only 'csv' and 'json' types are supported.
	FileTypes *string `json:"file_types,omitempty"`
	// The directory to search files for sync
	FolderPath *string `json:"folder_path,omitempty"`
	// The server host address
	Host string `json:"host"`
	// The server port
	Port       int64          `json:"port"`
	SourceType SourceSftpSftp `json:"sourceType"`
	// The server user
	User string `json:"user"`
}

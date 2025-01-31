// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceSftpUpdateAuthenticationWildcardSSHKeyAuthenticationAuthMethod - Connect through ssh key
type SourceSftpUpdateAuthenticationWildcardSSHKeyAuthenticationAuthMethod string

const (
	SourceSftpUpdateAuthenticationWildcardSSHKeyAuthenticationAuthMethodSSHKeyAuth SourceSftpUpdateAuthenticationWildcardSSHKeyAuthenticationAuthMethod = "SSH_KEY_AUTH"
)

func (e SourceSftpUpdateAuthenticationWildcardSSHKeyAuthenticationAuthMethod) ToPointer() *SourceSftpUpdateAuthenticationWildcardSSHKeyAuthenticationAuthMethod {
	return &e
}

func (e *SourceSftpUpdateAuthenticationWildcardSSHKeyAuthenticationAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = SourceSftpUpdateAuthenticationWildcardSSHKeyAuthenticationAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSftpUpdateAuthenticationWildcardSSHKeyAuthenticationAuthMethod: %v", v)
	}
}

// SourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication - The server authentication method
type SourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication struct {
	// Connect through ssh key
	AuthMethod SourceSftpUpdateAuthenticationWildcardSSHKeyAuthenticationAuthMethod `json:"auth_method"`
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	AuthSSHKey string `json:"auth_ssh_key"`
}

// SourceSftpUpdateAuthenticationWildcardPasswordAuthenticationAuthMethod - Connect through password authentication
type SourceSftpUpdateAuthenticationWildcardPasswordAuthenticationAuthMethod string

const (
	SourceSftpUpdateAuthenticationWildcardPasswordAuthenticationAuthMethodSSHPasswordAuth SourceSftpUpdateAuthenticationWildcardPasswordAuthenticationAuthMethod = "SSH_PASSWORD_AUTH"
)

func (e SourceSftpUpdateAuthenticationWildcardPasswordAuthenticationAuthMethod) ToPointer() *SourceSftpUpdateAuthenticationWildcardPasswordAuthenticationAuthMethod {
	return &e
}

func (e *SourceSftpUpdateAuthenticationWildcardPasswordAuthenticationAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = SourceSftpUpdateAuthenticationWildcardPasswordAuthenticationAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSftpUpdateAuthenticationWildcardPasswordAuthenticationAuthMethod: %v", v)
	}
}

// SourceSftpUpdateAuthenticationWildcardPasswordAuthentication - The server authentication method
type SourceSftpUpdateAuthenticationWildcardPasswordAuthentication struct {
	// Connect through password authentication
	AuthMethod SourceSftpUpdateAuthenticationWildcardPasswordAuthenticationAuthMethod `json:"auth_method"`
	// OS-level password for logging into the jump server host
	AuthUserPassword string `json:"auth_user_password"`
}

type SourceSftpUpdateAuthenticationWildcardType string

const (
	SourceSftpUpdateAuthenticationWildcardTypeSourceSftpUpdateAuthenticationWildcardPasswordAuthentication SourceSftpUpdateAuthenticationWildcardType = "source-sftp-update_Authentication *_Password Authentication"
	SourceSftpUpdateAuthenticationWildcardTypeSourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication   SourceSftpUpdateAuthenticationWildcardType = "source-sftp-update_Authentication *_SSH Key Authentication"
)

type SourceSftpUpdateAuthenticationWildcard struct {
	SourceSftpUpdateAuthenticationWildcardPasswordAuthentication *SourceSftpUpdateAuthenticationWildcardPasswordAuthentication
	SourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication   *SourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication

	Type SourceSftpUpdateAuthenticationWildcardType
}

func CreateSourceSftpUpdateAuthenticationWildcardSourceSftpUpdateAuthenticationWildcardPasswordAuthentication(sourceSftpUpdateAuthenticationWildcardPasswordAuthentication SourceSftpUpdateAuthenticationWildcardPasswordAuthentication) SourceSftpUpdateAuthenticationWildcard {
	typ := SourceSftpUpdateAuthenticationWildcardTypeSourceSftpUpdateAuthenticationWildcardPasswordAuthentication

	return SourceSftpUpdateAuthenticationWildcard{
		SourceSftpUpdateAuthenticationWildcardPasswordAuthentication: &sourceSftpUpdateAuthenticationWildcardPasswordAuthentication,
		Type: typ,
	}
}

func CreateSourceSftpUpdateAuthenticationWildcardSourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication(sourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication SourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication) SourceSftpUpdateAuthenticationWildcard {
	typ := SourceSftpUpdateAuthenticationWildcardTypeSourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication

	return SourceSftpUpdateAuthenticationWildcard{
		SourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication: &sourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication,
		Type: typ,
	}
}

func (u *SourceSftpUpdateAuthenticationWildcard) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceSftpUpdateAuthenticationWildcardPasswordAuthentication := new(SourceSftpUpdateAuthenticationWildcardPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSftpUpdateAuthenticationWildcardPasswordAuthentication); err == nil {
		u.SourceSftpUpdateAuthenticationWildcardPasswordAuthentication = sourceSftpUpdateAuthenticationWildcardPasswordAuthentication
		u.Type = SourceSftpUpdateAuthenticationWildcardTypeSourceSftpUpdateAuthenticationWildcardPasswordAuthentication
		return nil
	}

	sourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication := new(SourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication); err == nil {
		u.SourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication = sourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication
		u.Type = SourceSftpUpdateAuthenticationWildcardTypeSourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSftpUpdateAuthenticationWildcard) MarshalJSON() ([]byte, error) {
	if u.SourceSftpUpdateAuthenticationWildcardPasswordAuthentication != nil {
		return json.Marshal(u.SourceSftpUpdateAuthenticationWildcardPasswordAuthentication)
	}

	if u.SourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication != nil {
		return json.Marshal(u.SourceSftpUpdateAuthenticationWildcardSSHKeyAuthentication)
	}

	return nil, nil
}

type SourceSftpUpdate struct {
	// The server authentication method
	Credentials *SourceSftpUpdateAuthenticationWildcard `json:"credentials,omitempty"`
	// The regular expression to specify files for sync in a chosen Folder Path
	FilePattern *string `json:"file_pattern,omitempty"`
	// Coma separated file types. Currently only 'csv' and 'json' types are supported.
	FileTypes *string `json:"file_types,omitempty"`
	// The directory to search files for sync
	FolderPath *string `json:"folder_path,omitempty"`
	// The server host address
	Host string `json:"host"`
	// The server port
	Port int64 `json:"port"`
	// The server user
	User string `json:"user"`
}

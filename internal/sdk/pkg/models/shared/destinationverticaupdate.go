// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// DestinationVerticaUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationVerticaUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationVerticaUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationVerticaUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationVerticaUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationVerticaUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationVerticaUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationVerticaUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationVerticaUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationVerticaUpdateSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationVerticaUpdateSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationVerticaUpdateSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationVerticaUpdateSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationVerticaUpdateSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationVerticaUpdateSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationVerticaUpdateSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationVerticaUpdateSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaUpdateSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationVerticaUpdateSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaUpdateSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationVerticaUpdateSSHTunnelMethodNoTunnelTunnelMethod `json:"tunnel_method"`
}

type DestinationVerticaUpdateSSHTunnelMethodType string

const (
	DestinationVerticaUpdateSSHTunnelMethodTypeDestinationVerticaUpdateSSHTunnelMethodNoTunnel               DestinationVerticaUpdateSSHTunnelMethodType = "destination-vertica-update_SSH Tunnel Method_No Tunnel"
	DestinationVerticaUpdateSSHTunnelMethodTypeDestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication   DestinationVerticaUpdateSSHTunnelMethodType = "destination-vertica-update_SSH Tunnel Method_SSH Key Authentication"
	DestinationVerticaUpdateSSHTunnelMethodTypeDestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication DestinationVerticaUpdateSSHTunnelMethodType = "destination-vertica-update_SSH Tunnel Method_Password Authentication"
)

type DestinationVerticaUpdateSSHTunnelMethod struct {
	DestinationVerticaUpdateSSHTunnelMethodNoTunnel               *DestinationVerticaUpdateSSHTunnelMethodNoTunnel
	DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication   *DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication
	DestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication *DestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication

	Type DestinationVerticaUpdateSSHTunnelMethodType
}

func CreateDestinationVerticaUpdateSSHTunnelMethodDestinationVerticaUpdateSSHTunnelMethodNoTunnel(destinationVerticaUpdateSSHTunnelMethodNoTunnel DestinationVerticaUpdateSSHTunnelMethodNoTunnel) DestinationVerticaUpdateSSHTunnelMethod {
	typ := DestinationVerticaUpdateSSHTunnelMethodTypeDestinationVerticaUpdateSSHTunnelMethodNoTunnel

	return DestinationVerticaUpdateSSHTunnelMethod{
		DestinationVerticaUpdateSSHTunnelMethodNoTunnel: &destinationVerticaUpdateSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationVerticaUpdateSSHTunnelMethodDestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication(destinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication) DestinationVerticaUpdateSSHTunnelMethod {
	typ := DestinationVerticaUpdateSSHTunnelMethodTypeDestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication

	return DestinationVerticaUpdateSSHTunnelMethod{
		DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication: &destinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationVerticaUpdateSSHTunnelMethodDestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication(destinationVerticaUpdateSSHTunnelMethodPasswordAuthentication DestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication) DestinationVerticaUpdateSSHTunnelMethod {
	typ := DestinationVerticaUpdateSSHTunnelMethodTypeDestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication

	return DestinationVerticaUpdateSSHTunnelMethod{
		DestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication: &destinationVerticaUpdateSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationVerticaUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationVerticaUpdateSSHTunnelMethodNoTunnel := new(DestinationVerticaUpdateSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationVerticaUpdateSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationVerticaUpdateSSHTunnelMethodNoTunnel = destinationVerticaUpdateSSHTunnelMethodNoTunnel
		u.Type = DestinationVerticaUpdateSSHTunnelMethodTypeDestinationVerticaUpdateSSHTunnelMethodNoTunnel
		return nil
	}

	destinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication := new(DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication = destinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationVerticaUpdateSSHTunnelMethodTypeDestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationVerticaUpdateSSHTunnelMethodPasswordAuthentication := new(DestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationVerticaUpdateSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication = destinationVerticaUpdateSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationVerticaUpdateSSHTunnelMethodTypeDestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationVerticaUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationVerticaUpdateSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationVerticaUpdateSSHTunnelMethodNoTunnel)
	}

	if u.DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type DestinationVerticaUpdate struct {
	// Name of the database.
	Database string `json:"database"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port int64 `json:"port"`
	// Schema for vertica destination
	Schema string `json:"schema"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationVerticaUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

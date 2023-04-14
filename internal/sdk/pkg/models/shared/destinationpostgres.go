// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationPostgresPostgresEnum string

const (
	DestinationPostgresPostgresEnumPostgres DestinationPostgresPostgresEnum = "postgres"
)

func (e *DestinationPostgresPostgresEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "postgres":
		*e = DestinationPostgresPostgresEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresPostgresEnum: %s", s)
	}
}

type DestinationPostgresSSLModesVerifyFullModeEnum string

const (
	DestinationPostgresSSLModesVerifyFullModeEnumVerifyFull DestinationPostgresSSLModesVerifyFullModeEnum = "verify-full"
)

func (e *DestinationPostgresSSLModesVerifyFullModeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "verify-full":
		*e = DestinationPostgresSSLModesVerifyFullModeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesVerifyFullModeEnum: %s", s)
	}
}

// DestinationPostgresSSLModesVerifyFull - Verify-full SSL mode.
type DestinationPostgresSSLModesVerifyFull struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate
	ClientCertificate string `json:"client_certificate"`
	// Client key
	ClientKey string `json:"client_key"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                       `json:"client_key_password,omitempty"`
	Mode              DestinationPostgresSSLModesVerifyFullModeEnum `json:"mode"`
}

type DestinationPostgresSSLModesVerifyCaModeEnum string

const (
	DestinationPostgresSSLModesVerifyCaModeEnumVerifyCa DestinationPostgresSSLModesVerifyCaModeEnum = "verify-ca"
)

func (e *DestinationPostgresSSLModesVerifyCaModeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "verify-ca":
		*e = DestinationPostgresSSLModesVerifyCaModeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesVerifyCaModeEnum: %s", s)
	}
}

// DestinationPostgresSSLModesVerifyCa - Verify-ca SSL mode.
type DestinationPostgresSSLModesVerifyCa struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                     `json:"client_key_password,omitempty"`
	Mode              DestinationPostgresSSLModesVerifyCaModeEnum `json:"mode"`
}

type DestinationPostgresSSLModesRequireModeEnum string

const (
	DestinationPostgresSSLModesRequireModeEnumRequire DestinationPostgresSSLModesRequireModeEnum = "require"
)

func (e *DestinationPostgresSSLModesRequireModeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "require":
		*e = DestinationPostgresSSLModesRequireModeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesRequireModeEnum: %s", s)
	}
}

// DestinationPostgresSSLModesRequire - Require SSL mode.
type DestinationPostgresSSLModesRequire struct {
	Mode DestinationPostgresSSLModesRequireModeEnum `json:"mode"`
}

type DestinationPostgresSSLModesPreferModeEnum string

const (
	DestinationPostgresSSLModesPreferModeEnumPrefer DestinationPostgresSSLModesPreferModeEnum = "prefer"
)

func (e *DestinationPostgresSSLModesPreferModeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "prefer":
		*e = DestinationPostgresSSLModesPreferModeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesPreferModeEnum: %s", s)
	}
}

// DestinationPostgresSSLModesPrefer - Prefer SSL mode.
type DestinationPostgresSSLModesPrefer struct {
	Mode DestinationPostgresSSLModesPreferModeEnum `json:"mode"`
}

type DestinationPostgresSSLModesAllowModeEnum string

const (
	DestinationPostgresSSLModesAllowModeEnumAllow DestinationPostgresSSLModesAllowModeEnum = "allow"
)

func (e *DestinationPostgresSSLModesAllowModeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "allow":
		*e = DestinationPostgresSSLModesAllowModeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesAllowModeEnum: %s", s)
	}
}

// DestinationPostgresSSLModesAllow - Allow SSL mode.
type DestinationPostgresSSLModesAllow struct {
	Mode DestinationPostgresSSLModesAllowModeEnum `json:"mode"`
}

type DestinationPostgresSSLModesDisableModeEnum string

const (
	DestinationPostgresSSLModesDisableModeEnumDisable DestinationPostgresSSLModesDisableModeEnum = "disable"
)

func (e *DestinationPostgresSSLModesDisableModeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "disable":
		*e = DestinationPostgresSSLModesDisableModeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesDisableModeEnum: %s", s)
	}
}

// DestinationPostgresSSLModesDisable - Disable SSL.
type DestinationPostgresSSLModesDisable struct {
	Mode DestinationPostgresSSLModesDisableModeEnum `json:"mode"`
}

type DestinationPostgresSSLModesType string

const (
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesDisable    DestinationPostgresSSLModesType = "destination-postgres_SSL modes_disable"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesAllow      DestinationPostgresSSLModesType = "destination-postgres_SSL modes_allow"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesPrefer     DestinationPostgresSSLModesType = "destination-postgres_SSL modes_prefer"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesRequire    DestinationPostgresSSLModesType = "destination-postgres_SSL modes_require"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyCa   DestinationPostgresSSLModesType = "destination-postgres_SSL modes_verify-ca"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyFull DestinationPostgresSSLModesType = "destination-postgres_SSL modes_verify-full"
)

type DestinationPostgresSSLModes struct {
	DestinationPostgresSSLModesDisable    *DestinationPostgresSSLModesDisable
	DestinationPostgresSSLModesAllow      *DestinationPostgresSSLModesAllow
	DestinationPostgresSSLModesPrefer     *DestinationPostgresSSLModesPrefer
	DestinationPostgresSSLModesRequire    *DestinationPostgresSSLModesRequire
	DestinationPostgresSSLModesVerifyCa   *DestinationPostgresSSLModesVerifyCa
	DestinationPostgresSSLModesVerifyFull *DestinationPostgresSSLModesVerifyFull

	Type DestinationPostgresSSLModesType
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesDisable(destinationPostgresSSLModesDisable DestinationPostgresSSLModesDisable) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesDisable

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesDisable: &destinationPostgresSSLModesDisable,
		Type:                               typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesAllow(destinationPostgresSSLModesAllow DestinationPostgresSSLModesAllow) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesAllow

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesAllow: &destinationPostgresSSLModesAllow,
		Type:                             typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesPrefer(destinationPostgresSSLModesPrefer DestinationPostgresSSLModesPrefer) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesPrefer

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesPrefer: &destinationPostgresSSLModesPrefer,
		Type:                              typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesRequire(destinationPostgresSSLModesRequire DestinationPostgresSSLModesRequire) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesRequire

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesRequire: &destinationPostgresSSLModesRequire,
		Type:                               typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesVerifyCa(destinationPostgresSSLModesVerifyCa DestinationPostgresSSLModesVerifyCa) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyCa

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesVerifyCa: &destinationPostgresSSLModesVerifyCa,
		Type:                                typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesVerifyFull(destinationPostgresSSLModesVerifyFull DestinationPostgresSSLModesVerifyFull) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyFull

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesVerifyFull: &destinationPostgresSSLModesVerifyFull,
		Type:                                  typ,
	}
}

func (u *DestinationPostgresSSLModes) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationPostgresSSLModesDisable := new(DestinationPostgresSSLModesDisable)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesDisable); err == nil {
		u.DestinationPostgresSSLModesDisable = destinationPostgresSSLModesDisable
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesDisable
		return nil
	}

	destinationPostgresSSLModesAllow := new(DestinationPostgresSSLModesAllow)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesAllow); err == nil {
		u.DestinationPostgresSSLModesAllow = destinationPostgresSSLModesAllow
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesAllow
		return nil
	}

	destinationPostgresSSLModesPrefer := new(DestinationPostgresSSLModesPrefer)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesPrefer); err == nil {
		u.DestinationPostgresSSLModesPrefer = destinationPostgresSSLModesPrefer
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesPrefer
		return nil
	}

	destinationPostgresSSLModesRequire := new(DestinationPostgresSSLModesRequire)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesRequire); err == nil {
		u.DestinationPostgresSSLModesRequire = destinationPostgresSSLModesRequire
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesRequire
		return nil
	}

	destinationPostgresSSLModesVerifyCa := new(DestinationPostgresSSLModesVerifyCa)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesVerifyCa); err == nil {
		u.DestinationPostgresSSLModesVerifyCa = destinationPostgresSSLModesVerifyCa
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyCa
		return nil
	}

	destinationPostgresSSLModesVerifyFull := new(DestinationPostgresSSLModesVerifyFull)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesVerifyFull); err == nil {
		u.DestinationPostgresSSLModesVerifyFull = destinationPostgresSSLModesVerifyFull
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyFull
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPostgresSSLModes) MarshalJSON() ([]byte, error) {
	if u.DestinationPostgresSSLModesDisable != nil {
		return json.Marshal(u.DestinationPostgresSSLModesDisable)
	}

	if u.DestinationPostgresSSLModesAllow != nil {
		return json.Marshal(u.DestinationPostgresSSLModesAllow)
	}

	if u.DestinationPostgresSSLModesPrefer != nil {
		return json.Marshal(u.DestinationPostgresSSLModesPrefer)
	}

	if u.DestinationPostgresSSLModesRequire != nil {
		return json.Marshal(u.DestinationPostgresSSLModesRequire)
	}

	if u.DestinationPostgresSSLModesVerifyCa != nil {
		return json.Marshal(u.DestinationPostgresSSLModesVerifyCa)
	}

	if u.DestinationPostgresSSLModesVerifyFull != nil {
		return json.Marshal(u.DestinationPostgresSSLModesVerifyFull)
	}

	return nil, nil
}

// DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum - Connect through a jump server tunnel host using username and password authentication
type DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum string

const (
	DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodEnumSSHPasswordAuth DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum = "SSH_PASSWORD_AUTH"
)

func (e *DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum: %s", s)
	}
}

// DestinationPostgresSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum - Connect through a jump server tunnel host using username and ssh key
type DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum string

const (
	DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnumSSHKeyAuth DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum = "SSH_KEY_AUTH"
)

func (e *DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "SSH_KEY_AUTH":
		*e = DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum: %s", s)
	}
}

// DestinationPostgresSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodEnum - No ssh tunnel needed to connect to database
type DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodEnum string

const (
	DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodEnumNoTunnel DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodEnum = "NO_TUNNEL"
)

func (e *DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "NO_TUNNEL":
		*e = DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodEnum: %s", s)
	}
}

// DestinationPostgresSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodEnum `json:"tunnel_method"`
}

type DestinationPostgresSSHTunnelMethodType string

const (
	DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodNoTunnel               DestinationPostgresSSHTunnelMethodType = "destination-postgres_SSH Tunnel Method_No Tunnel"
	DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodSSHKeyAuthentication   DestinationPostgresSSHTunnelMethodType = "destination-postgres_SSH Tunnel Method_SSH Key Authentication"
	DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodPasswordAuthentication DestinationPostgresSSHTunnelMethodType = "destination-postgres_SSH Tunnel Method_Password Authentication"
)

type DestinationPostgresSSHTunnelMethod struct {
	DestinationPostgresSSHTunnelMethodNoTunnel               *DestinationPostgresSSHTunnelMethodNoTunnel
	DestinationPostgresSSHTunnelMethodSSHKeyAuthentication   *DestinationPostgresSSHTunnelMethodSSHKeyAuthentication
	DestinationPostgresSSHTunnelMethodPasswordAuthentication *DestinationPostgresSSHTunnelMethodPasswordAuthentication

	Type DestinationPostgresSSHTunnelMethodType
}

func CreateDestinationPostgresSSHTunnelMethodDestinationPostgresSSHTunnelMethodNoTunnel(destinationPostgresSSHTunnelMethodNoTunnel DestinationPostgresSSHTunnelMethodNoTunnel) DestinationPostgresSSHTunnelMethod {
	typ := DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodNoTunnel

	return DestinationPostgresSSHTunnelMethod{
		DestinationPostgresSSHTunnelMethodNoTunnel: &destinationPostgresSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationPostgresSSHTunnelMethodDestinationPostgresSSHTunnelMethodSSHKeyAuthentication(destinationPostgresSSHTunnelMethodSSHKeyAuthentication DestinationPostgresSSHTunnelMethodSSHKeyAuthentication) DestinationPostgresSSHTunnelMethod {
	typ := DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodSSHKeyAuthentication

	return DestinationPostgresSSHTunnelMethod{
		DestinationPostgresSSHTunnelMethodSSHKeyAuthentication: &destinationPostgresSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationPostgresSSHTunnelMethodDestinationPostgresSSHTunnelMethodPasswordAuthentication(destinationPostgresSSHTunnelMethodPasswordAuthentication DestinationPostgresSSHTunnelMethodPasswordAuthentication) DestinationPostgresSSHTunnelMethod {
	typ := DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodPasswordAuthentication

	return DestinationPostgresSSHTunnelMethod{
		DestinationPostgresSSHTunnelMethodPasswordAuthentication: &destinationPostgresSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationPostgresSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationPostgresSSHTunnelMethodNoTunnel := new(DestinationPostgresSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationPostgresSSHTunnelMethodNoTunnel = destinationPostgresSSHTunnelMethodNoTunnel
		u.Type = DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodNoTunnel
		return nil
	}

	destinationPostgresSSHTunnelMethodSSHKeyAuthentication := new(DestinationPostgresSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication = destinationPostgresSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationPostgresSSHTunnelMethodPasswordAuthentication := new(DestinationPostgresSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationPostgresSSHTunnelMethodPasswordAuthentication = destinationPostgresSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPostgresSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationPostgresSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationPostgresSSHTunnelMethodNoTunnel)
	}

	if u.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationPostgresSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationPostgresSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

// DestinationPostgres - The values required to configure the destination.
type DestinationPostgres struct {
	// Name of the database.
	Database        string                          `json:"database"`
	DestinationType DestinationPostgresPostgresEnum `json:"destinationType"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port int64 `json:"port"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Schema string `json:"schema"`
	// SSL connection modes.
	//  <b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database
	//  <b>allow</b> - Chose this mode to enable encryption only when required by the source database
	//  <b>prefer</b> - Chose this mode to allow unencrypted connection only if the source database does not support encryption
	//  <b>require</b> - Chose this mode to always require encryption. If the source database server does not support encryption, connection will fail
	//   <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the source database server has a valid SSL certificate
	//   <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the source database server
	//  See more information - <a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"> in the docs</a>.
	SslMode *DestinationPostgresSSLModes `json:"ssl_mode,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationPostgresSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

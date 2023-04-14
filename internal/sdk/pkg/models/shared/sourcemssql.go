// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceMssqlReplicationMethodLogicalReplicationCDCDataToSyncEnum - What data should be synced under the CDC. "Existing and New" will read existing data as a snapshot, and sync new changes through CDC. "New Changes Only" will skip the initial snapshot, and only sync new changes through CDC.
type SourceMssqlReplicationMethodLogicalReplicationCDCDataToSyncEnum string

const (
	SourceMssqlReplicationMethodLogicalReplicationCDCDataToSyncEnumExistingAndNew SourceMssqlReplicationMethodLogicalReplicationCDCDataToSyncEnum = "Existing and New"
	SourceMssqlReplicationMethodLogicalReplicationCDCDataToSyncEnumNewChangesOnly SourceMssqlReplicationMethodLogicalReplicationCDCDataToSyncEnum = "New Changes Only"
)

func (e *SourceMssqlReplicationMethodLogicalReplicationCDCDataToSyncEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Existing and New":
		fallthrough
	case "New Changes Only":
		*e = SourceMssqlReplicationMethodLogicalReplicationCDCDataToSyncEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlReplicationMethodLogicalReplicationCDCDataToSyncEnum: %s", s)
	}
}

type SourceMssqlReplicationMethodLogicalReplicationCDCMethodEnum string

const (
	SourceMssqlReplicationMethodLogicalReplicationCDCMethodEnumCdc SourceMssqlReplicationMethodLogicalReplicationCDCMethodEnum = "CDC"
)

func (e *SourceMssqlReplicationMethodLogicalReplicationCDCMethodEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "CDC":
		*e = SourceMssqlReplicationMethodLogicalReplicationCDCMethodEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlReplicationMethodLogicalReplicationCDCMethodEnum: %s", s)
	}
}

// SourceMssqlReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelEnum - Existing data in the database are synced through an initial snapshot. This parameter controls the isolation level that will be used during the initial snapshotting. If you choose the "Snapshot" level, you must enable the <a href="https://docs.microsoft.com/en-us/dotnet/framework/data/adonet/sql/snapshot-isolation-in-sql-server">snapshot isolation mode</a> on the database.
type SourceMssqlReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelEnum string

const (
	SourceMssqlReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelEnumSnapshot      SourceMssqlReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelEnum = "Snapshot"
	SourceMssqlReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelEnumReadCommitted SourceMssqlReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelEnum = "Read Committed"
)

func (e *SourceMssqlReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Snapshot":
		fallthrough
	case "Read Committed":
		*e = SourceMssqlReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelEnum: %s", s)
	}
}

// SourceMssqlReplicationMethodLogicalReplicationCDC - CDC uses {TBC} to detect inserts, updates, and deletes. This needs to be configured on the source database itself.
type SourceMssqlReplicationMethodLogicalReplicationCDC struct {
	// What data should be synced under the CDC. "Existing and New" will read existing data as a snapshot, and sync new changes through CDC. "New Changes Only" will skip the initial snapshot, and only sync new changes through CDC.
	DataToSync *SourceMssqlReplicationMethodLogicalReplicationCDCDataToSyncEnum `json:"data_to_sync,omitempty"`
	// The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">initial waiting time</a>.
	InitialWaitingSeconds *int64                                                      `json:"initial_waiting_seconds,omitempty"`
	Method                SourceMssqlReplicationMethodLogicalReplicationCDCMethodEnum `json:"method"`
	// Existing data in the database are synced through an initial snapshot. This parameter controls the isolation level that will be used during the initial snapshotting. If you choose the "Snapshot" level, you must enable the <a href="https://docs.microsoft.com/en-us/dotnet/framework/data/adonet/sql/snapshot-isolation-in-sql-server">snapshot isolation mode</a> on the database.
	SnapshotIsolation *SourceMssqlReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelEnum `json:"snapshot_isolation,omitempty"`
}

type SourceMssqlReplicationMethodStandardMethodEnum string

const (
	SourceMssqlReplicationMethodStandardMethodEnumStandard SourceMssqlReplicationMethodStandardMethodEnum = "STANDARD"
)

func (e *SourceMssqlReplicationMethodStandardMethodEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "STANDARD":
		*e = SourceMssqlReplicationMethodStandardMethodEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlReplicationMethodStandardMethodEnum: %s", s)
	}
}

// SourceMssqlReplicationMethodStandard - Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.
type SourceMssqlReplicationMethodStandard struct {
	Method SourceMssqlReplicationMethodStandardMethodEnum `json:"method"`
}

type SourceMssqlReplicationMethodType string

const (
	SourceMssqlReplicationMethodTypeSourceMssqlReplicationMethodStandard              SourceMssqlReplicationMethodType = "source-mssql_Replication Method_Standard"
	SourceMssqlReplicationMethodTypeSourceMssqlReplicationMethodLogicalReplicationCDC SourceMssqlReplicationMethodType = "source-mssql_Replication Method_Logical Replication (CDC)"
)

type SourceMssqlReplicationMethod struct {
	SourceMssqlReplicationMethodStandard              *SourceMssqlReplicationMethodStandard
	SourceMssqlReplicationMethodLogicalReplicationCDC *SourceMssqlReplicationMethodLogicalReplicationCDC

	Type SourceMssqlReplicationMethodType
}

func CreateSourceMssqlReplicationMethodSourceMssqlReplicationMethodStandard(sourceMssqlReplicationMethodStandard SourceMssqlReplicationMethodStandard) SourceMssqlReplicationMethod {
	typ := SourceMssqlReplicationMethodTypeSourceMssqlReplicationMethodStandard

	return SourceMssqlReplicationMethod{
		SourceMssqlReplicationMethodStandard: &sourceMssqlReplicationMethodStandard,
		Type:                                 typ,
	}
}

func CreateSourceMssqlReplicationMethodSourceMssqlReplicationMethodLogicalReplicationCDC(sourceMssqlReplicationMethodLogicalReplicationCDC SourceMssqlReplicationMethodLogicalReplicationCDC) SourceMssqlReplicationMethod {
	typ := SourceMssqlReplicationMethodTypeSourceMssqlReplicationMethodLogicalReplicationCDC

	return SourceMssqlReplicationMethod{
		SourceMssqlReplicationMethodLogicalReplicationCDC: &sourceMssqlReplicationMethodLogicalReplicationCDC,
		Type: typ,
	}
}

func (u *SourceMssqlReplicationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMssqlReplicationMethodStandard := new(SourceMssqlReplicationMethodStandard)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlReplicationMethodStandard); err == nil {
		u.SourceMssqlReplicationMethodStandard = sourceMssqlReplicationMethodStandard
		u.Type = SourceMssqlReplicationMethodTypeSourceMssqlReplicationMethodStandard
		return nil
	}

	sourceMssqlReplicationMethodLogicalReplicationCDC := new(SourceMssqlReplicationMethodLogicalReplicationCDC)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlReplicationMethodLogicalReplicationCDC); err == nil {
		u.SourceMssqlReplicationMethodLogicalReplicationCDC = sourceMssqlReplicationMethodLogicalReplicationCDC
		u.Type = SourceMssqlReplicationMethodTypeSourceMssqlReplicationMethodLogicalReplicationCDC
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMssqlReplicationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMssqlReplicationMethodStandard != nil {
		return json.Marshal(u.SourceMssqlReplicationMethodStandard)
	}

	if u.SourceMssqlReplicationMethodLogicalReplicationCDC != nil {
		return json.Marshal(u.SourceMssqlReplicationMethodLogicalReplicationCDC)
	}

	return nil, nil
}

type SourceMssqlMssqlEnum string

const (
	SourceMssqlMssqlEnumMssql SourceMssqlMssqlEnum = "mssql"
)

func (e *SourceMssqlMssqlEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "mssql":
		*e = SourceMssqlMssqlEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlMssqlEnum: %s", s)
	}
}

type SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum string

const (
	SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnumEncryptedVerifyCertificate SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum = "encrypted_verify_certificate"
)

func (e *SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "encrypted_verify_certificate":
		*e = SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum: %s", s)
	}
}

// SourceMssqlSSLMethodEncryptedVerifyCertificate - Verify and use the certificate provided by the server.
type SourceMssqlSSLMethodEncryptedVerifyCertificate struct {
	// Specifies the host name of the server. The value of this property must match the subject property of the certificate.
	HostNameInCertificate *string                                                     `json:"hostNameInCertificate,omitempty"`
	SslMethod             SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum `json:"ssl_method"`
}

type SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum string

const (
	SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnumEncryptedTrustServerCertificate SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum = "encrypted_trust_server_certificate"
)

func (e *SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "encrypted_trust_server_certificate":
		*e = SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum: %s", s)
	}
}

// SourceMssqlSSLMethodEncryptedTrustServerCertificate - Use the certificate provided by the server without verification. (For testing purposes only!)
type SourceMssqlSSLMethodEncryptedTrustServerCertificate struct {
	SslMethod SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum `json:"ssl_method"`
}

type SourceMssqlSSLMethodType string

const (
	SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedTrustServerCertificate SourceMssqlSSLMethodType = "source-mssql_SSL Method_Encrypted (trust server certificate)"
	SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedVerifyCertificate      SourceMssqlSSLMethodType = "source-mssql_SSL Method_Encrypted (verify certificate)"
)

type SourceMssqlSSLMethod struct {
	SourceMssqlSSLMethodEncryptedTrustServerCertificate *SourceMssqlSSLMethodEncryptedTrustServerCertificate
	SourceMssqlSSLMethodEncryptedVerifyCertificate      *SourceMssqlSSLMethodEncryptedVerifyCertificate

	Type SourceMssqlSSLMethodType
}

func CreateSourceMssqlSSLMethodSourceMssqlSSLMethodEncryptedTrustServerCertificate(sourceMssqlSSLMethodEncryptedTrustServerCertificate SourceMssqlSSLMethodEncryptedTrustServerCertificate) SourceMssqlSSLMethod {
	typ := SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedTrustServerCertificate

	return SourceMssqlSSLMethod{
		SourceMssqlSSLMethodEncryptedTrustServerCertificate: &sourceMssqlSSLMethodEncryptedTrustServerCertificate,
		Type: typ,
	}
}

func CreateSourceMssqlSSLMethodSourceMssqlSSLMethodEncryptedVerifyCertificate(sourceMssqlSSLMethodEncryptedVerifyCertificate SourceMssqlSSLMethodEncryptedVerifyCertificate) SourceMssqlSSLMethod {
	typ := SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedVerifyCertificate

	return SourceMssqlSSLMethod{
		SourceMssqlSSLMethodEncryptedVerifyCertificate: &sourceMssqlSSLMethodEncryptedVerifyCertificate,
		Type: typ,
	}
}

func (u *SourceMssqlSSLMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMssqlSSLMethodEncryptedTrustServerCertificate := new(SourceMssqlSSLMethodEncryptedTrustServerCertificate)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlSSLMethodEncryptedTrustServerCertificate); err == nil {
		u.SourceMssqlSSLMethodEncryptedTrustServerCertificate = sourceMssqlSSLMethodEncryptedTrustServerCertificate
		u.Type = SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedTrustServerCertificate
		return nil
	}

	sourceMssqlSSLMethodEncryptedVerifyCertificate := new(SourceMssqlSSLMethodEncryptedVerifyCertificate)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlSSLMethodEncryptedVerifyCertificate); err == nil {
		u.SourceMssqlSSLMethodEncryptedVerifyCertificate = sourceMssqlSSLMethodEncryptedVerifyCertificate
		u.Type = SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedVerifyCertificate
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMssqlSSLMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMssqlSSLMethodEncryptedTrustServerCertificate != nil {
		return json.Marshal(u.SourceMssqlSSLMethodEncryptedTrustServerCertificate)
	}

	if u.SourceMssqlSSLMethodEncryptedVerifyCertificate != nil {
		return json.Marshal(u.SourceMssqlSSLMethodEncryptedVerifyCertificate)
	}

	return nil, nil
}

// SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum - Connect through a jump server tunnel host using username and password authentication
type SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum string

const (
	SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnumSSHPasswordAuth SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum = "SSH_PASSWORD_AUTH"
)

func (e *SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "SSH_PASSWORD_AUTH":
		*e = SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum: %s", s)
	}
}

// SourceMssqlSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMssqlSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum - Connect through a jump server tunnel host using username and ssh key
type SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum string

const (
	SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnumSSHKeyAuth SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum = "SSH_KEY_AUTH"
)

func (e *SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "SSH_KEY_AUTH":
		*e = SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum: %s", s)
	}
}

// SourceMssqlSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMssqlSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// SourceMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum - No ssh tunnel needed to connect to database
type SourceMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum string

const (
	SourceMssqlSSHTunnelMethodNoTunnelTunnelMethodEnumNoTunnel SourceMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum = "NO_TUNNEL"
)

func (e *SourceMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "NO_TUNNEL":
		*e = SourceMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum: %s", s)
	}
}

// SourceMssqlSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMssqlSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod SourceMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum `json:"tunnel_method"`
}

type SourceMssqlSSHTunnelMethodType string

const (
	SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodNoTunnel               SourceMssqlSSHTunnelMethodType = "source-mssql_SSH Tunnel Method_No Tunnel"
	SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodSSHKeyAuthentication   SourceMssqlSSHTunnelMethodType = "source-mssql_SSH Tunnel Method_SSH Key Authentication"
	SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodPasswordAuthentication SourceMssqlSSHTunnelMethodType = "source-mssql_SSH Tunnel Method_Password Authentication"
)

type SourceMssqlSSHTunnelMethod struct {
	SourceMssqlSSHTunnelMethodNoTunnel               *SourceMssqlSSHTunnelMethodNoTunnel
	SourceMssqlSSHTunnelMethodSSHKeyAuthentication   *SourceMssqlSSHTunnelMethodSSHKeyAuthentication
	SourceMssqlSSHTunnelMethodPasswordAuthentication *SourceMssqlSSHTunnelMethodPasswordAuthentication

	Type SourceMssqlSSHTunnelMethodType
}

func CreateSourceMssqlSSHTunnelMethodSourceMssqlSSHTunnelMethodNoTunnel(sourceMssqlSSHTunnelMethodNoTunnel SourceMssqlSSHTunnelMethodNoTunnel) SourceMssqlSSHTunnelMethod {
	typ := SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodNoTunnel

	return SourceMssqlSSHTunnelMethod{
		SourceMssqlSSHTunnelMethodNoTunnel: &sourceMssqlSSHTunnelMethodNoTunnel,
		Type:                               typ,
	}
}

func CreateSourceMssqlSSHTunnelMethodSourceMssqlSSHTunnelMethodSSHKeyAuthentication(sourceMssqlSSHTunnelMethodSSHKeyAuthentication SourceMssqlSSHTunnelMethodSSHKeyAuthentication) SourceMssqlSSHTunnelMethod {
	typ := SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodSSHKeyAuthentication

	return SourceMssqlSSHTunnelMethod{
		SourceMssqlSSHTunnelMethodSSHKeyAuthentication: &sourceMssqlSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateSourceMssqlSSHTunnelMethodSourceMssqlSSHTunnelMethodPasswordAuthentication(sourceMssqlSSHTunnelMethodPasswordAuthentication SourceMssqlSSHTunnelMethodPasswordAuthentication) SourceMssqlSSHTunnelMethod {
	typ := SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodPasswordAuthentication

	return SourceMssqlSSHTunnelMethod{
		SourceMssqlSSHTunnelMethodPasswordAuthentication: &sourceMssqlSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *SourceMssqlSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMssqlSSHTunnelMethodNoTunnel := new(SourceMssqlSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlSSHTunnelMethodNoTunnel); err == nil {
		u.SourceMssqlSSHTunnelMethodNoTunnel = sourceMssqlSSHTunnelMethodNoTunnel
		u.Type = SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodNoTunnel
		return nil
	}

	sourceMssqlSSHTunnelMethodSSHKeyAuthentication := new(SourceMssqlSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.SourceMssqlSSHTunnelMethodSSHKeyAuthentication = sourceMssqlSSHTunnelMethodSSHKeyAuthentication
		u.Type = SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	sourceMssqlSSHTunnelMethodPasswordAuthentication := new(SourceMssqlSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlSSHTunnelMethodPasswordAuthentication); err == nil {
		u.SourceMssqlSSHTunnelMethodPasswordAuthentication = sourceMssqlSSHTunnelMethodPasswordAuthentication
		u.Type = SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMssqlSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMssqlSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.SourceMssqlSSHTunnelMethodNoTunnel)
	}

	if u.SourceMssqlSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.SourceMssqlSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.SourceMssqlSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.SourceMssqlSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

// SourceMssql - The values required to configure the source.
type SourceMssql struct {
	// The name of the database.
	Database string `json:"database"`
	// The hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with the username.
	Password *string `json:"password,omitempty"`
	// The port of the database.
	Port int64 `json:"port"`
	// The replication method used for extracting data from the database. STANDARD replication requires no setup on the DB side but will not be able to represent deletions incrementally. CDC uses {TBC} to detect inserts, updates, and deletes. This needs to be configured on the source database itself.
	ReplicationMethod *SourceMssqlReplicationMethod `json:"replication_method,omitempty"`
	// The list of schemas to sync from. Defaults to user. Case sensitive.
	Schemas    []string             `json:"schemas,omitempty"`
	SourceType SourceMssqlMssqlEnum `json:"sourceType"`
	// The encryption method which is used when communicating with the database.
	SslMethod *SourceMssqlSSLMethod `json:"ssl_method,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *SourceMssqlSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}

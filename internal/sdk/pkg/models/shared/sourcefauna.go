// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceFaunaCollectionDeletionModeEnabledDeletionMode string

const (
	SourceFaunaCollectionDeletionModeEnabledDeletionModeDeletedField SourceFaunaCollectionDeletionModeEnabledDeletionMode = "deleted_field"
)

func (e SourceFaunaCollectionDeletionModeEnabledDeletionMode) ToPointer() *SourceFaunaCollectionDeletionModeEnabledDeletionMode {
	return &e
}

func (e *SourceFaunaCollectionDeletionModeEnabledDeletionMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleted_field":
		*e = SourceFaunaCollectionDeletionModeEnabledDeletionMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFaunaCollectionDeletionModeEnabledDeletionMode: %v", v)
	}
}

// SourceFaunaCollectionDeletionModeEnabled - <b>This only applies to incremental syncs.</b> <br>
// Enabling deletion mode informs your destination of deleted documents.<br>
// Disabled - Leave this feature disabled, and ignore deleted documents.<br>
// Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.
type SourceFaunaCollectionDeletionModeEnabled struct {
	// Name of the "deleted at" column.
	Column       string                                               `json:"column"`
	DeletionMode SourceFaunaCollectionDeletionModeEnabledDeletionMode `json:"deletion_mode"`
}

type SourceFaunaCollectionDeletionModeDisabledDeletionMode string

const (
	SourceFaunaCollectionDeletionModeDisabledDeletionModeIgnore SourceFaunaCollectionDeletionModeDisabledDeletionMode = "ignore"
)

func (e SourceFaunaCollectionDeletionModeDisabledDeletionMode) ToPointer() *SourceFaunaCollectionDeletionModeDisabledDeletionMode {
	return &e
}

func (e *SourceFaunaCollectionDeletionModeDisabledDeletionMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ignore":
		*e = SourceFaunaCollectionDeletionModeDisabledDeletionMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFaunaCollectionDeletionModeDisabledDeletionMode: %v", v)
	}
}

// SourceFaunaCollectionDeletionModeDisabled - <b>This only applies to incremental syncs.</b> <br>
// Enabling deletion mode informs your destination of deleted documents.<br>
// Disabled - Leave this feature disabled, and ignore deleted documents.<br>
// Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.
type SourceFaunaCollectionDeletionModeDisabled struct {
	DeletionMode SourceFaunaCollectionDeletionModeDisabledDeletionMode `json:"deletion_mode"`
}

type SourceFaunaCollectionDeletionModeType string

const (
	SourceFaunaCollectionDeletionModeTypeSourceFaunaCollectionDeletionModeDisabled SourceFaunaCollectionDeletionModeType = "source-fauna_Collection_Deletion Mode_Disabled"
	SourceFaunaCollectionDeletionModeTypeSourceFaunaCollectionDeletionModeEnabled  SourceFaunaCollectionDeletionModeType = "source-fauna_Collection_Deletion Mode_Enabled"
)

type SourceFaunaCollectionDeletionMode struct {
	SourceFaunaCollectionDeletionModeDisabled *SourceFaunaCollectionDeletionModeDisabled
	SourceFaunaCollectionDeletionModeEnabled  *SourceFaunaCollectionDeletionModeEnabled

	Type SourceFaunaCollectionDeletionModeType
}

func CreateSourceFaunaCollectionDeletionModeSourceFaunaCollectionDeletionModeDisabled(sourceFaunaCollectionDeletionModeDisabled SourceFaunaCollectionDeletionModeDisabled) SourceFaunaCollectionDeletionMode {
	typ := SourceFaunaCollectionDeletionModeTypeSourceFaunaCollectionDeletionModeDisabled

	return SourceFaunaCollectionDeletionMode{
		SourceFaunaCollectionDeletionModeDisabled: &sourceFaunaCollectionDeletionModeDisabled,
		Type: typ,
	}
}

func CreateSourceFaunaCollectionDeletionModeSourceFaunaCollectionDeletionModeEnabled(sourceFaunaCollectionDeletionModeEnabled SourceFaunaCollectionDeletionModeEnabled) SourceFaunaCollectionDeletionMode {
	typ := SourceFaunaCollectionDeletionModeTypeSourceFaunaCollectionDeletionModeEnabled

	return SourceFaunaCollectionDeletionMode{
		SourceFaunaCollectionDeletionModeEnabled: &sourceFaunaCollectionDeletionModeEnabled,
		Type:                                     typ,
	}
}

func (u *SourceFaunaCollectionDeletionMode) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceFaunaCollectionDeletionModeDisabled := new(SourceFaunaCollectionDeletionModeDisabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFaunaCollectionDeletionModeDisabled); err == nil {
		u.SourceFaunaCollectionDeletionModeDisabled = sourceFaunaCollectionDeletionModeDisabled
		u.Type = SourceFaunaCollectionDeletionModeTypeSourceFaunaCollectionDeletionModeDisabled
		return nil
	}

	sourceFaunaCollectionDeletionModeEnabled := new(SourceFaunaCollectionDeletionModeEnabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFaunaCollectionDeletionModeEnabled); err == nil {
		u.SourceFaunaCollectionDeletionModeEnabled = sourceFaunaCollectionDeletionModeEnabled
		u.Type = SourceFaunaCollectionDeletionModeTypeSourceFaunaCollectionDeletionModeEnabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceFaunaCollectionDeletionMode) MarshalJSON() ([]byte, error) {
	if u.SourceFaunaCollectionDeletionModeDisabled != nil {
		return json.Marshal(u.SourceFaunaCollectionDeletionModeDisabled)
	}

	if u.SourceFaunaCollectionDeletionModeEnabled != nil {
		return json.Marshal(u.SourceFaunaCollectionDeletionModeEnabled)
	}

	return nil, nil
}

// SourceFaunaCollection - Settings for the Fauna Collection.
type SourceFaunaCollection struct {
	// <b>This only applies to incremental syncs.</b> <br>
	// Enabling deletion mode informs your destination of deleted documents.<br>
	// Disabled - Leave this feature disabled, and ignore deleted documents.<br>
	// Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.
	Deletions SourceFaunaCollectionDeletionMode `json:"deletions"`
	// The page size used when reading documents from the database. The larger the page size, the faster the connector processes documents. However, if a page is too large, the connector may fail. <br>
	// Choose your page size based on how large the documents are. <br>
	// See <a href="https://docs.fauna.com/fauna/current/learn/understanding/types#page">the docs</a>.
	PageSize int64 `json:"page_size"`
}

type SourceFaunaFauna string

const (
	SourceFaunaFaunaFauna SourceFaunaFauna = "fauna"
)

func (e SourceFaunaFauna) ToPointer() *SourceFaunaFauna {
	return &e
}

func (e *SourceFaunaFauna) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fauna":
		*e = SourceFaunaFauna(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFaunaFauna: %v", v)
	}
}

type SourceFauna struct {
	// Settings for the Fauna Collection.
	Collection *SourceFaunaCollection `json:"collection,omitempty"`
	// Domain of Fauna to query. Defaults db.fauna.com. See <a href=https://docs.fauna.com/fauna/current/learn/understanding/region_groups#how-to-use-region-groups>the docs</a>.
	Domain string `json:"domain"`
	// Endpoint port.
	Port int64 `json:"port"`
	// URL scheme.
	Scheme string `json:"scheme"`
	// Fauna secret, used when authenticating with the database.
	Secret     string           `json:"secret"`
	SourceType SourceFaunaFauna `json:"sourceType"`
}

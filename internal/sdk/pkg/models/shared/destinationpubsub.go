// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DestinationPubsubPubsub string

const (
	DestinationPubsubPubsubPubsub DestinationPubsubPubsub = "pubsub"
)

func (e DestinationPubsubPubsub) ToPointer() *DestinationPubsubPubsub {
	return &e
}

func (e *DestinationPubsubPubsub) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pubsub":
		*e = DestinationPubsubPubsub(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPubsubPubsub: %v", v)
	}
}

type DestinationPubsub struct {
	// Number of ms before the buffer is flushed
	BatchingDelayThreshold *int64 `json:"batching_delay_threshold,omitempty"`
	// Number of messages before the buffer is flushed
	BatchingElementCountThreshold *int64 `json:"batching_element_count_threshold,omitempty"`
	// If TRUE messages will be buffered instead of sending them one by one
	BatchingEnabled bool `json:"batching_enabled"`
	// Number of bytes before the buffer is flushed
	BatchingRequestBytesThreshold *int64 `json:"batching_request_bytes_threshold,omitempty"`
	// The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.com/integrations/destinations/pubsub">docs</a> if you need help generating this key.
	CredentialsJSON string                  `json:"credentials_json"`
	DestinationType DestinationPubsubPubsub `json:"destinationType"`
	// If TRUE PubSub publisher will have <a href="https://cloud.google.com/pubsub/docs/ordering">message ordering</a> enabled. Every message will have an ordering key of stream
	OrderingEnabled bool `json:"ordering_enabled"`
	// The GCP project ID for the project containing the target PubSub.
	ProjectID string `json:"project_id"`
	// The PubSub topic ID in the given GCP project ID.
	TopicID string `json:"topic_id"`
}

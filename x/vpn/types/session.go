package types

import (
	"encoding/hex"
	"fmt"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Session struct {
	ID                  sdkTypes.ID        `json:"id"`
	SubscriptionID      sdkTypes.ID        `json:"subscription_id"`
	Bandwidth           sdkTypes.Bandwidth `json:"bandwidth"`
	CalculatedBandwidth sdkTypes.Bandwidth `json:"calculated_bandwidth"`
	NodeOwnerSign       []byte             `json:"node_owner_sign"`
	ClientSign          []byte             `json:"client_sign"`
	Status              string             `json:"status"`
	StatusModifiedAt    int64              `json:"status_modified_at"`
}

func (s Session) String() string {
	nodeOwnerSign := hex.EncodeToString(s.NodeOwnerSign)
	clientSign := hex.EncodeToString(s.ClientSign)

	return fmt.Sprintf(`Session
  ID:                   %s
  Subscription ID:      %s
  Bandwidth:            %s
  Calculated Bandwidth: %s
  Node Owner Signature: %s
  Client Signature:     %s
  Status:               %s
  Status Modified At:   %d`, s.ID, s.SubscriptionID, s.Bandwidth, s.CalculatedBandwidth,
		nodeOwnerSign, clientSign, s.Status, s.StatusModifiedAt)
}

func (s Session) IsValid() error {
	if s.ID == nil || s.ID.Len() < 26 {
		return fmt.Errorf("invalid id")
	}
	if s.SubscriptionID == nil || s.SubscriptionID.Len() < 24 {
		return fmt.Errorf("invalid subscription id")
	}
	if s.Bandwidth.AnyNil() {
		return fmt.Errorf("invalid bandwidth")
	}
	if s.CalculatedBandwidth.AnyNil() {
		return fmt.Errorf("invalid calculated bandwidth")
	}
	if s.Status != StatusActive && s.Status != StatusInactive {
		return fmt.Errorf("invalid status")
	}

	return nil
}

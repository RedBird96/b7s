package response

import (
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/RedBird96/b7s/consensus"
	"github.com/RedBird96/b7s/models/codes"
)

// FormCluster describes the `MessageFormClusteRr` response.
type FormCluster struct {
	Type      string         `json:"type,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
	From      peer.ID        `json:"from,omitempty"`
	Code      codes.Code     `json:"code,omitempty"`
	Consensus consensus.Type `json:"consensus,omitempty"`
}

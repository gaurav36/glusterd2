// Package peer implements the Peer type
package peer

import (
	"github.com/pborman/uuid"
)

// Peer reperesents a GlusterD
type Peer struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Addresses []string  `json:"addresses"`
}

// PeerAddRequest represents the structure to be added into the store
type PeerAddRequest struct {
	Addresses []string `json:"addresses"`
	Name      string   `json:"name,omitempty"`
}

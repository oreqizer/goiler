package ids

import (
	"github.com/oreqizer/go-relaygen/relay"
	"strconv"
)

// IntID holds information about an int ID
type IntID struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}

// FromGlobalIntID creates a local ID from a global int-based ID
func FromGlobalIntID(global string) (*IntID, error) {
	local := relay.FromGlobalID(global)
	id, err := strconv.Atoi(local.ID)
	if err != nil {
		return nil, err
	}

	return &IntID{ID: id, Type: local.Type}, nil
}

package ids

import (
	"github.com/oreqizer/go-relay"
	"strconv"
)

type IntID struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}

func FromGlobalIntID(global string) (*IntID, error) {
	local := relay.FromGlobalID(global)
	id, err := strconv.Atoi(local.ID)
	if err != nil {
		return nil, err
	}

	return &IntID{ID: id, Type: local.Type}, nil
}

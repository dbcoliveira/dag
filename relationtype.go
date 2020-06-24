package main

import (
	"strings"

	guuid "github.com/google/uuid"
)

// RelationshipStore structure
type RelationshipStore struct {
	ID guuid.UUID
	// Items Map of Relations
	items map[string]RelationType
}

// RelationType structure
type RelationType struct {
	ID          guuid.UUID
	Name        string
	Description string
	Bidirect    bool
}

// New Initialise RelationshipStore
func (r *RelationshipStore) New() {
	r.items = make(map[string]RelationType)
}

// NewRelation Create new Relation on store
func (r *RelationshipStore) NewRelation(name, description string, bidirect bool) {
	r.items[strings.ToUpper(name)] = RelationType{
		ID:          guuid.New(),
		Name:        strings.ToUpper(name),
		Description: description,
		Bidirect:    bidirect,
	}
}

// RelationType retrieve relation by NAME from store map
func (r *RelationshipStore) RelationType(name string) RelationType {
	return r.items[strings.ToUpper(name)]
}

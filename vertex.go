package main

import (
	"time"

	guuid "github.com/google/uuid"
)

// EntityType interface
type EntityType interface{}

// Vertex structure
type Vertex struct {
	ID     guuid.UUID
	Entity EntityType
	Edges  []*directedEdge
	cTime  time.Time
}

// New Vertex
func New(et EntityType) *Vertex {
	return &Vertex{
		ID:     guuid.New(),
		Entity: et,
		cTime:  time.Now(),
	}
}

func (v *Vertex) timeSinceCreation() time.Duration {
	return v.cTime.Sub(time.Now())
}

func (v *Vertex) existsEdge(e *directedEdge) bool {
	for _, re := range v.Edges {
		if (re.Relation.ID == e.Relation.ID) &&
			(re.Target == e.Target) {
			return true
		}
	}
	return false
}

func (v *Vertex) edgesCount(target *Vertex) int {
	return len(v.Edges)
}

func (v *Vertex) edgeByTarget(target *Vertex) []*directedEdge {
	var edgeList []*directedEdge
	for _, e := range v.Edges {
		if e.Target == target {
			edgeList = append(edgeList, e)
		}
	}
	return edgeList
}

func (v *Vertex) addEdgeTarget(rt RelationType, target *Vertex) {
	e := &directedEdge{
		ID:       guuid.New(),
		Relation: &rt,
		Target:   target,
		cTime:    time.Now(),
	}
	if !v.existsEdge(e) {
		v.Edges = append(v.Edges, e)
	} else {
		return
	}
	// add new edge to the target vertex (bidirectional edge)
	if rt.Bidirect {
		target.addEdgeTarget(rt, v)
	}
}

package main

import (
	"errors"
	"time"

	guuid "github.com/google/uuid"
)

type directedEdge struct {
	ID       guuid.UUID
	Relation *RelationType
	Target   *Vertex
	cTime    time.Time
}

func (e *directedEdge) addTarget(v *Vertex) error {
	if v != nil {
		e.Target = v
		return nil
	}
	return errors.New("No vertex")
}

func (e *directedEdge) creationTime() time.Time {
	return e.cTime
}

func (e *directedEdge) timeSinceCreation() time.Duration {
	return e.cTime.Sub(time.Now())
}

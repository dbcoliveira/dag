package main

import (
	"fmt"
	"reflect"

	guuid "github.com/google/uuid"
)

const (
	// ExportDisplayDOT Export graph in DOT format
	ExportDisplayDOT = "DOT"
	// ExportDisplayCYPHER Export graph in CYPHER format
	ExportDisplayCYPHER = "CYPHER"
)

// Search struct
type Search struct {
	visitedMap map[guuid.UUID]*Vertex
	vertex     *Vertex
	cyclic     bool
	display    string
}

// New Vertex
func (s *Search) New(baseVertex *Vertex) {
	s.visitedMap = make(map[guuid.UUID]*Vertex)
	s.vertex = baseVertex
	s.cyclic = false
}

// ShowGraph Show graph topology
func (s *Search) ShowGraph() {
	s.display = ExportDisplayCYPHER
	s.traverseGraph(s.vertex, "Name", true) // "Name" element default in s.vertex.Entity struct.
}

// ExportDOT export in DOT format
func (s *Search) ExportDOT(name string) {
	s.display = ExportDisplayDOT
	fmt.Println(fmt.Sprintf("digraph %s {", name))
	s.traverseGraph(s.vertex, "Name", true) // "Name" element default in s.vertex.Entity struct.
	fmt.Println(fmt.Sprintf("}"))
}

func (s *Search) graphDisplay(v *Vertex, e *directedEdge, field string) {
	fmt.Println(fmt.Sprintf("(%s {%s: \"%s\"}) - [:%s] -> (%s {%s: \"%s\"})",
		reflect.TypeOf(v.Entity).Elem().Name(),
		field,
		reflect.ValueOf(v.Entity).Elem().FieldByName(field), //Field name needs to be present
		e.Relation.Name,
		reflect.TypeOf(e.Target.Entity).Elem().Name(),
		field,
		reflect.ValueOf(e.Target.Entity).Elem().FieldByName(field))) //Field name needs to be present
}

func (s *Search) graphDisplayDOT(v *Vertex, e *directedEdge, field string) {
	fmt.Println(fmt.Sprintf("%s -> %s[label=%s];",
		reflect.ValueOf(v.Entity).Elem().FieldByName(field),        //Field name needs to be present
		reflect.ValueOf(e.Target.Entity).Elem().FieldByName(field), //Field name needs to be present
		e.Relation.Name))
}

func (s *Search) traverseGraph(v *Vertex, field string, display bool) {
	targetList := make(map[guuid.UUID]*Vertex) // Adjacent slice ?
	s.visitedMap[v.ID] = v
	for _, e := range v.Edges {
		targetList[e.Target.ID] = e.Target
		if display {
			switch s.display {
			case "DOT":
				s.graphDisplayDOT(v, e, field)
				break
			case "CYPHER":
				s.graphDisplay(v, e, field)
				break
			}
		}
	}
	for _, target := range targetList {
		if _, ok := s.visitedMap[target.ID]; !ok {
			s.traverseGraph(target, field, display)
		} else {
			s.cyclic = true
		}
	}
}

func (s *Search) isCyclic() bool {
	return s.cyclic
}

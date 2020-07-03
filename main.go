package main

import (
	"fmt"
	"os"
)

// Person Data structure
type Person struct {
	Name string
	Age  int
}

// Company Data structure
type Company struct {
	Name string
}

func main() {
	john:= New(&Person{
		Name: "John",
		Age:  37,
	})
	alice:= New(&Person{
		Name: "Alice",
		Age:  38,
	})
	doe := New(&Person{
		Name: "Doe",
		Age:  6,
	})

	eve := New(&Person{
		Name: "Eve",
		Age:  4,
	})

	bob := New(&Person{
		Name: "Bob",
		Age:  4,
	})

	foobar := New(&Company{
		Name: "Foobar",
	})

	foobargroup:= New(&Company{
		Name: "Foobar Group",
	})

	var RelationDB RelationshipStore
	RelationDB.New()

	RelationDB.NewRelation("Works", "Subject works for Target", false)
	RelationDB.NewRelation("Parent", "Subject is Parent of Target", false)
	RelationDB.NewRelation("Husband", "Subject is husband of Target", false)
	RelationDB.NewRelation("Household", "Subject lives with Target", true)
	RelationDB.NewRelation("Son", "Subject is Son of Target", false)
	RelationDB.NewRelation("Sibling", "Subject is Sibling of Target", true)
	RelationDB.NewRelation("Self", "Subject is itself (identity)", false)

	relationHusband := RelationDB.RelationType("Husband")
	relationHouseHold := RelationDB.RelationType("Household")
	relationParent := RelationDB.RelationType("Parent")
	relationSon := RelationDB.RelationType("Son")
	relationSibling := RelationDB.RelationType("Sibling")
	worksTo := RelationDB.RelationType("Works")
	itSelf := RelationDB.RelationType("Self")

	bob.addEdgeTarget(worksTo, foobargroup)
	foobar.addEdgeTarget(worksTo, bob)
	john.addEdgeTarget(worksTo, foobar)
	john.addEdgeTarget(relationHusband, alice)
	john.addEdgeTarget(relationHouseHold, alice)
	john.addEdgeTarget(relationParent, doe)
	alice.addEdgeTarget(relationParent, doe)
	alice.addEdgeTarget(relationParent, eve)
	alice.addEdgeTarget(itSelf, alice)
	doe.addEdgeTarget(relationSon, john)
	doe.addEdgeTarget(relationSon, alice)
	eve.addEdgeTarget(relationSibling, doe)
	eve.addEdgeTarget(relationSon, john)
	eve.addEdgeTarget(relationSon, alice)

	var search Search
	search.New(john)
	search.ExportDOT("My_exported_DOT_graph")
	fmt.Println("isCyclic():", search.isCyclic())
	fmt.Println(john.Edges[0].timeSinceCreation())
	os.Exit(0)
}

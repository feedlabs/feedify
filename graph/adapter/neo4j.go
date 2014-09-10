package adapter

import (
	"fmt"

	"github.com/feedlabs/feedify/neo4j"
	"github.com/feedlabs/feedify/graph/adapter/neo4jlang"
)

type Neo4jAdapter struct {
	client		*neo4j.Neo4jClient
	adapter		*neo4jlang.Neo4jCypher
}

func (n *Neo4jAdapter) Node() {}

func (n *Neo4jAdapter) Relation() {}

func (n *Neo4jAdapter) Query() {
	n.adapter.Query()
}

func (n *Neo4jAdapter) Connect() {
	db, err := n.client.Connect()
	if err != nil {
		fmt.Println("Cannot connect to neo4j database")
	}
	n.adapter.Db = db
}

func NewNeo4jAdapter(queryLanguage string) *Neo4jAdapter {
	client := neo4j.NewNeo4jClient()
	adapter := neo4jlang.NewNeo4jCypher()
	return &Neo4jAdapter{client, adapter}
}

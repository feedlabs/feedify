package adapter

import (
	"fmt"

	"github.com/jmcvetta/neoism"

	"github.com/feedlabs/feedify/neo4j"
)

type Neo4jAdapter struct {
	client		*neo4j.Neo4jClient
	queryLang	string
	db			  *neoism.Database
}

func (n *Neo4jAdapter) Node() {}

func (n *Neo4jAdapter) Relation() {}

func (n *Neo4jAdapter) Query() {
	if n.queryLang != "cypher" {
		fmt.Println("Unknown neo4j query language `" + n.queryLang + "`")
		return
	}

	cq := neoism.CypherQuery{
		Statement: `
			START n=node(*)
			MATCH (n)-[r:outranks]->(m)
			WHERE n.shirt = {color}
			RETURN n.name, type(r), m.name
			`,
		Parameters: neoism.Props{"color": "blue"},
		Result: &[]struct {
			N   string `json:"n.name"`
			Rel string `json:"type(r)"`
			M   string `json:"m.name"`
		}{},
	}

	n.db.Cypher(&cq)

	fmt.Println(cq.Result)
}

func (n *Neo4jAdapter) Connect() {
	db, err := n.client.Connect()
	if err != nil {
		fmt.Println("Cannot connect to neo4j database")
	}
	n.db = db
}

func NewNeo4jAdapter(queryLanguage string) *Neo4jAdapter {
	client := neo4j.NewNeo4jClient()
	return &Neo4jAdapter{client, queryLanguage, nil}
}

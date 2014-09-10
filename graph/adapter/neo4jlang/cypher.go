package neo4jlang

import (
	"fmt"

	"github.com/jmcvetta/neoism"
)

type Neo4jCypher struct {
	Db	*neoism.Database
}

func (n *Neo4jCypher) Query() {
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

	n.Db.Cypher(&cq)

	fmt.Println(cq.Result)
}

func NewNeo4jCypher() *Neo4jCypher {
	return &Neo4jCypher{}
}

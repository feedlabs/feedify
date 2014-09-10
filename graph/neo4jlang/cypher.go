package neo4jlang

import (
	"fmt"

	"github.com/jmcvetta/neoism"

	"github.com/feedlabs/feedify/neo4j"
	"github.com/feedlabs/feedify/graph"
)

const (
	NEO4J_CYPHER_PACKAGE_NAME = "neo4j-cypher"
)

func init() {
	graph.RegisterAdapterStore(NEO4J_CYPHER_PACKAGE_NAME, newAdapterStore, createNewNeo4jClient)
}

func createNewNeo4jClient(options graph.Options) error {
	return nil
}

func newAdapterStore(options graph.Options) (graph.GraphAdapterStore, error) {
	client := neo4j.NewNeo4jClient()
	return &GraphAdapterStore{client, nil}, nil
}

type GraphAdapterStore struct {
	client	*neo4j.Neo4jClient
	db		*neoism.Database
}

func (n *GraphAdapterStore) Query(query string) *graph.GraphQuery {
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

	return &graph.GraphQuery{}
}

func (n *GraphAdapterStore) Connect() {
	db, err := n.client.Connect()
	if err != nil {
		fmt.Println("Cannot connect to neo4j database")
	}
	n.db = db
}

func (n *GraphAdapterStore) Name() string {
	return NEO4J_CYPHER_PACKAGE_NAME
}

func (n *GraphAdapterStore) Disconnect() {}

func (n *GraphAdapterStore) Database(name string) *graph.GraphDatabase {
	return &graph.GraphDatabase{}
}

func (n *GraphAdapterStore) Node(id int) *graph.GraphNode {
	return &graph.GraphNode{}
}

func (n *GraphAdapterStore) Relation(id int) *graph.GraphRelation {
	return &graph.GraphRelation{}
}

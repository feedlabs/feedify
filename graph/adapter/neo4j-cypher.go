package adapter

import (
	"fmt"

	"github.com/jmcvetta/neoism"

	"github.com/feedlabs/feedify/neo4j"
	"github.com/feedlabs/feedify/graph"
	"github.com/feedlabs/feedify/graph/entity"
)

const (
	NEO4J_CYPHER_PACKAGE_NAME = "neo4j_cypher"
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

func (n *GraphAdapterStore) Query(statement string) *entity.GraphQuery {
	cq := neoism.CypherQuery{
		Statement: statement,
		Parameters: neoism.Props{"color": "blue"},
		Result: &[]struct {
			N   string `json:"n.name"`
			Rel string `json:"type(r)"`
			M   string `json:"m.name"`
		}{},
	}

	n.db.Cypher(&cq)

	fmt.Println(cq.Result)

	return &entity.GraphQuery{}
}

func (n *GraphAdapterStore) Connect() {
	db, err := n.client.Connect()
	if err != nil {
		fmt.Println("Cannot connect to neo4j database")
	}
	n.db = db
}

func (n *GraphAdapterStore) Disconnect() {}

func (n *GraphAdapterStore) Name() string {
	return NEO4J_CYPHER_PACKAGE_NAME
}

func (n *GraphAdapterStore) Database(name string) *entity.GraphDatabase {
	return &entity.GraphDatabase{}
}

func (n *GraphAdapterStore) Node(id int) *entity.GraphNode {
	return &entity.GraphNode{}
}

func (n *GraphAdapterStore) NewNode() *entity.GraphNode {
	return &entity.GraphNode{}
}

func (n *GraphAdapterStore) Relation(id int) *entity.GraphRelation {
	return &entity.GraphRelation{}
}

func (n *GraphAdapterStore) NewRelation() *entity.GraphRelation {
	return &entity.GraphRelation{}
}

func (n *GraphAdapterStore) FindNodes(params map[string]string) *entity.GraphNode {
	return &entity.GraphNode{}
}

func (n *GraphAdapterStore) FindRelations(params map[string]string) *entity.GraphRelation {
	return &entity.GraphRelation{}
}

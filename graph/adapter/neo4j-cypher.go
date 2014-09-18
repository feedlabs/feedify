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
	return &GraphAdapterStore{client, nil, false}, nil
}

type GraphAdapterStore struct {
	client	*neo4j.Neo4jClient
	db		*neoism.Database

	isConnected bool
}

func (n *GraphAdapterStore) Query(statement string) *entity.GraphQuery {
	if !n.isConnected {
		n.Connect()
	}

	cq := neoism.CypherQuery{
		Statement: statement,
		Parameters: neoism.Props{},
		Result: &[]struct {
			N   map[string]string `json:"n"`
//			N   string `json:"n.data"`
		}{},
	}

	n.db.Cypher(&cq)

	return &entity.GraphQuery{cq.Result}
}

func (n *GraphAdapterStore) Connect() {
	db, err := n.client.Connect()
	if err != nil {
		fmt.Println("Cannot connect to neo4j database")
	}
	n.isConnected = true
	n.db = db
}

func (n *GraphAdapterStore) Disconnect() {
	n.isConnected = false
}

func (n *GraphAdapterStore) Name() string {
	return NEO4J_CYPHER_PACKAGE_NAME
}

func (n *GraphAdapterStore) Database(name string) *entity.GraphDatabase {
	return &entity.GraphDatabase{}
}

func (n *GraphAdapterStore) Node(id int) (*entity.GraphNode, error) {
	if !n.isConnected {
		n.Connect()
	}

	_n, err := n.db.Node(id)

	if err != nil {
		return nil, err
	}

	return &entity.GraphNode{_n.Id(), _n.Data, _n.Extensions}, nil
}

func (n *GraphAdapterStore) NewNode(p graph.Props, label string) (*entity.GraphNode, error) {
	if !n.isConnected {
		n.Connect()
	}

	_n, err := n.db.CreateNode(neoism.Props(p))
	_n.AddLabel(label)

	if err != nil {
		return nil, err
	}

	return &entity.GraphNode{_n.Id(), _n.Data, _n.Extensions}, nil
}

func (n *GraphAdapterStore) DeleteNode(id int) (error) {
	if !n.isConnected {
		n.Connect()
	}

	_n, err := n.db.Node(id)

	if err != nil {
		return err
	}

	return _n.Delete()
}

func (n *GraphAdapterStore) SetPropertyNode(id int, key string, value string) (error) {
	if !n.isConnected {
		n.Connect()
	}

	_n, err := n.db.Node(id)

	if err != nil {
		return err
	}

	return _n.SetProperty(key, value)
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

func (n *GraphAdapterStore) FindNodesByLabel(label string) ([]*entity.GraphNode, error) {
	if !n.isConnected {
		n.Connect()
	}

	_n, err := n.db.NodesByLabel(label)

	if err != nil {
		return nil, err
	}

	var nodes []*entity.GraphNode

	for _, node := range _n {
		node := &entity.GraphNode{node.Id(), node.Data, node.Extensions}
		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (n *GraphAdapterStore) FindRelations(params map[string]string) *entity.GraphRelation {
	return &entity.GraphRelation{}
}

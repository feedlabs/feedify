package graph

import (
	"errors"
	"github.com/feedlabs/feedify/config"
)

type Neo4jAdapter struct {
	adapter	GraphAdapterStore
}

func (n *Neo4jAdapter) Node(id int) {
	n.adapter.Node(id)
}

func (n *Neo4jAdapter) Relation(id int) {
	n.adapter.Relation(id)
}

func (n *Neo4jAdapter) Query(query string) {
	n.adapter.Query(query)
}

func (n *Neo4jAdapter) Connect() {
	n.adapter.Connect()
}

func (n *Neo4jAdapter) Disconnect() {
	n.adapter.Disconnect()
}

func (m Neo4jAdapter) SetAdapter(adapter GraphAdapterStore) {
	m.adapter = adapter
}

func (m Neo4jAdapter) GetAdapter() GraphAdapterStore {
	return m.adapter
}

func NewNeo4jAdapter() (*Neo4jAdapter, error) {
	adapter_type := config.GetConfigKey("neo4j::adapter")
	adapter, err := NewAdapterStore("neo4j-" + adapter_type, nil)
	if err != nil {
		return nil, errors.New("Cannot load stream message adapter")
	}

	return &Neo4jAdapter{adapter}, nil
}

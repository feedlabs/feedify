package graph

import (
	"errors"
	"github.com/feedlabs/feedify/config"
	"github.com/feedlabs/feedify/graph/entity"
)

type GraphStorage struct {
	adapter GraphAdapterStore
}

func (g *GraphStorage) Connect() {
	g.adapter.Connect()
}

func (g *GraphStorage) Node(id int) (*entity.GraphNode, error) {
	return g.adapter.Node(id)
}

func (g *GraphStorage) DeleteNode(id int) (error) {
	return g.adapter.DeleteNode(id)
}

func (g *GraphStorage) SetPropertyNode(id int, key string, value string) (error) {
	return g.adapter.SetPropertyNode(id, key, value)
}

func (g *GraphStorage) NewNode(p Props, label string) (*entity.GraphNode, error) {
	return g.adapter.NewNode(p, label)
}

func (g *GraphStorage) RelateNodes(sourceNodeId int, destNodeId int, relationName string, p Props) (*entity.GraphRelation, error) {
	return g.adapter.RelateNodes(sourceNodeId, destNodeId, relationName, p)
}

func (g *GraphStorage) RelationshipsNode(id int, name ...string) ([]*entity.GraphRelation, error) {
	return g.adapter.RelationshipsNode(id, name...)
}

func (g *GraphStorage) DeleteRelation(id int) (error) {
	return g.adapter.DeleteRelation(id)
}

func (g *GraphStorage) FindNodesByLabel(label string) ([]*entity.GraphNode, error) {
	return g.adapter.FindNodesByLabel(label)
}

func (g *GraphStorage) Query(query string) *entity.GraphQuery {
	return g.adapter.Query(query)
}

func NewGraphStorage() (*GraphStorage, error) {
	adapter_type := config.GetConfigKey("graph::storage_adapter")
	adapter, err := NewAdapterStore(adapter_type, nil)
	if err != nil {
		return nil, errors.New("Cannot load graph storage adapter")
	}

	return &GraphStorage{adapter}, nil
}

package graph

import (
	"errors"
	"github.com/feedlabs/feedify/config"
	"github.com/feedlabs/feedify/graph/entity"
)

type GraphStorage struct {
	adapter GraphAdapterStore
}

func (g *GraphStorage) Node(id int) *entity.GraphNode {
	return g.adapter.Node(id)
}

func (g *GraphStorage) Connect() {
	g.adapter.Connect()
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

package service

import (
	"errors"
	"github.com/feedlabs/feedify/config"
	"github.com/feedlabs/feedify/graph"
)

type GraphService struct {}

func (s *GraphService) Name() string {
	return "graph-service"
}

func NewGraph() (graph.GraphAdapterStore, error) {
	graphAdapter := config.GetConfigKey("service::graph")

	if graphAdapter != "" {
		return graph.NewAdapterStore(graphAdapter, nil)
	}

	return nil, errors.New("Cannot load graph adapter '" + graphAdapter + "'")
}

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

func NewGraph() (*graph.Neo4jAdapter, error) {
	graphAdapter := config.GetConfigKey("service::graph")

	if graphAdapter == "neo4j" {
		return graph.NewNeo4jAdapter()
	}

	return nil, errors.New("Cannot load graph adapter")
}

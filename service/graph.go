package service

import (
	"github.com/feedlabs/feedify/config"
	"github.com/feedlabs/feedify/graph/adapter"
)

func NewGraph() *adapter.Neo4jAdapter {
	graphAdapter := config.GetConfigKey("service::graph")
	graphLanguage := config.GetConfigKey(graphAdapter + "::query")

	return adapter.NewNeo4jAdapter(graphLanguage)
}

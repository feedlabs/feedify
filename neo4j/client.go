package neo4j

import (
	"errors"

	"github.com/jmcvetta/neoism"

	"github.com/feedlabs/feedify/config"
)

type Neo4jClient struct {
	host	string
	port	string
	db		string
}

func (n *Neo4jClient) Connect() (*neoism.Database, error) {
	db, err := neoism.Connect("http://" + n.host + ":" + n.port + "/" + n.db)
	if err != nil {
		return nil, errors.New("Cannot connect to neo4j")
	}
	return db, nil
}

func NewNeo4jClient() *Neo4jClient {
	host := config.GetConfigKey("neo4j::host")
	port := config.GetConfigKey("neo4j::port")
	db := config.GetConfigKey("neo4j::db")

	return &Neo4jClient{host, port, db}
}

package entity

type GraphRelation struct {
	Id			int
	Type       	string      `json:"type"`
	Data       	interface{} `json:"data"`
	Extensions 	interface{} `json:"extensions"`

	StartNode	*GraphNode
	EndNode		*GraphNode
}

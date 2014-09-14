package entity

type GraphNode struct {
	relations *[]GraphRelation
}

func (n *GraphNode) GetRelation(int) (*GraphRelation) {
	return &GraphRelation{}
}

func (n *GraphNode) GetRelations(int) (*[]GraphRelation) {
	return n.relations
}

package graph

import (
	"errors"
	"github.com/barakmich/glog"
	"github.com/feedlabs/feedify/graph/entity"
)

// A Props is a set of key/value properties.
type Props map[string]interface{}

type GraphAdapterStore interface {
	Name() string

	Connect()
	Disconnect()

	Database(string) *entity.GraphDatabase

	Node(int) (*entity.GraphNode, error)
	DeleteNode(int) (error)
	SetPropertyNode(int, string, string) (error)
	NewNode(Props, string) (*entity.GraphNode, error)
	RelateNodes(int, int, string, Props) (*entity.GraphRelation, error)
	RelationshipsNode(int, ...string) ([]*entity.GraphRelation, error)

	Relation(int) *entity.GraphRelation
	NewRelation() *entity.GraphRelation
	DeleteRelation(int) (error)

	FindNodes(map[string]string) (*entity.GraphNode)
	FindNodesByLabel(string) ([]*entity.GraphNode, error)

	FindRelations(map[string]string) *entity.GraphRelation

	Query(string) *entity.GraphQuery
}

type Options map[string]interface{}

func (d Options) IntKey(key string) (int, bool) {
	if val, ok := d[key]; ok {
		switch vv := val.(type) {
		case float64:
			return int(vv), true
		default:
			glog.Fatalln("Invalid", key, "parameter type from config.")
		}
	}
	return 0, false
}

func (d Options) StringKey(key string) (string, bool) {
	if val, ok := d[key]; ok {
		switch vv := val.(type) {
		case string:
			return vv, true
		default:
			glog.Fatalln("Invalid", key, "parameter type from config.")
		}
	}
	return "", false
}

func (d Options) BoolKey(key string) (bool, bool) {
	if val, ok := d[key]; ok {
		switch vv := val.(type) {
		case bool:
			return vv, true
		default:
			glog.Fatalln("Invalid", key, "parameter type from config.")
		}
	}
	return false, false
}

type NewStoreFunc func(Options) (GraphAdapterStore, error)
type InitStoreFunc func(Options) error

type register struct {
	newFunc  NewStoreFunc
	initFunc InitStoreFunc
}

var storeRegistry = make(map[string]register)

func RegisterAdapterStore(name string, newFunc NewStoreFunc, initFunc InitStoreFunc) {
	if _, found := storeRegistry[name]; found {
		panic("already registered stream adapter " + name)
	}
	storeRegistry[name] = register{
		newFunc:  newFunc,
		initFunc: initFunc,
	}
}

func NewAdapterStore(name string, opts Options) (GraphAdapterStore, error) {
	r, registered := storeRegistry[name]
	if !registered {
		return nil, errors.New("adapterstore: name '" + string(name) + "' is not registered")
	}
	return r.newFunc(opts)
}

func InitAdapterStore(name string, opts Options) error {
	r, registered := storeRegistry[name]
	if registered {
		return r.initFunc(opts)
	}
	return errors.New("adapterstore: name '" + string(name) + "' is not registered")
}

func AdapterStores() []string {
	t := make([]string, 0, len(storeRegistry))
	for n := range storeRegistry {
		t = append(t, n)
	}
	return t
}

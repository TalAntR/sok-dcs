package dcs

import (
	"reflect"
)

type Node struct {
	descendants []*Node
	key         string
	value       interface{}
}

func (n *Node) Key() string {
	return n.key
}

func (n *Node) Resolve(path string) interface{} {
  	return nil
}

func Map2Tree(key string, m map[string]interface{}) *Node {
	return map2Tree(reflect.ValueOf(key), reflect.ValueOf(m))
}

func map2Tree(key reflect.Value, value reflect.Value) *Node {
	nodes := make([]*Node, 0)
	if reflect.Map == value.Kind() {
		i := 0
		nodes = make([]*Node, value.Len())
		for _, k := range value.MapKeys() {
			nodes[i] = map2Tree(k, value.MapIndex(k))
			i++
		}
	}
	return &Node{key: key.String(), value: value.Interface(), descendants: nodes}
}

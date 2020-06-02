package dcs

import (
	"reflect"
)

func (n *Node) Resolve(path []Label) interface{} {
	node := n.FindSubtree(path)
	if node == nil {
		return nil
	}
	return node.Value()
}

func Map2Tree(key string, m map[string]interface{}) Tree {
	return map2Tree(reflect.ValueOf(key), reflect.ValueOf(m))
}

func map2Tree(key reflect.Value, value reflect.Value) Tree {
	nodes := make([]Tree, 0)
	if reflect.Map == value.Kind() {
		i := 0
		nodes = make([]Tree, value.Len())
		for _, k := range value.MapKeys() {
			nodes[i] = map2Tree(k, value.MapIndex(k)).(*Node)
			i++
		}
	}
	return &Node{vertex{Label(key.String()), value.Interface()}, forest{nodes}}
}

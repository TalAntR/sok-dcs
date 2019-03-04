package dcs

import (
	"fmt"
	"reflect"
	"strings"
)

type Node struct {
	descendants []*Node
	key         string
	value       interface{}
}

func (n *Node) Key() string {
	return n.key
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Subtree(key string) *Node {
	for _, i := range n.descendants {
		if key == i.key {
			return i
		}
	}
	return nil
}

func (n *Node) Resolve(path string) interface{} {
	node := n.subtree(strings.Split(path, ":"))
	if node == nil {
		return nil
	}
	return node.value
}

func (n *Node) subtree(keys []string) *Node {
	node := n
	fmt.Printf("%v", keys)
	for _, key := range keys {
		fmt.Printf("%v", key)
		node = node.Subtree(key)
		if node == nil {
			break
		}
	}
	return node
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

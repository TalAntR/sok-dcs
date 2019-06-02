package dcs

// The Tree interface represents a configuration tree with labeled vertex.
type Tree interface {
	// The ChildLabels returns a collection of labels for descendants of the root vertex.
	Descendants() []Tree

	// The Label method returns label of the root vertex in the tree.
	Label() string

	// The Value method returns value associated with root vertex.
	Value() interface{}
}

// The TreeSearch interface represents a tree traversal algorithm.
type TreeSearch interface {

	// The Subtree method returns a child tree by specified key.
	Subtree(path []string) Tree
}

// The Vertex struct is a representation of a vertex in a tree.
type Vertex struct {
	key   string
	value interface{}
}

// Label method returns a key of a vertex;
func (v *Vertex) Label() string {
	return v.key
}

// Value method returns a value associated with vertex;
func (v *Vertex) Value() interface{} {
	return v.value
}

// The Node struct is a representation of a vertex in tree.
// Configuration tree is a labeled tree each node of them has some label
// and assosiated value.
type Node struct {
	Vertex
	descendants []*Node
}

// Descendants method returns an array of ...
func (n *Node) Descendants() []Tree {
	subtrees := make([]Tree, len(n.descendants))
	for i := range n.descendants {
		subtrees[i] = n.descendants[i]
	}
	return subtrees
}

// Subtree method returns ...
func (n *Node) Subtree(path []string) Tree {
	if len(path) < 1 {
		return n
	}
	t := n.getDescendant(path[0])
	if t != nil {
		return t.Subtree(path[1:])
	}
	return nil
}

func (n *Node) getDescendant(label string) *Node {
	for _, i := range n.descendants {
		if label == i.Label() {
			return i
		}
	}
	return nil
}

// MakePath function constructs a path (linear tree) from a collection of vertices
func MakePath(v0 Vertex, vertices ...Vertex) Tree {
	n := Node{v0, make([]*Node, 1)}
	p := &n
	for _, v := range vertices {
		n := Node{v, make([]*Node, 1)}
		p.descendants[0], p = &n, &n
	}
	p.descendants = []*Node{}
	return &n
}

// MergeTree function ...
func MergeTree(t ...Tree) Tree {
	size := len(t)
	if size == 1 {
		return t[0]
	}
	return nil
}

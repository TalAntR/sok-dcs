package dcs

// The Label type is used to mark Tree vertex
type Label string

type Pair struct {
	a, b interface{}
}

// The Vertex a data container for a tree.
type Vertex interface {

	// The Label method returns label or key of a vertex in tree.
	Label() Label

	// The Value method returns data associated with this vertex.
	Value() interface{}
}

type null struct{}

// Label method returns a key of a vertex;
func (v *null) Label() Label {
	return ""
}

// Value method returns a value associated with vertex;
func (v *null) Value() interface{} {
	return nil
}

// The Forest interface represents collection of labeled trees.
type Forest interface {

	// The GetTree retruns a tree which root vertex has given label
	// or nil if tree for this label is not found
	GetSubtree(label Label) Tree

	// The GetLabels returns array of root labels for trees in forest
	GetLabels() []Label
}

// The Tree interface represents a tree with labeled vertex.
type Tree interface {
	Forest
	Vertex
}

// The Vertex struct is a representation of a vertex in a tree.
type vertex struct {
	key   Label
	value interface{}
}

// The Label method returns a label of this vertex;
func (v *vertex) Label() Label {
	return v.key
}

// Value method returns a value associated with vertex;
func (v *vertex) Value() interface{} {
	return v.value
}

// V function creates a Vertex object for given label and value parameters
func V(label Label, value interface{}) Vertex {
	return &vertex{label, value}
}

type forest struct {
	descendants []Tree
}

// GetLabels retruns slice of labels for all child subtrees
func (f *forest) GetLabels() []Label {
	labels := make([]Label, len(f.descendants))
	for i, t := range f.descendants {
		labels[i] = t.Label()
	}
	return labels
}

// GetSubtree retruns child tree by given label
func (f *forest) GetSubtree(label Label) Tree {
	for _, t := range f.descendants {
		if label == t.Label() {
			return t
		}
	}
	return nil
}

// F function creates Forest object for given slice of trees
func F(descendants []Tree) Forest {
	return &forest{descendants}
}

// The Node struct is a representation of a vertex in tree.
// Configuration tree is a labeled tree each node of them has some label
// and assosiated value.
type Node struct {
	vertex
	forest
}

// FindSubtree method returns a tree by given path
func (n *Node) FindSubtree(path []Label) Tree {
	if len(path) < 1 {
		return n
	}
	t := n.GetSubtree(path[0])
	if t != nil {
		return t.(*Node).FindSubtree(path[1:])
	}
	return nil
}

// // DiffByLeafs ...
// func DiffByLeafs(t0 Tree, t1 Tree) Tree {
// 	root := Node{
// 		vertex{t1.Label(), Pair{t0.Value(), t1.Value()}},
// 		make([]*Node, 1),
// 	}

// 	labels := t1.GetLabels()
// 	for _, l := range labels {
// 		tree = t0.GetSubtree(l)
// 		if tree == nil {
// 			Node{v0, make([]*Node, 1)}
// 		} else {

// 		}
// 	}
// 	return 0
// }

func createBufferedLeaf(v Vertex, size int) *Node {
	root := vertex{v.Label(), v.Value()}
	childs := forest{make([]Tree, size)}
	return &Node{root, childs}
}

// MakePath constructs a path (linear tree) from a collection of vertices
func MakePath(v0 Vertex, vertices ...Vertex) Tree {
	root := createBufferedLeaf(v0, 1)
	item := root
	for _, v := range vertices {
		root := createBufferedLeaf(v, 1)
		item.descendants[0], item = root, root
	}
	item.descendants = []Tree{}
	return root
}

// MergeTree function ...
func MergeTree(t Tree, trees ...Tree) Forest {
	merged := []Tree{t}
	if len(trees) > 0 {
		for _, tree := range trees {
			merged = mergeTree(merged, tree)
		}
	}
	return F(merged)
}

func mergeTree(f []Tree, t Tree) []Tree {
	lidx := -1
	for i, tree := range f {
		if tree.Label() == t.Label() {
			lidx = i
		}
	}
	result := make([]Tree, len(f)+1)
	if lidx < 0 {
		copy(result, f)
		result[len(f)] = t
		return result
	}

	tree := f[lidx]
	subtrees := make([]Tree, len(tree.GetLabels()))
	for i, l := range tree.GetLabels() {
		subtrees[i] = tree.GetSubtree(l)
	}

	for _, l := range t.GetLabels() {
		subtrees = mergeTree(subtrees, t.GetSubtree(l))
	}
	r := createBufferedLeaf(t, len(subtrees))
	copy(r.forest.descendants, subtrees)
	return []Tree{r}
}

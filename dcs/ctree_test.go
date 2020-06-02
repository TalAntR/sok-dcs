package dcs

import (
	"testing"
)

func TestPathCompose(t *testing.T) {
	abc := MakePath(V("a", "a"), V("b", "ab"), V("c", "abc"))

	if abc.Label() != "a" {
		t.Errorf("Expected 'a' label is not equal to actual, got: '%s', want: '%s'.", abc.Label(), "a")
	}

	if abc.Value() != "a" {
		t.Errorf("Expected 'a' value is not equal to actual, got: '%s', want: '%s'.", abc.Value(), "a")
	}

	if len(abc.GetLabels()) != 1 {
		t.Errorf("Expected descendants number for 'a' is not equal to actual, got: '%d', want: '%d'.", len(abc.GetLabels()), 1)
	}

	bc := abc.GetSubtree("b")
	if string(bc.Label()) != "b" {
		t.Errorf("Expected 'b' label is not equal to actual, got: '%s', want: '%s'.", bc.Label(), "b")
	}

	if bc.Value() != "ab" {
		t.Errorf("Expected 'b' value is not equal to actual, got: '%s', want: '%s'.", bc.Value(), "ab")
	}

	if len(bc.GetLabels()) != 1 {
		t.Errorf("Expected descendants number for 'b' vertex is not equal to actual, got: '%d', want: '%d'.", len(bc.GetLabels()), 1)
	}

	c := bc.GetSubtree("c")
	if c.Label() != "c" {
		t.Errorf("Expected 'c' label is not equal to actual, got: '%s', want: '%s'.", c.Label(), "c")
	}

	if c.Value() != "abc" {
		t.Errorf("Expected 'c' value is not equal to actual, got: '%s', want: '%s'.", c.Value(), "abc")
	}

	if len(c.GetLabels()) != 0 {
		t.Errorf("Expected descendants number for 'c' vertex is not equal to actual, got: '%d', want: '%d'.", len(c.GetLabels()), 0)
	}
}

func TestMergeSingleTree(t *testing.T) {
	abc := MakePath(V("a", "a"), V("b", "ab"), V("c", "abc"))
	f := MergeTree(abc)
	if f == nil {
		t.Errorf("Expected forest after merge paths but not nil.")
	}

	if len(f.GetLabels()) != 1 {
		t.Errorf("Expected number of trees in forest is not equal to actual, got: '%d', want: '%d'.", len(f.GetLabels()), 1)
	}

	fabc := f.GetSubtree(abc.Label())
	if fabc.Value() != abc.Value() {
		t.Errorf("Expected value of the root element is not equal to actual, got: '%s', want: '%s'.", fabc.Value(), abc.Value())
	}

}

func TestMergePathsWithDifferentRoots(t *testing.T) {
	abc := MakePath(V("a", "a"), V("b", "ab"), V("c", "abc"))
	xyz := MakePath(V("x", "x"), V("y", "xy"), V("z", "xyz"))
	f := MergeTree(abc, xyz)
	if f == nil {
		t.Errorf("Expected an tree after merge paths but not nil.")
	}

	if len(f.GetLabels()) != 2 {
		t.Errorf("Expected number of trees in forest is not equal to actual, got: '%d', want: '%d'.", len(f.GetLabels()), 2)
	}

	fabc := f.GetSubtree(abc.Label())
	if fabc.Value() != abc.Value() {
		t.Errorf("Expected label of root element is not equal to actual, got: '%s', want: '%s'.", fabc.Value(), abc.Value())
	}

	fxyz := f.GetSubtree(xyz.Label())
	if fxyz.Value() != xyz.Value() {
		t.Errorf("Expected label of root element is not equal to actual, got: '%s', want: '%s'.", fxyz.Value(), xyz.Value())
	}
}

func TestMergePathsWithSameRoots(t *testing.T) {
	abc := MakePath(V("a", "a"), V("b", "ab"), V("c", "abc"))
	ayz := MakePath(V("a", "x"), V("y", "xy"), V("z", "xyz"))
	f := MergeTree(abc, ayz)
	if f == nil {
		t.Errorf("Expected an tree after merge paths but not nil.")
	}

	if len(f.GetLabels()) != 1 {
		t.Errorf("Expected number of trees in forest is not equal to actual, got: '%d', want: '%d'.", len(f.GetLabels()), 1)
	}

	fayz := f.GetSubtree(ayz.Label())
	if fayz.Value() != "x" {
		t.Errorf("Expected label of root element is not equal to actual, got: '%s', want: '%s'.", fayz.Value(), ayz.Value())
	}

	if len(fayz.GetLabels()) != 2 {
		t.Errorf("Expected number of trees in forest is not equal to actual, got: '%d', want: '%d'.", len(f.GetLabels()), 1)
	}

	fbc := fayz.GetSubtree("b")
	if fbc.Value() != "ab" {
		t.Errorf("Expected label of root element is not equal to actual, got: '%s', want: '%s'.", fbc.Value(), "ab")
	}

	fyz := fayz.GetSubtree("y")
	if fyz.Value() != "xy" {
		t.Errorf("Expected label of root element is not equal to actual, got: '%s', want: '%s'.", fyz.Value(), "xy")
	}
}

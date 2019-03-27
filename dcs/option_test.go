package dcs

import "testing"

func TestMap2Tree(t *testing.T) {
	m := make(map[string]interface{})

	m["a"] = "va"
	m["b"] = "vb"

	n := Map2Tree(":", m)
	if n.Label() != ":" {
		t.Errorf("Expected root key is not equal to actual, got: '%s', want: '%s'.", n.Label(), ":")
	}

	a := n.Subtree("a")
	if a.Value() != "va" {
		t.Errorf("Expected value for 'a' key is not equal to actual, got: '%s', want: '%s'.", a.Value(), "va")
	}

	bv := n.Resolve("b")
	if bv != "vb" {
		t.Errorf("The b element value is not equal to expected one, got: '%s', want: '%s'.", bv, "vb")
	}
}

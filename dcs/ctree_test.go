package dcs

import "testing"

func TestTree(t *testing.T) {
	m := make(map[string]interface{})

	m["a"] = "va"
	m["b"] = "vb"

	n := Map2Tree(":", m)
	if n.Label() != ":" {
		t.Errorf("Expected root key is not equal to actual, got: '%s', want: '%s'.", n.Label(), ":")
	}

}

package dcs

import "testing"

func TestMap2Tree(t *testing.T) {
	m := make(map[string]interface{})

	m["a"] = "va"
	m["b"] = "vb"

	n := Map2Tree("/", m)
	if n.getKey() != "test" {
		t.Errorf("Expected root key is not equal to actual, got: '%s', want: '%s'.", n.getKey(), "test")
	}
}

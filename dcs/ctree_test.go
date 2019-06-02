package dcs

import (
	"testing"
)

func TestPathCompose(t *testing.T) {
	abc := MakePath(Vertex{"a", "a"}, Vertex{"b", "ab"}, Vertex{"c", "abc"})

	if abc.Label() != "a" {
		t.Errorf("Expected 'a' label is not equal to actual, got: '%s', want: '%s'.", abc.Label(), "a")
	}

	if abc.Value() != "a" {
		t.Errorf("Expected 'a' value is not equal to actual, got: '%s', want: '%s'.", abc.Value(), "a")
	}

	if len(abc.Descendants()) != 1 {
		t.Errorf("Expected descendants number for 'a' is not equal to actual, got: '%d', want: '%d'.", len(abc.Descendants()), 1)
	}

	bc := abc.Descendants()[0]
	if bc.Label() != "b" {
		t.Errorf("Expected 'b' label is not equal to actual, got: '%s', want: '%s'.", bc.Label(), "b")
	}

	if bc.Value() != "ab" {
		t.Errorf("Expected 'b' value is not equal to actual, got: '%s', want: '%s'.", bc.Value(), "ab")
	}

	if len(bc.Descendants()) != 1 {
		t.Errorf("Expected descendants number for 'b' vertex is not equal to actual, got: '%d', want: '%d'.", len(bc.Descendants()), 1)
	}

	c := bc.Descendants()[0]
	if c.Label() != "c" {
		t.Errorf("Expected 'c' label is not equal to actual, got: '%s', want: '%s'.", c.Label(), "c")
	}

	if c.Value() != "abc" {
		t.Errorf("Expected 'c' value is not equal to actual, got: '%s', want: '%s'.", c.Value(), "abc")
	}

	if len(c.Descendants()) != 0 {
		t.Errorf("Expected descendants number for 'c' vertex is not equal to actual, got: '%d', want: '%d'.", len(c.Descendants()), 0)
	}
}

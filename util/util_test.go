package util

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicateFromSlice(t *testing.T) {
	checkExpectedValue(
		t,
		[]string{"a", "b", "a"},
		[]string{"a", "b"},
	)

	checkExpectedValue(
		t,
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c"},
	)

	checkExpectedValue(
		t,
		[]string{"a", "a", "a", "a"},
		[]string{"a"},
	)

	var emptySlice []string
	checkExpectedValue(
		t,
		[]string{},
		emptySlice,
	)
}

func checkExpectedValue(t *testing.T, args []string, expected []string) {
	removed := RemoveDuplicateFromSlice(args)
	if reflect.DeepEqual(removed, expected) == false {
		t.Fatalf("Should return %q, but got %q", expected, removed)
	}
}

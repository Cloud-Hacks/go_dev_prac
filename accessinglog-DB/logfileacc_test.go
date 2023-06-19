package main

import (
	"reflect"
	"testing"
)

func TestFindTopNMenuItems(t *testing.T) {
	menuItemCounts := map[string]int{
		"Item1": 10,
		"Item2": 5,
		"Item3": 8,
		"Item4": 3,
		"Item5": 12,
	}

	expected := []string{"Item5", "Item1", "Item3"}

	topNMenuItems := findTopNMenuItems(menuItemCounts, 3)

	if !reflect.DeepEqual(topNMenuItems, expected) {
		t.Errorf("Mismatched top N menu items. Expected: %v, Got: %v", expected, topNMenuItems)
	}
}

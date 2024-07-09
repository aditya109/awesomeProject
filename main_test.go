package main

import (
	"testing"
)

func TestExpandCronField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		minValue int
		maxValue int
		expected []int
	}{
		{"Test star field", "*", 0, 5, []int{0, 1, 2, 3, 4, 5}},
		{"Test specific values", "1,3,5", 0, 5, []int{1, 3, 5}},
		{"Test range", "1-3", 0, 5, []int{1, 2, 3}},
		{"Test step values", "*/2", 0, 5, []int{0, 2, 4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cf := NewCronField(test.name, test.field, test.minValue, test.maxValue)
			cf.Expand()
			if !equalSlices(cf.AllowedVals, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, cf.AllowedVals)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

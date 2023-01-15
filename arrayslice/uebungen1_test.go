package main

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	l1 := []int{1, 3, 5, 7, 9}
	var tests = []struct {
		a    []int
		x    int
		want bool
	}{
		{l1, 1, true},
		{l1, 5, true},
		{l1, 9, true},
		{l1, 4, false},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Contains(tt.a, tt.x)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

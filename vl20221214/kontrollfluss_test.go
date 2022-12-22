// Aufgaben zu Funktionen, Schleifen und If-Then-Else
package main

import (
	"fmt"
	"testing"
)

func TestDouble(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{0, 0},
		{1, 2},
		{-2, -4},
		{-1, -2},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := double(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{1, 1},
		{2, 3},
		{5, 15},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := sum(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestSumMultiples(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{3, 10, 18},
		{2, 10, 20},
		{10, 2, 0},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := sumMultiples(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	var tests = []struct {
		a    int
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{31, true},
		{101, true},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := isPrime(tt.a)
			if ans != tt.want {

				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestLcm(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{3, 5, 15},
		{2, 5, 10},
		{4, 10, 20},
		{20, 5, 20},
		{25, 10, 50},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := lcm(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestBinrep(t *testing.T) {
	var tests = []struct {
		a    int
		want string
	}{
		{5, "101"},
		{50, "110010"},
		{9000, "10001100101000"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := binrep(tt.a)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

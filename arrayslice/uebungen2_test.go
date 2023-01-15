package main

import (
	"fmt"
	"testing"
)

func TestContain1s(t *testing.T) {
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

func ExampleContainsChain2_nonempty() {
	l1 := []int{1, 5, 6, 2, 7, 9, 5}
	l2 := []int{1, 7, 3, 5, 4, 5, 6}
	l3 := []int{1, 3, 2, 5, 4, 9, 5}

	fmt.Println(ContainsChain2(l1))
	fmt.Println(ContainsChain2(l2))
	fmt.Println(ContainsChain2(l3))

	// Output:
	// true
	// true
	// false

}

func ExampleContainsChain2_short() {
	l1 := []int{2, 3}
	l2 := []int{1, 3, 5}

	fmt.Println(ContainsChain2(l1))
	fmt.Println(ContainsChain2(l2))

	// Output:
	// false
	// true
}

func ExampleContainsChain2_empty() {
	l1 := []int{}

	fmt.Println(ContainsChain2(l1))

	// Output:
	// false
}

func ExampleContainsSum_nonempty() {
	l1 := []int{1, 14, 14, 14, 2, 23, 5}
	l2 := []int{1, 3, 5, 14, 2, 23, 17}
	l3 := []int{1, 7, 2, 3, 5, 8, 1, 15, 2}

	fmt.Println(ContainsSum(l1))
	fmt.Println(ContainsSum(l2))
	fmt.Println(ContainsSum(l3))

	// Output:
	// true
	// true
	// false
}

func ExampleContainsSum_short() {
	l1 := []int{2, 3}
	l2 := []int{14, 14, 14}

	fmt.Println(ContainsSum(l1))
	fmt.Println(ContainsSum(l2))

	// Output:
	// false
	// true
}

func ExampleContainsSum_empty() {
	l1 := []int{}

	fmt.Println(ContainsSum(l1))

	// Output:
	// false
}

package filter

import "fmt"

func ExampleFilterLess() {
	l1 := []int{10, 3, 17, 5, 23, 12}

	lesser := FilterLess(l1, 11)
	fmt.Println(lesser)

	// Output:
	// [10 3 5]
}

func ExampleFilterGreater() {
	l1 := []int{10, 3, 17, 5, 23, 12}

	greater := FilterGreater(l1, 11)
	fmt.Println(greater)

	// Output:
	// [17 23 12]
}

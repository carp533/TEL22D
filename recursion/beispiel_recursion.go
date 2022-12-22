package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

/*Aufgabe programmiere die Ackermann Funktion.
  A(m, n) =
  n+1                              if m = 0
  A(m-1, 1)                        if m > 0 and n = 0
  A(m-1, A(m, n-1))                if m > 0 and n > 0
*/

func ackerman(m, n int) int {
	//TODO
	return 0
}

func main() {
	fmt.Println(fact(7))
	fmt.Println(fib(7))
	fmt.Println(ackerman(3, 2)) //->29
}

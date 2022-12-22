package main

import "fmt"

func main() {

	//definiert einen array der Länge 5 vom Typ int
	var a [5]int
	fmt.Println("emp:", a)

	//Zugriff auf Elemente des Array
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//eingebaute funktion len gibt Länge des Array zurück
	fmt.Println("len:", len(a))

	//Kurzform Variable deklarieren und Werte zuweisen
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//Zweidimensionaler Array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

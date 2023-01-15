package main

import "fmt"

func BeispielSlice() {

	//slices sind eine Referenz auf ein Array
	// Eingebaute Funktion make um ein slice zu erzeugen
	s := make([]string, 3)
	fmt.Println("emp:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	//Eingebaute Funktion append um ein Element zu einem sclice hinzuzufügen
	s = append(s, "d")
	s = append(s, "e", "f")
	s = append(s, "g", "h")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	//Eingabute Funktion copy um einen sclice zu kopieren
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	c[0] = "X"
	fmt.Println("c = ", c)
	fmt.Println("s = ", s)

	//"sclice" Operator -> Zugriff auf Elemente 2,3,4
	// der linke Index (2) ist inklusiv; der rechte (5) exclusiv
	l := s[2:5]
	fmt.Println("sl1:", l)

	//"sclice" Operator -> Zugriff auf Elemente 0 bis 4
	l = s[:5]
	fmt.Println("sl2:", l)

	//"sclice" Operator -> Zugriff auf Elemente 2 bis Ende
	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

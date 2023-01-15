package main

import (
	"fmt"
)

type Kunde struct {
	name    string
	vorname string
	kdnr    string
	punkte  int
}

func main() {

	// Zwei einzelne Variablen vom Typ Kunde definieren.
	k2 := Kunde{
		"Beispiel",
		"Barbara",
		"456a",
		77,
	}
	k3 := Kunde{
		punkte:  77,
		vorname: "Monika",
		kdnr:    "456a",
		name:    "Musterfrau",
	}

	fmt.Println(k2)
	fmt.Println(k3)

	// Eine Liste von Kunden:
	l1 := []Kunde{k2, k3, Kunde{"Mustermann", "Max", "123a", 42}}

	fmt.Println(l1)

	// Kundennummer des 2. Kunden ausgeben:
	kunde2 := l1[1]
	kdnr_kunde2 := kunde2.kdnr
	fmt.Println(kdnr_kunde2)

}

// Aufgaben zu Funktionen, Schleifen und If-Then-Else
// in der Datei kontrollfluss.go sind Tests für die Funktionen enthalten.
// prüfe deine Funktionen mit "go test" oder "go test -v" aus dem Terminal
package main

import (
	"fmt"
)

// Liefert das Doppelte von n.
func double(n int) int {
	// TODO: Hier das Ergebnis berechnen.
	return 0
}

// Liefert die Summe der Zahlen von 1 bis n.
func sum(n int) int {
	// TODO: Hier das Ergebnis berechnen.
	return 0
}

// Liefert die Summe aller Vielfachen von x, die kleiner als n sind.
func sumMultiples(x, n int) int {
	result := 0
	// TODO: Hier das Ergebnis berechnen.
	return result
}

// Liefert genau dann true, wenn n eine Primzahl ist.
func isPrime(n int) bool {
	// TODO: Das pauschale Return durch etwas sinnvolles ersetzen.
	return true
}

// Liefert den größten Primfaktor der Zahl n.
func largestPrimeFactor(n int) int {
	// TODO: Hier das Ergebnis berechnen.
	return 0
}

// Liefert das kleinste gemeinsame Vielfache von m und n.
func lcm(m, n int) int {
	// TODO: Hier das Ergebnis berechnen.
	return 0
}

// binär Darstellung
// Schreibe eine Funktion, die zu einer ganzen Zahl die binär Darstellung der Zahl ausgibt.
// die Funtkion fmt.Sprintf sollte nicht verwendet werden.
// Beispiel f(5) = 101
// Beispiel f(50) = 110010
// Beispiel f(9000) = 10001100101000
func binrep(n int) string {
	return fmt.Sprintf("%b\n", n)
}

// Main-Funktion
// Rufen Sie hier Ihre Funktionen auf, die Sie testen möchten
func main() {
	//fmt.Println(lcm(25, 10)) // Erwarte:   50
	fmt.Println(binrep(9000))
}

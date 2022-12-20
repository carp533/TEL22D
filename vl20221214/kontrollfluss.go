// Aufgaben zu Funktionen, Schleifen und If-Then-Else
// in der Datei kontrollfluss.go sind Tests für die Funktionen enthalten.
// prüfe deine Funktionen mit "go test" oder "go test -v" aus dem Terminal
package main

import (
	"fmt"
	"math"
)

// Liefert das Doppelte von n.
func double(n int) int {
	// TODO: Hier das Ergebnis berechnen.
	//return 0
	return 2 * n
}

// Liefert die Summe der Zahlen von 1 bis n.
func sum(n int) int {
	// TODO: Hier das Ergebnis berechnen.
	//return 0
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		return sum
	}
	return sum
}

// Liefert die Summe aller Vielfachen von x, die kleiner als n sind.
func sumMultiples(x, n int) int {
	result := 0
	// TODO: Hier das Ergebnis berechnen.
	for i := x; i < n; i += x {
		result += i
	}
	return result
}

// Liefert genau dann true, wenn n eine Primzahl ist.
func isPrime(n int) bool {
	// TODO: Das pauschale Return durch etwas sinnvolles ersetzen.
	if n < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq_root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Liefert den größten Primfaktor der Zahl n.
func largestPrimeFactor(n int) int {
	// TODO: Hier das Ergebnis berechnen.
	//return 0
	if n < 2 {
		return 0
	}
	lpf := 0
	for n%2 == 0 {
		lpf = 2
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			lpf = i
			n = n / i
		}
	}
	if n > 2 {
		lpf = n
	}
	return lpf
}

// Liefert das kleinste gemeinsame Vielfache von m und n.
func lcm(m, n int) int {
	// TODO: Hier das Ergebnis berechnen.
	//return 0
	for i := 2; i <= n*m; i++ {
		if i%n == 0 && i%m == 0 {
			return i
		}
	}
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
	fmt.Println(sum(5))
}

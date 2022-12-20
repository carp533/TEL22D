package main

import (
	"fmt"
	//"reflect"
)

/* Aufgabe 1:
 *
 * Schreiben Sie eine Funktion, die den Benutzer immer wieder nach einer Zahl fragt,
 * bis er eine 0 eingibt. Anschließend soll die Summe der eingegebenen Zahlen
 * auf die Konsole ausgegeben werden.
 *
 * Hinweis: In der main()-Funktion steht ein Beispiel, wie man eine einzelne Zahl
 *          einliest und wieder ausgibt.
 */
func readSum() int {
	// TODO
	return 0
}

/* Aufgabe 2:
 *
 * Schreiben Sie eine Funktion, die den Benutzer immer wieder nach einer Zahl fragt,
 * bis er eine 0 eingibt. Anschließend soll die Anzahl der eingegebenen Fünfen
 * auf die Konsole ausgegeben werden.
 */
func readFives() int {
	// TODO
	return 0
}

/* Aufgabe 3:
 *
 * Verallgemeinern Sie die Funktion readFives() so, dass man die zu zählende Zahl
 * als Parameter übergeben kann.
 */

func main() {

	// Beispiel: Eine Zahl einlesen und sie wieder ausgeben:
	var input int
	fmt.Print("Bitte eine Zahl eingeben: ")
	fmt.Scanln(&input)
	fmt.Printf("Sie haben %v eingegeben.\n", input)

	// Hier wird die Funktion readSum() aufgerufen.
	readSum()

	// Hier wird die Funktion readFives() aufgerufen.
	readFives()
}

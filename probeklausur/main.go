package main

import (
	"fmt"
	"strings"
)

/* AUFGABE: Listen
 * Erreichbare Punkte: 10
 */

// AUFGABENSTELLUNG:
// Erwartet drei Zahlen m,n,max > 0 und liefert eine Liste mit allen
// gemeinsamen Vielfachen von m und n, die nicht größer als max sind.
func CommonMultiples(m, n, max int) []int {
	result := []int{}
	// TODO
	return result
}

/* AUFGABE: Listen
 * Erreichbare Punkte: 10
 */

// AUFGABENSTELLUNG:
// Erwartet eine int-Slice list.
// Liefert eine int-Slice, die an Stelle n die Summe
// der Elemente aus list bis zu Stelle n enthält.
func ArraySums(list []int) []int {
	result := []int{}
	// TODO
	return result
}

/* AUFGABE: Rekursion
 * Erreichbare Punkte: 10
 */

// AUFGABENSTELLUNG:
// Liefert die Potenz "2 hoch x".
// Zusatzanforderung: Die Funktion muss rekursiv sein.
func Power2(x int) float64 {
	// TODO
	return 0
}

/* AUFGABE: Listen
 * Erreichbare Punkte: 10
 */

// AUFGABENSTELLUNG:
// Liefert eine Liste mit allen Stellen in list,
// an denen x vorkommt.
func FindAll(list []int, x int) []int {
	result := []int{}
	// TODO
	return result
}

/* AUFGABE: Structs
 * Erreichbare Punkte: 20
 */

// AUFGABENSTELLUNG:
// Reparieren Sie die folgende Methode.
// Es müssen alle Tests grün werden.

// Fügt ein Element zu dict hinzu,
// das das Wortpaar (de,en) enthält.
// Falls dict schon einen Eintrag für de enthält,
// soll dessen en ersetzt werden.
func (dict *Dictionary) SetEntry(de, en string) {
	dict.Entries = append(dict.Entries, Entry{de, en})
}

type Entry struct {
	De string
	En string
}

func (entry Entry) String() string {
	return fmt.Sprintf("%s : %s", entry.De, entry.En)
}

type Dictionary struct {
	Entries []Entry
}

func (dict Dictionary) String() string {
	lines := []string{}
	for i, entry := range dict.Entries {
		lines = append(lines, fmt.Sprintf("%d - %s", i, entry))
	}
	return strings.Join(lines, "\n")
}
func ExampleDictionary_SetEntry() {
	d1 := Dictionary{[]Entry{}}

	d1.SetEntry("Himmel", "sky")
	fmt.Println(d1)
	fmt.Println()

	d1.SetEntry("Erde", "earth")
	fmt.Println(d1)
	fmt.Println()

	d1.SetEntry("Himmel", "heaven")
	fmt.Println(d1)
	fmt.Println()

	// Output:
	// 0 - Himmel : sky
	//
	// 0 - Himmel : sky
	// 1 - Erde : earth
	//
	// 0 - Himmel : heaven
	// 1 - Erde : earth
}

/* AUFGABE: Structs
 * Erreichbare Punkte: 20
 */

// AUFGABENSTELLUNG:
// Im Folgenden ist ein Datentyp für die Einträge eines Wörterbuchs gegeben.
// Ändern Sie diesen Datentyp so ab, dass er anstelle eines einzelnen englischen
// Worts eine Liste von Wörtern enthält.
// Passen Sie die Funktionen NewEntry(), Translations() und AddTranslation() an.

// Datentyp für Einträge eines Wörterbuchs.
type Eintrag struct {
	De string
	En string
}

// Liefert einen neuen Eintrag.
func NewEntry(de, en string) Entry {
	return Entry{de, en}
}

// Liefert eine String-Repräsentation eines Eintrags.
func (entry Eintrag) String() string {
	return fmt.Sprintf("%s : %v", entry.De, entry.Translations())
}

// Liefert einen String mit allen englischen Wörtern aus entry.
// Die einzelnen Wörter sollen mit Kommata getrennt sein.
func (entry Eintrag) Translations() string {
	return entry.En
}

// Fügt eine neue Übersetzung zu entry hinzu.
func (entry *Eintrag) AddTranslation(newEn string) {
	entry.En = newEn
}

/* AUFGABE: Rekursion
 * Erreichbare Punkte: 15
 */

// AUFGABENSTELLUNG:
// Reparieren Sie die folgende Funktion.
// Es müssen alle Tests grün werden.
// Zusatzanforderung: Die Funktion muss rekursiv bleiben.

// Berechnet den ganzzahligen Quotienten x / y.
// D.h. die Funktion ignoriert Nachkommastellen bzw. den Rest.
func Div(x, y int) int {
	return 1 + Div(x-y, y)
}

/* AUFGABE: Schere-Stein-Papier
 * Erreichbare Punkte: 25
 */

// AUFGABENSTELLUNG:
// Hier ist ein Gerüst für das Spiel "Schere-Stein-Papier" vorgegeben.
// Vervollständigen Sie das Spiel, d.h. implementieren Sie die leeren Funktionen.

// Spielt eine Spielrunde Schere-Stein-Papier.
func PlayRPS() {
	inputP1 := ChooseItem(1)
	inputP2 := ChooseItem(2)

	result := GetResult(inputP1, inputP2)
	PrintResult(result)
}

// Fragt den Spieler mit der gegebenen Nummer
// nach seiner Wahl und liefert sie als int.
func ChooseItem(player int) int {
	// TODO
	return 0
}

// Bestimmt das Ergebnis des Spiels.
// 0: Unentschieden
// 1: Spieler 1 gewinnt.
// 2: Spieler 2 gewinnt.
func GetResult(input1, input2 int) int {
	// TODO
	return 0
}

// Gibt das Ergebnis auf der Konsole aus.
func PrintResult(result int) {
	switch result {
	case 0:
		fmt.Println("Draw.")
	case 1:
		fmt.Println("Player 1 wins.")
	case 2:
		fmt.Println("Player 2 wins.")
	}
}

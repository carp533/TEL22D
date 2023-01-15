package main

import (
	"fmt"
)

func f(s string) {
	s = "Funktion f"
	fmt.Println(s)
}

func fr(p *string) {
	*p = "Funktion fr"
	fmt.Println(*p)

}
func main() {

	i, j := 42, 2701

	p := &i                                        // definiert einen Pointer auf die Variable i
	*p = 21                                        // (Schreib-)Zugriff auf den Wert des Pointer (Inhalt Speicher)
	fmt.Printf("Der Wert von    p (*p): %v\n", *p) // (Lese-)Zugriff auf den Wert des Pointer (Inhalt Speicher)
	fmt.Printf("Die Adresse von p (p) : %v\n", p)  //
	fmt.Printf("Der Wert von    i (i) : %v\n", i)  //
	fmt.Printf("Die Adresse von i (&i): %v\n", &i) // Zugriff auf den Speicher, in dem i steht
	p = &j                                         // p zeigt nun auf die Variable j
	*p = *p / 37                                   // ändere den Wert (Inhalt Speicher)
	fmt.Println(j)

	var k *int
	fmt.Printf("%v\n", k) //ein initialer Pointer ist <nil> -> hiermit kann man nichts machen
	//*k = 42             // -> Programm bricht ab, weil k noch auf keinen Speicherbereich zeigt
	//fmt.Printf("%v", k)

	str := "test"
	f(str)
	fmt.Println(str)

	fr(&str)
	fmt.Println(str)
	main2()

}

func main2() {
	// Beispiel 1: Definition einer Variable und eines Pointers darauf.
	//             Eine Veränderung der Variable ändert auch den Pointer und umgekehrt.

	// Die Variablen erzeugen und einmal ausgeben.
	v1 := 42
	p_v1 := &v1
	fmt.Println(v1, p_v1, *p_v1) // Gibt den Wert von v1 sowie die Adresse und den Wert von p_v1 aus.

	v1 = 23
	fmt.Println(v1, p_v1, *p_v1) // Es haben sich sowohl v1 als auch *p_v1 verändert.

	*p_v1 = 77
	fmt.Println(v1, p_v1, *p_v1) // Es haben sich sowohl v1 als auch *p_v1 verändert.

	// Beispiel 2: Eine Funktion nutzen, die einen Pointer erwartet und ihn verändert.
	changePtr(p_v1)
	fmt.Println(v1, p_v1, *p_v1) // Es haben sich sowohl v1 als auch *p_v1 verändert.

	// Beispiel 3: Einen Pointer mit new erzeugen.
	p2 := new(int)
	fmt.Println(p2, *p2) // Adresse und Wert von p2 ausgeben.
	changePtr(p2)
	fmt.Println(p2, *p2) // Adresse und Wert von p2 ausgeben.

	// Beispiel 4: Das Struct unten instanziieren und beide Methoden ausprobieren.
	s1 := MyStruct{10}
	fmt.Println(s1)
	s1.change()
	fmt.Println(s1)
	s1.noChange()
	fmt.Println(s1)

}

// Beispiel-Funktion: Erwartet einen int-Pointer und ändert den Wert dahinter.
func changePtr(i_ptr *int) {
	*i_ptr += 10
}

// Beispiel-Struct mit jeweils einer Methode, die es verändert und einer, die es nicht verändert.
type MyStruct struct {
	i1 int
}

func (m *MyStruct) change() {
	m.i1 += 100
}
func (m MyStruct) noChange() {
	m.i1 += 100 // Hat keinen Effekt, weil m nicht als Pointer vorliegt.
}

## Informatik 1 TEL22D
Hier finden Sie die Programme zur Vorlesung informatik 1. Dieses GIT Repository in das workspace Verzeichnis kopieren/auschecken.
(über GIT Zip Datei herunter laden oder git clone https://github.com/carp533/TEL22D)

Anleitung eigenes Modul in den workspace einbinden:

<in workspace Verzeichnis wechseln>
$ mkdir myfirstmodule
$ cd myfirstmodule
$ go mod init myfirstmodule

<eine go Datei erzeugen myfirst.go>

package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello")
}

<datei starten>
$ go run myfirst.go

<zurück in das workspace Verzeichnis wechseln>
$ cd ..

<das modul zum workspace hinzufügen>
$ go work use .\myfirstmodule\

<falls Testdateien vorhanden sind (*_test.go) können die Tests mit "go test" oder "go test -v" ausgeführt werden. oder über VS Code in der Testdatei (Zeile 1) run package tests oder run file tests>

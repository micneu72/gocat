package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	for _, fname := range os.Args[1:] {
		fmt.Println(fname)

		if len(os.Args) == 1 {
			fmt.Println("Mindestens eine Datei als Parameter erwartet")
			os.Exit(1)
		}
		f, err := os.Open(fname)
		if err != nil {
			log.Printf("Fehler beim Öffnen der Datei: %s/n%s", fname, err)
			f.Close()
			continue
		}
		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			log.Printf("Fehler bei der Ausgabe von %s\n%s", fname, err)
		}
		f.Close()
	}
}

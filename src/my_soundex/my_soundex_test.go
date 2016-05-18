package main

import (
	"log"
	"testing"
)

func TestCalculateSoundex(t *testing.T) {
	log.SetFlags(0)
	log.Println("TEST calculateSoundex()")

	word := &Word{Value: "Soundex"}
	word.calculateSoundex()
	if word.Soundex != "S532" {
		t.Fatalf("%q != %q", word.Soundex, "S532")
	}

	word.Value = "Beth"
	word.calculateSoundex()
	if word.Soundex != "B300" {
		t.Fatalf("%q != %q", word.Soundex, "B300")
	}
}

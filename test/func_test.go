package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestNewline(t *testing.T) {
	tab := []string{"", ""}
	result := NewLine(tab)
	if !result {
		t.Error("Test FAILED.")
	} else {
		t.Log("Test PASSED.")
	}

}
func TestPrintable(t *testing.T) {

	tab2 := []rune{'d'}
	result2 := Printable(tab2)
	if !result2 {
		t.Error("Test FAILED.")
	} else {
		t.Log("Test PASSED.")
	}
}

func TestBanner(t *testing.T) {
	fichier := "standard"
	expected := "./standard.txt"
	result := Banner(fichier)
	if expected != result {
		t.Error("Test FAILED.")
	} else {
		t.Log("Test PASSED.")
	}
}

func TestMatch(t *testing.T) {
	ascii := make(map[byte][]string)
	var index byte = 32
	file, err := os.ReadFile("./standard.txt")
	if err != nil {
		log.Fatal("Error : Not a ascci file in the repertory")
	}
	Split := strings.Split(string(file), "\n")
	for i := 1; i+8 < len(Split); i += 9 {
		ascii[index] = Split[i : i+8]
		index++
	}
	expected := " _       "
	result := Match('L', 0, ascii)
	if expected != result {
		t.Error("Test FAILED.")
	} else {
		t.Log("Test PASSED.")
	}
}

func TestFlag(t *testing.T) {
	flag1 := "--color=weex"
	result1 := Flag(flag1)
	if !result1 {
		t.Error("Test FAILED.")
	} else {
		t.Log("Test PASSED.")
	}
}

func TestColorFlag(t *testing.T) {
	flag := "--color=chartreuse"
	expected := "chartreuse"
	result := ColorFlag(flag)
	if result == expected {
		t.Log("Test PASSED.")
	} else {
		t.Error("Test FAILED.")
	}
}

func TestToColor(t *testing.T) {
	s := "Ab"
	r := 'b'
	if ToColor(s, r) {
		t.Log("Test PASSED.")
	} else {
		t.Error("Test FAILED.")
	}
}

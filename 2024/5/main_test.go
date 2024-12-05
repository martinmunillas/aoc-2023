package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	result := A("example.txt")
	if result != 143 {
		t.Fatalf("Invalid a example result: %d", result)
	}
	fmt.Println(A("input.txt"))
}

func TestB(t *testing.T) {
	result := B("example.txt")
	if result != 123 {
		t.Fatalf("Invalid b example result: %d", result)
	}
	fmt.Println(B("input.txt"))
}

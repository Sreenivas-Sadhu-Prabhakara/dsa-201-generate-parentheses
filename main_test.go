package main

import (
	"reflect"
	"testing"
)

func TestGenerateParens(t *testing.T) {
	got := GenerateParens(3)
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GenerateParens(3) = %v, want %v", got, want)
	}

	if got := GenerateParens(1); !reflect.DeepEqual(got, []string{"()"}) {
		t.Errorf("n=1: %v", got)
	}
	if got := GenerateParens(0); !reflect.DeepEqual(got, []string{""}) {
		t.Errorf("n=0: %v", got)
	}
	// Catalan numbers: C(4) = 14.
	if n := len(GenerateParens(4)); n != 14 {
		t.Errorf("n=4 count = %d, want 14", n)
	}
}

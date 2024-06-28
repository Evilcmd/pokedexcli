package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"Hello World", []string{"hello", "world"}},
		{"   GoLang   Test   ", []string{"golang", "test"}},
		{"UPPER case", []string{"upper", "case"}},
		{"MixedCASE Words", []string{"mixedcase", "words"}},
		{"", []string{}},
	}

	for _, test := range tests {
		result := cleanInput(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

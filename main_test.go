package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestGetLinesChannel(t *testing.T) {
	cases := []struct {
		name     string
		content  string
		expected []string
	}{
		{"empty", "", nil},
		{"lines with newline", "a\nb\n", []string{"a", "b"}},
		{"last line without newline", "first longer than eight bytes\nlast", []string{"first longer than eight bytes", "last"}},
		{"empty middle line", "a\n\nb", []string{"a", "", "b"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []string
			for line := range getLinesChannel(io.NopCloser(strings.NewReader(tc.content))) {
				got = append(got, line)
			}
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("lines = %#v, want %#v", got, tc.expected)
			}
		})
	}
}

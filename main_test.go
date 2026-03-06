package main

import "testing"

func TestFixPrefix(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"strips prefix from all lines", "  line one\n  line two\n  line three", "line one\nline two\nline three"},
		{"mixed prefixed and unprefixed", "  prefixed\nnot prefixed\n  prefixed again", "prefixed\nnot prefixed\nprefixed again"},
		{"single space not stripped", " one space\n  two spaces", " one space\ntwo spaces"},
		{"empty string", "", ""},
		{"no prefix", "no prefix here\nnor here", "no prefix here\nnor here"},
		{"blank lines preserved", "  first\n\n  third", "first\n\nthird"},
		{"tabs after prefix preserved", "  \ttabbed", "\ttabbed"},
		{"only spaces line becomes empty", "  ", ""},
		{"CRLF line endings", "  one\r\n  two\r\n  three", "one\r\ntwo\r\nthree"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fixPrefix(tt.input)
			if got != tt.want {
				t.Errorf("fixPrefix(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

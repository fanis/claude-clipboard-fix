package main

import "testing"

func TestFixPrefix(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"wrapped line rejoined", "  a long line that\n  continues here", "a long line that continues here"},
		{"blank line prevents rejoin", "  first\n\n  third", "first\n\nthird"},
		{"mixed prefixed and unprefixed", "  prefixed\nnot prefixed\n  prefixed again", "prefixed\nnot prefixed\nprefixed again"},
		{"single space not stripped", " one space\n  two spaces", " one space\ntwo spaces"},
		{"empty string", "", ""},
		{"no prefix", "no prefix here\nnor here", "no prefix here\nnor here"},
		{"tabs after prefix preserved", "  \ttabbed", "\ttabbed"},
		{"only spaces line becomes empty", "  ", ""},
		{"CRLF line endings", "  one\r\n  two\r\n  three", "one two three"},
		{"bullet list not rejoined", "  intro\n  - item one\n  - item two", "intro\n- item one\n- item two"},
		{"numbered list not rejoined", "  intro\n  1. first\n  2. second", "intro\n1. first\n2. second"},
		{"heading not rejoined", "  some text\n  # Heading", "some text\n# Heading"},
		{"indented next line not rejoined", "  line one\n  \tindented", "line one\n\tindented"},
		{"three wrapped lines rejoined", "  this is a very\n  long paragraph that\n  spans three lines", "this is a very long paragraph that spans three lines"},
		{"unprefixed line breaks rejoin", "  first part\nsecond part\n  third part", "first part\nsecond part\nthird part"},
		{"star bullet not rejoined", "  intro\n  * item", "intro\n* item"},
		{"paren numbered list not rejoined", "  intro\n  2) item", "intro\n2) item"},
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

package main

import "testing"

func TestFixPrefix(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"wrapped line rejoined", "  This is a long enough line that it would plausibly hit the wrap boundary in a\n  terminal window so it should be rejoined", "This is a long enough line that it would plausibly hit the wrap boundary in a terminal window so it should be rejoined"},
		{"short lines not rejoined", "  first line\n  second line\n  third line", "first line\nsecond line\nthird line"},
		{"blank line prevents rejoin", "  first\n\n  third", "first\n\nthird"},
		{"mixed prefixed and unprefixed", "  prefixed\nnot prefixed\n  prefixed again", "prefixed\nnot prefixed\nprefixed again"},
		{"single space not stripped", " one space\n  two spaces", " one space\ntwo spaces"},
		{"empty string", "", ""},
		{"no prefix", "no prefix here\nnor here", "no prefix here\nnor here"},
		{"tabs after prefix preserved", "  \ttabbed", "\ttabbed"},
		{"only spaces line becomes empty", "  ", ""},
		{"CRLF line endings", "  This line is long enough to plausibly be a soft wrap in the terminal output\r\n  and this continues it", "This line is long enough to plausibly be a soft wrap in the terminal output and this continues it"},
		{"CRLF short lines not rejoined", "  one\r\n  two\r\n  three", "one\r\ntwo\r\nthree"},
		{"bullet list not rejoined", "  This is a long introductory line that would wrap in a terminal window for sure\n  - item one\n  - item two", "This is a long introductory line that would wrap in a terminal window for sure\n- item one\n- item two"},
		{"numbered list not rejoined", "  This is a long introductory line that would wrap in a terminal window for sure\n  1. first\n  2. second", "This is a long introductory line that would wrap in a terminal window for sure\n1. first\n2. second"},
		{"heading not rejoined", "  This is a long introductory line that would wrap in a terminal window for sure\n  # Heading", "This is a long introductory line that would wrap in a terminal window for sure\n# Heading"},
		{"indented next line not rejoined", "  This is a long introductory line that would wrap in a terminal window for sure\n  \tindented", "This is a long introductory line that would wrap in a terminal window for sure\n\tindented"},
		{"three wrapped lines rejoined", "  This is a really long paragraph that would definitely wrap in the terminal so\n  the next line continues the thought naturally and then the third line also has\n  some more content that wraps", "This is a really long paragraph that would definitely wrap in the terminal so the next line continues the thought naturally and then the third line also has some more content that wraps"},
		{"unprefixed line breaks rejoin", "  first part\nsecond part\n  third part", "first part\nsecond part\nthird part"},
		{"star bullet not rejoined", "  This is a long introductory line that would wrap in a terminal window for sure\n  * item", "This is a long introductory line that would wrap in a terminal window for sure\n* item"},
		{"paren numbered list not rejoined", "  This is a long introductory line that would wrap in a terminal window for sure\n  2) item", "This is a long introductory line that would wrap in a terminal window for sure\n2) item"},
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

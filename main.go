// Copyright (c) 2026 Fanis Hatzidakis
// Licensed under PolyForm Internal Use License 1.0.0 - see LICENCE.md

package main

import (
	"strings"
	"syscall"
	"unsafe"
)

var (
	user32          = syscall.NewLazyDLL("user32.dll")
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	openClipboard   = user32.NewProc("OpenClipboard")
	closeClipboard  = user32.NewProc("CloseClipboard")
	getClipboardData = user32.NewProc("GetClipboardData")
	setClipboardData = user32.NewProc("SetClipboardData")
	emptyClipboard  = user32.NewProc("EmptyClipboard")
	globalAlloc     = kernel32.NewProc("GlobalAlloc")
	globalLock      = kernel32.NewProc("GlobalLock")
	globalUnlock    = kernel32.NewProc("GlobalUnlock")
)

const (
	cfUnicodeText = 13
	gmemMoveable  = 0x0002
)

func getClipboardText() (string, error) {
	r, _, err := openClipboard.Call(0)
	if r == 0 {
		return "", err
	}
	defer closeClipboard.Call()

	h, _, err := getClipboardData.Call(cfUnicodeText)
	if h == 0 {
		return "", err
	}

	ptr, _, err := globalLock.Call(h)
	if ptr == 0 {
		return "", err
	}
	defer globalUnlock.Call(h)

	// Find null terminator length
	n := 0
	for p := ptr; *(*uint16)(unsafe.Pointer(p)) != 0; p += 2 {
		n++
	}
	utf16 := unsafe.Slice((*uint16)(unsafe.Pointer(ptr)), n)
	return syscall.UTF16ToString(utf16), nil
}

func setClipboardText(s string) error {
	r, _, err := openClipboard.Call(0)
	if r == 0 {
		return err
	}
	defer closeClipboard.Call()

	utf16, err := syscall.UTF16FromString(s)
	if err != nil {
		return err
	}

	size := len(utf16) * 2
	h, _, err := globalAlloc.Call(gmemMoveable, uintptr(size))
	if h == 0 {
		return err
	}

	ptr, _, err := globalLock.Call(h)
	if ptr == 0 {
		return err
	}
	copy(unsafe.Slice((*uint16)(unsafe.Pointer(ptr)), len(utf16)), utf16)
	globalUnlock.Call(h)

	emptyClipboard.Call()
	setClipboardData.Call(cfUnicodeText, h)
	return nil
}

func fixPrefix(text string) string {
	lines := strings.Split(text, "\n")

	// Track which lines originally had the 2-space prefix
	hadPrefix := make([]bool, len(lines))
	for i, line := range lines {
		if strings.HasPrefix(line, "  ") {
			hadPrefix[i] = true
			lines[i] = line[2:]
		}
	}

	// Rejoin wrapped lines: if both the current and next line had the prefix,
	// and the break looks like a soft wrap (not an intentional newline), join them.
	var result []string
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for i+1 < len(lines) && isWrappedContinuation(lines, hadPrefix, i) {
			i++
			line = strings.TrimRight(line, "\r") + " " + lines[i]
		}
		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

// isWrappedContinuation returns true if lines[i+1] looks like a soft-wrapped
// continuation of lines[i] rather than an intentional new line.
func isWrappedContinuation(lines []string, hadPrefix []bool, i int) bool {
	cur := lines[i]
	next := lines[i+1]

	// Both lines must have originally had the 2-space prefix
	if !hadPrefix[i] || !hadPrefix[i+1] {
		return false
	}

	// Current line must have content (blank line = intentional break)
	if strings.TrimRight(cur, " \t\r") == "" {
		return false
	}

	// Next line must have content
	trimmed := strings.TrimRight(next, " \t\r")
	if trimmed == "" {
		return false
	}

	// Next line starting with whitespace suggests structure (code, indentation)
	if len(next) > 0 && (next[0] == ' ' || next[0] == '\t') {
		return false
	}

	// Next line starting with a bullet or list marker suggests intentional break
	if len(trimmed) > 0 && (trimmed[0] == '-' || trimmed[0] == '*' || trimmed[0] == '#') {
		return false
	}

	// Numbered list (e.g., "1. " or "1) ")
	if len(trimmed) > 1 && trimmed[0] >= '0' && trimmed[0] <= '9' {
		for j := 1; j < len(trimmed); j++ {
			if trimmed[j] >= '0' && trimmed[j] <= '9' {
				continue
			}
			if trimmed[j] == '.' || trimmed[j] == ')' {
				return false
			}
			break
		}
	}

	return true
}

func main() {
	text, err := getClipboardText()
	if err != nil || text == "" {
		return
	}
	fixed := fixPrefix(text)
	if fixed != text {
		setClipboardText(fixed)
	}
}

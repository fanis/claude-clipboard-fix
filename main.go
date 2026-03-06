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
	for i, line := range lines {
		if strings.HasPrefix(line, "  ") {
			lines[i] = line[2:]
		}
	}
	return strings.Join(lines, "\n")
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

# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
go build -o dist/ccfix.exe -ldflags="-s -w" main.go

# Test
go test -v

# Release (tag format: no 'v' prefix)
gh release create X.Y.Z dist/ccfix.exe --repo fanis/claude-clipboard-fix --title "vX.Y.Z" --notes "..."
```

## Architecture

Single-file Go program (`main.go`) with no dependencies. Accesses Windows clipboard via direct Win32 syscalls (user32.dll, kernel32.dll) using CF_UNICODETEXT (UTF-16). Reads clipboard, strips 2-space prefix from each line, writes back. Silent operation - no console output, no GUI.

Key functions:
- `fixPrefix` - pure string transform, the only testable unit
- `getClipboardText` / `setClipboardText` - Win32 clipboard via syscall, not unit-testable

## Conventions

- Windows-only (Win32 API)
- Zero external dependencies - stdlib and syscall only
- Binary output goes in `dist/` (gitignored)
- Licence: PolyForm Internal Use 1.0.0 - all source files carry a copyright header

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2026-03-10

### Fixed
- Soft-wrapped lines are now rejoined after prefix stripping (long lines no longer break mid-sentence)
- Intentional line breaks preserved: blank lines, bullets, numbered lists, headings, indented lines

## [1.0.0] - 2026-03-06

### Added
- Initial release
- Reads clipboard, strips 2-space prefix from each line, writes back
- Win32 clipboard access via direct syscalls (CF_UNICODETEXT, UTF-16)
- Silent operation - no console output, no GUI
- Zero external dependencies - stdlib and syscall only

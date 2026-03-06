# Build Procedure

## Building

```bash
go build -o dist/ccfix.exe -ldflags="-s -w" main.go
```

## Running Tests

```bash
go test -v
```

# Shell in Go

A small shell implementation written in Go.

## Requirements
- Go 1.22+ (or newer)

## Run
```bash
go run ./app
```

Or run the helper script:
```bash
./your_program.sh
```

## What this project currently supports
- Prompt loop (`$ `)
- Builtins: `exit`, `echo`, `type`
- Unknown command handling

## Project structure
- `app/main.go`: shell implementation
- `go.mod`: Go module definition
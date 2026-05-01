# Rendang 🍛 CLI Ordering 📝 with Cobra

A small project for learning Go that uses [Cobra](https://github.com/spf13/cobra) for a nice terminal experience for ordering Rendang. Also can't wait for Hari Raya Aidilfitri! 🥳

## Getting Started

Make sure Go is installed, then run:

```bash
go run . --help
```

Useful commands:

```bash
go run . status
go run . order
go run . order --first-name Faiq --last-name Adzlan --email faiq@example.com --packs 2 --no-input
go run . session
```

## Commands

### `status`

Shows the current event name and starting stock for the run.

### `order`

Places a single order. You can provide all values as flags, let the CLI prompt for them, or mix both approaches.

Example:

```bash
go run . order --first-name Faiq --last-name Adzlan --email faiq@example.com --packs 2 --no-input
```

### `session`

Starts an interactive multi-order flow that keeps asking for orders until you stop or the stock runs out.

## Global Flags

You can customize the event branding and stock amount on any command:

```bash
go run . status --event "Hari Raya Booth" --stock 250
```

## Example Output

![cli-multi-turn-sample](/assets/cli-multi-turn-sample.png)

## Development

Format and test the project with:

```bash
gofmt -w main.go cmd/*.go internal/rendang/*.go helper/*.go
go test ./...
```

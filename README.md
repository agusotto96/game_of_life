# Game of Life

A Conway's Game of Life implementation in Go using [Ebitengine](https://ebitengine.org/).

## Usage

```bash
go run .
```

### Options

See available flags and defaults:

```bash
go run . -h
```

## Examples

HighLife:
```bash
go run . -rule "B36/S23"
```

Diamoeba:
```bash
go run . -rule "B35678/S5678" -chance 0.45 -tps 30
```

Mazectric:
```bash
go run . -rule "B3/S1234" -chance 0.08 -tps 30
```

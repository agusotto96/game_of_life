# Game of Life

A Conway's Game of Life implementation in Go using [Ebitengine](https://ebitengine.org/).

[![Play Online](https://img.shields.io/badge/Play_Online-GitHub_Pages-blue?style=for-the-badge&logo=webassembly)](https://agusotto96.github.io/game_of_life/)

## Usage

```bash
go run .
```

### Options

See available flags and defaults:

```bash
go run . -h
```

**Web Version**: When playing online, you can configure the simulation by passing parameters in the URL just like CLI flags (e.g., `?rule=B36/S23&chance=0.2&tps=30`).

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

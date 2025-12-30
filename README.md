# draw

[![Go Reference](https://pkg.go.dev/badge/github.com/Doarakko/draw.svg)](https://pkg.go.dev/github.com/Doarakko/draw)
[![Go Report Card](https://goreportcard.com/badge/github.com/Doarakko/draw)](https://goreportcard.com/report/github.com/Doarakko/draw)

Draw Yu-Gi-Oh! Card.

## Install

go

```sh
go install github.com/Doarakko/draw@latest
```

Homebrew

```sh
brew install Doarakko/tap/draw
```

## Usage

```
draw
```

![](sample.gif)

## Credits

- [Yu-Gi-Oh! API by YGOPRODeck](https://db.ygoprodeck.com/api-guide/)
- [nyanko](https://github.com/mattn/nyanko)

## Release

```sh
# 1. Create and push tag
git tag v0.1.0
git push origin v0.1.0

# 2. Update homebrew-tap
./scripts/update-homebrew.sh 0.1.0
```

## License

MIT

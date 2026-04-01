# charm-banner

ASCII art banner generator for terminal UIs — styled with [Lip Gloss](https://github.com/charmbracelet/lipgloss).

> [Leia em Portugues](README.pt-br.md)

## Features

- Convert text to large ASCII art using Unicode block elements (`▀▄█░▒▓`)
- Three built-in font sizes: **Small** (3×3), **Medium** (5×5), **Large** (7×7)
- Full [Lip Gloss](https://github.com/charmbracelet/lipgloss) styling support
- Supports letters A-Z, digits 0-9, and common punctuation

## Install

```bash
go get github.com/junhinhow/charm-banner@latest
```

## Usage

```go
package main

import (
    "fmt"

    banner "github.com/junhinhow/charm-banner"
    "charm.land/lipgloss/v2"
)

func main() {
    // Simple banner
    fmt.Println(banner.Render("HELLO"))

    // With font size
    fmt.Println(banner.Render("GO", banner.WithFont(banner.Large)))

    // With Lip Gloss styling
    style := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FF79C6")).
        Bold(true)
    fmt.Println(banner.Render("CHARM", banner.WithStyle(style)))
}
```

## API

| Function | Description |
|----------|-------------|
| `Render(text string, opts ...Option) string` | Render text as ASCII art banner |
| `WithFont(Font) Option` | Set font size: `Small`, `Medium`, `Large` |
| `WithStyle(lipgloss.Style) Option` | Apply a Lip Gloss style to the output |

## License

[MIT](LICENSE)

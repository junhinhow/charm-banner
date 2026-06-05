# charm-banner

Gerador de banners em arte ASCII para interfaces de terminal — estilizado com [Lip Gloss](https://github.com/charmbracelet/lipgloss).

> [Read in English](README.md)

## Funcionalidades

- Converte texto em arte ASCII grande usando elementos de bloco Unicode (`▀▄█░▒▓`)
- Tres tamanhos de fonte integrados: **Small** (3x3), **Medium** (5x5), **Large** (7x7)
- Suporte completo a estilos [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- Suporta letras A-Z, digitos 0-9 e pontuacao comum

## Instalacao

```bash
go get github.com/junhinhow/charm-banner@latest
```

## Uso

```go
package main

import (
    "fmt"

    banner "github.com/junhinhow/charm-banner"
    "charm.land/lipgloss/v2"
)

func main() {
    // Banner simples
    fmt.Println(banner.Render("HELLO"))

    // Com tamanho de fonte
    fmt.Println(banner.Render("GO", banner.WithFont(banner.Large)))

    // Com estilo Lip Gloss
    style := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FF79C6")).
        Bold(true)
    fmt.Println(banner.Render("CHARM", banner.WithStyle(style)))
}
```

## API

| Funcao | Descricao |
|--------|-----------|
| `Render(text string, opts ...Option) string` | Renderiza texto como banner em arte ASCII |
| `WithFont(Font) Option` | Define o tamanho da fonte: `Small`, `Medium`, `Large` |
| `WithStyle(lipgloss.Style) Option` | Aplica um estilo Lip Gloss na saida |

## Licenca

[MIT](LICENSE)

# charm-banner

Gerador de banners em arte ASCII para interfaces de terminal — estilizado com [Lip Gloss](https://github.com/charmbracelet/lipgloss).

> [Read in English](README.md)

## Funcionalidades

- Converte texto em arte ASCII grande usando elementos de bloco Unicode (`▀▄█`)
- Três tamanhos de fonte integrados com **glifos quadrados**: **Small** (3×3), **Medium** (5×5), **Large** (7×7)
- Suporte completo a estilos [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- Suporta letras A-Z, dígitos 0-9 e pontuação comum (`! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~`)
- Letras acentuadas (`Á À Â Ã Ä É Ê Ë Í Î Ï Ó Ô Õ Ö Ú Û Ü Ç Ñ` …) reutilizam o glifo da letra base
- Caracteres que não fazem parte da fonte são desenhados como espaço

## Instalação

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

| Função | Descrição |
|--------|-----------|
| `Render(text string, opts ...Option) string` | Renderiza texto como banner em arte ASCII |
| `WithFont(Font) Option` | Define o tamanho da fonte: `Small`, `Medium`, `Large` |
| `WithStyle(lipgloss.Style) Option` | Aplica um estilo Lip Gloss na saída |

## Observações

- `Render("")` devolve uma string vazia.
- Cada caractere da entrada ocupa exatamente a largura da fonte (3, 5 ou 7 colunas) e todas as
  linhas de um banner têm sempre a mesma largura.
- O texto é normalizado antes da renderização: letras minúsculas viram maiúsculas e letras
  acentuadas caem na letra base (`"Á"` é desenhado exatamente como `"A"`).

## Licença

[MIT](LICENSE)

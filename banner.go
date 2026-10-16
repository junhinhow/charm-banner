// Package banner gera arte ASCII estilizada para interfaces de terminal.
//
// Utiliza elementos de bloco Unicode (▀▄█) para renderizar texto em tamanho grande,
// com suporte a estilos lipgloss e tamanhos de fonte configuráveis.
package banner

import (
	"strings"
	"unicode"

	"charm.land/lipgloss/v2"
)

// Font define o tamanho da fonte para renderização do banner.
type Font int

const (
	// Small renderiza caracteres com 3 colunas e 3 linhas.
	Small Font = iota
	// Medium renderiza caracteres com 5 colunas e 5 linhas.
	Medium
	// Large renderiza caracteres com 7 colunas e 7 linhas.
	Large
)

// Option configura o comportamento do Render.
type Option func(*config)

type config struct {
	style *lipgloss.Style
	font  Font
}

// WithStyle aplica um estilo lipgloss ao banner renderizado.
func WithStyle(s lipgloss.Style) Option {
	return func(c *config) {
		c.style = &s
	}
}

// WithFont define o tamanho da fonte do banner.
func WithFont(f Font) Option {
	return func(c *config) {
		c.font = f
	}
}

// Render converte texto em arte ASCII usando caracteres de bloco Unicode.
// Aceita opções para personalizar estilo e tamanho da fonte.
//
// O texto é normalizado antes da renderização: letras minúsculas são
// convertidas para maiúsculas e letras acentuadas (Á, Ç, Ê...) usam o glifo
// da letra base. Caracteres que não existem na fonte são desenhados como
// espaço. Uma entrada vazia devolve uma string vazia.
func Render(text string, opts ...Option) string {
	cfg := &config{
		font: Medium,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	runes := []rune(normalize(text))
	if len(runes) == 0 {
		return ""
	}

	font := getFont(cfg.font)
	height := fontHeight(cfg.font)
	width := fontWidth(cfg.font)

	lines := make([]string, height)

	// O índice do laço é a posição do rune (não o offset em bytes), para que o
	// separador entre caracteres seja acrescentado exatamente uma vez.
	for i, ch := range runes {
		pattern, ok := font[ch]
		if !ok {
			pattern = font[' ']
		}
		for row := 0; row < height; row++ {
			segment := ""
			if row < len(pattern.rows) {
				segment = pattern.rows[row]
			}
			lines[row] += padRight(segment, width)
		}
		// Espaço entre caracteres
		if i < len(runes)-1 {
			for row := 0; row < height; row++ {
				lines[row] += " "
			}
		}
	}

	result := strings.Join(lines, "\n")

	if cfg.style != nil {
		result = cfg.style.Render(result)
	}

	return result
}

// normalize converte o texto para maiúsculas e reduz letras acentuadas à
// letra base, para que "Á" use o mesmo glifo de "A".
func normalize(text string) string {
	upper := strings.ToUpper(text)
	return strings.Map(func(r rune) rune {
		if base, ok := baseRunes[r]; ok {
			return base
		}
		// Remove marcas de combinação (caso o texto venha em forma decomposta).
		if unicode.Is(unicode.Mn, r) {
			return -1
		}
		return r
	}, upper)
}

// baseRunes mapeia letras acentuadas e caracteres especiais para a letra base
// presente no mapa de glifos.
var baseRunes = map[rune]rune{
	'Á': 'A', 'À': 'A', 'Â': 'A', 'Ã': 'A', 'Ä': 'A', 'Å': 'A',
	'Ā': 'A', 'Ă': 'A', 'Ą': 'A', 'Ǎ': 'A',
	'Ç': 'C', 'Ć': 'C', 'Ĉ': 'C', 'Ċ': 'C', 'Č': 'C',
	'Ð': 'D', 'Ď': 'D', 'Đ': 'D',
	'É': 'E', 'È': 'E', 'Ê': 'E', 'Ë': 'E',
	'Ē': 'E', 'Ĕ': 'E', 'Ė': 'E', 'Ę': 'E', 'Ě': 'E',
	'Ĝ': 'G', 'Ğ': 'G', 'Ġ': 'G', 'Ģ': 'G',
	'Ĥ': 'H', 'Ħ': 'H',
	'Í': 'I', 'Ì': 'I', 'Î': 'I', 'Ï': 'I',
	'Ĩ': 'I', 'Ī': 'I', 'Ĭ': 'I', 'Į': 'I', 'İ': 'I',
	'Ĵ': 'J', 'Ķ': 'K',
	'Ĺ': 'L', 'Ļ': 'L', 'Ľ': 'L', 'Ŀ': 'L', 'Ł': 'L',
	'Ñ': 'N', 'Ń': 'N', 'Ņ': 'N', 'Ň': 'N',
	'Ó': 'O', 'Ò': 'O', 'Ô': 'O', 'Õ': 'O', 'Ö': 'O', 'Ø': 'O',
	'Ō': 'O', 'Ŏ': 'O', 'Ő': 'O',
	'Ŕ': 'R', 'Ŗ': 'R', 'Ř': 'R',
	'Ś': 'S', 'Ŝ': 'S', 'Ş': 'S', 'Š': 'S', 'Ș': 'S',
	'Ţ': 'T', 'Ť': 'T', 'Ŧ': 'T', 'Þ': 'T',
	'Ú': 'U', 'Ù': 'U', 'Û': 'U', 'Ü': 'U',
	'Ũ': 'U', 'Ū': 'U', 'Ŭ': 'U', 'Ů': 'U', 'Ű': 'U', 'Ų': 'U',
	'Ŵ': 'W',
	'Ý': 'Y', 'Ŷ': 'Y', 'Ÿ': 'Y',
	'Ź': 'Z', 'Ż': 'Z', 'Ž': 'Z',
	// Caracteres sem decomposição canônica que viram a letra mais próxima.
	'Æ': 'A', 'Œ': 'O',
}

// padRight preenche s com espaços até atingir a largura de exibição desejada.
func padRight(s string, width int) string {
	current := lipgloss.Width(s)
	if current >= width {
		return s
	}
	return s + strings.Repeat(" ", width-current)
}

// fontHeight retorna a altura em linhas para cada tamanho de fonte.
func fontHeight(f Font) int {
	switch f {
	case Small:
		return 3
	case Large:
		return 7
	default:
		return 5
	}
}

// fontWidth retorna a largura em colunas para cada tamanho de fonte.
func fontWidth(f Font) int {
	switch f {
	case Small:
		return 3
	case Large:
		return 7
	default:
		return 5
	}
}

// getFont retorna o mapa de caracteres para o tamanho de fonte especificado.
func getFont(f Font) map[rune]charPattern {
	switch f {
	case Small:
		return smallFont
	case Large:
		return largeFont
	default:
		return mediumFont
	}
}

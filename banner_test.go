package banner

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

func TestRenderRetornaSaida(t *testing.T) {
	result := Render("HI")
	if result == "" {
		t.Fatal("Render retornou string vazia")
	}
}

func TestRenderComFontes(t *testing.T) {
	fontes := []struct {
		nome   string
		font   Font
		altura int
	}{
		{"Small", Small, 3},
		{"Medium", Medium, 5},
		{"Large", Large, 7},
	}

	for _, f := range fontes {
		t.Run(f.nome, func(t *testing.T) {
			result := Render("A", WithFont(f.font))
			lines := strings.Split(result, "\n")
			if len(lines) != f.altura {
				t.Errorf("esperava %d linhas, obteve %d", f.altura, len(lines))
			}
		})
	}
}

func TestRenderComEstilo(t *testing.T) {
	style := lipgloss.NewStyle().Bold(true)
	result := Render("A", WithStyle(style))
	if result == "" {
		t.Fatal("Render com estilo retornou string vazia")
	}
}

func TestRenderCaractereDesconhecido(t *testing.T) {
	// Caracteres sem mapeamento devem usar espaço
	result := Render("@")
	if result == "" {
		t.Fatal("Render com caractere desconhecido retornou string vazia")
	}
}

func TestRenderEspaco(t *testing.T) {
	result := Render(" ")
	if result == "" {
		t.Fatal("Render com espaço retornou string vazia")
	}
}

func TestRenderMinusculasConvertidas(t *testing.T) {
	upper := Render("ABC")
	lower := Render("abc")
	if upper != lower {
		t.Error("minúsculas deveriam ser convertidas para maiúsculas")
	}
}

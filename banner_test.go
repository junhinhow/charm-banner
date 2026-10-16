package banner

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

// lineWidths devolve a largura de exibição de cada linha do banner.
func lineWidths(rendered string) []int {
	lines := strings.Split(rendered, "\n")
	widths := make([]int, len(lines))
	for i, line := range lines {
		widths[i] = lipgloss.Width(line)
	}
	return widths
}

// bannerWidth devolve a largura de exibição total do banner.
func bannerWidth(rendered string) int {
	return lipgloss.Width(rendered)
}

var allFonts = []struct {
	nome    string
	font    Font
	altura  int
	largura int
}{
	{"Small", Small, 3, 3},
	{"Medium", Medium, 5, 5},
	{"Large", Large, 7, 7},
}

func TestRenderVazioDevolveStringVazia(t *testing.T) {
	for _, f := range allFonts {
		t.Run(f.nome, func(t *testing.T) {
			resultado := Render("", WithFont(f.font))
			if resultado != "" {
				t.Errorf("Render(\"\") deveria devolver string vazia, obteve %q", resultado)
			}
		})
	}
}

func TestRenderAlturaDaFonte(t *testing.T) {
	for _, f := range allFonts {
		t.Run(f.nome, func(t *testing.T) {
			linhas := strings.Split(Render("A", WithFont(f.font)), "\n")
			if len(linhas) != f.altura {
				t.Errorf("esperava %d linhas, obteve %d", f.altura, len(linhas))
			}
		})
	}
}

// TestRenderLarguraDaFonte garante que um único caractere ocupe exatamente
// a largura documentada da fonte (3, 5 ou 7 colunas).
func TestRenderLarguraDaFonte(t *testing.T) {
	for _, f := range allFonts {
		t.Run(f.nome, func(t *testing.T) {
			largura := bannerWidth(Render("A", WithFont(f.font)))
			if largura != f.largura {
				t.Errorf("esperava largura %d, obteve %d", f.largura, largura)
			}
		})
	}
}

// TestRenderTodasLinhasMesmaLargura é o teste central de alinhamento:
// todas as linhas de um banner devem ter exatamente a mesma largura.
func TestRenderTodasLinhasMesmaLargura(t *testing.T) {
	entradas := []string{
		"A",
		"HELLO",
		"WIDE TEXT 123",
		"BÁ",
		"ÁÉÍÓÚ",
		"áàâãéêíóôõúç",
		"!@#$%^&*()",
		"Ç",
	}
	for _, f := range allFonts {
		for _, entrada := range entradas {
			t.Run(f.nome+"/"+entrada, func(t *testing.T) {
				larguras := lineWidths(Render(entrada, WithFont(f.font)))
				esperado := f.largura*len([]rune(normalize(entrada))) + len([]rune(normalize(entrada))) - 1
				for i, w := range larguras {
					if w != esperado {
						t.Errorf("linha %d: esperava largura %d, obteve %d", i, esperado, w)
					}
				}
			})
		}
	}
}

// TestRenderEntradaMultibyte verifica que o separador entre caracteres não é
// duplicado quando o último caractere ocupa mais de um byte.
func TestRenderEntradaMultibyte(t *testing.T) {
	// "BÁ" tem 2 caracteres: 2 glifos + 1 separador = 7 colunas em Small.
	const esperado = 7
	largura := bannerWidth(Render("BÁ", WithFont(Small)))
	if largura != esperado {
		t.Errorf("Render(\"BÁ\") esperava %d colunas, obteve %d", esperado, largura)
	}

	// O mesmo vale quando o caractere multibyte vem no início.
	largura = bannerWidth(Render("ÁB", WithFont(Small)))
	if largura != esperado {
		t.Errorf("Render(\"ÁB\") esperava %d colunas, obteve %d", esperado, largura)
	}
}

// TestRenderSemSeparadorFinal garante que não há espaço sobrando no fim da
// última coluna, o que happens quando o índice de byte é confundido com o
// índice de rune.
func TestRenderSemSeparadorFinal(t *testing.T) {
	for _, entrada := range []string{"A", "Á", "BÁ", "AB", "AÁ"} {
		linhas := strings.Split(Render(entrada, WithFont(Small)), "\n")
		for i, linha := range linhas {
			if strings.HasSuffix(linha, " ") {
				t.Errorf("Render(%q) linha %d termina com espaço: %q", entrada, i, linha)
			}
		}
	}
}

// TestRenderCaractereNaoSuportado verifica que caracteres ausentes da fonte
// são desenhados como espaço, mantendo a largura e o alinhamento.
func TestRenderCaractereNaoSuportado(t *testing.T) {
	resultado := Render("漢", WithFont(Small))
	linhas := strings.Split(resultado, "\n")
	if len(linhas) != 3 {
		t.Fatalf("esperava 3 linhas, obteve %d", len(linhas))
	}
	for i, linha := range linhas {
		if linha != "   " {
			t.Errorf("linha %d: esperava espaço para caractere não suportado, obteve %q", i, linha)
		}
	}
	if bannerWidth(resultado) != 3 {
		t.Errorf("esperava largura 3, obteve %d", bannerWidth(resultado))
	}
}

// TestFonteQuadrada garante que todos os glifos de todas as fontes tenham
// exatamente a largura e a altura declaradas.
func TestFonteQuadrada(t *testing.T) {
	fontes := map[string]map[rune]charPattern{
		"Small":  smallFont,
		"Medium": mediumFont,
		"Large":  largeFont,
	}
	alturas := map[string]int{"Small": 3, "Medium": 5, "Large": 7}
	larguras := map[string]int{"Small": 3, "Medium": 5, "Large": 7}

	for nome, fonte := range fontes {
		t.Run(nome, func(t *testing.T) {
			for caractere, padrao := range fonte {
				if len(padrao.rows) != alturas[nome] {
					t.Errorf("glifo %q: esperava %d linhas, obteve %d", caractere, alturas[nome], len(padrao.rows))
				}
				for i, linha := range padrao.rows {
					if w := lipgloss.Width(linha); w != larguras[nome] {
						t.Errorf("glifo %q linha %d: esperava largura %d, obteve %d", caractere, i, larguras[nome], w)
					}
				}
			}
		})
	}
}

// TestFonteContemAlfabetoEDigitos garante que o conjunto básico de glifos
// está presente em todas as fontes.
func TestFonteContemAlfabetoEDigitos(t *testing.T) {
	for _, f := range allFonts {
		t.Run(f.nome, func(t *testing.T) {
			fonte := getFont(f.font)
			for c := 'A'; c <= 'Z'; c++ {
				if _, ok := fonte[c]; !ok {
					t.Errorf("fonte %s não contém a letra %q", f.nome, string(c))
				}
			}
			for c := '0'; c <= '9'; c++ {
				if _, ok := fonte[c]; !ok {
					t.Errorf("fonte %s não contém o dígito %q", f.nome, string(c))
				}
			}
			if _, ok := fonte[' ']; !ok {
				t.Errorf("fonte %s não contém o espaço", f.nome)
			}
		})
	}
}

// TestFonteContemPontuacaoComum verifica a pontuação prometida pela
// documentação README.
func TestFonteContemPontuacaoComum(t *testing.T) {
	pontuacao := "!.,-:/+=()[]{}*#&@$%'\"<>_|~^`?"
	for _, f := range allFonts {
		t.Run(f.nome, func(t *testing.T) {
			fonte := getFont(f.font)
			for _, c := range pontuacao {
				if _, ok := fonte[c]; !ok {
					t.Errorf("fonte %s não contém o caractere %q", f.nome, string(c))
				}
			}
		})
	}
}

// TestPontuacaoNaoViraEspaco garante que a pontuação comum é realmente
// desenhada (e não cai no fallback de espaço).
func TestPontuacaoNaoViraEspaco(t *testing.T) {
	espaco := Render(" ", WithFont(Small))
	for _, c := range "!.,-:/+=()[]{}*#&@$%'\"<>_|~^`?" {
		resultado := Render(string(c), WithFont(Small))
		if resultado == espaco {
			t.Errorf("o caractere %q foi desenhado como espaço", string(c))
		}
	}
}

// TestLetrasAcentuadasUsamGlifoBase verifica que letras acentuadas são
// desenhadas com o glifo da letra base.
func TestLetrasAcentuadasUsamGlifoBase(t *testing.T) {
	pares := [][2]string{
		{"Á", "A"}, {"À", "A"}, {"Â", "A"}, {"Ã", "A"}, {"Ä", "A"},
		{"É", "E"}, {"Ê", "E"}, {"Í", "I"}, {"Ó", "O"}, {"Ô", "O"},
		{"Õ", "O"}, {"Ú", "U"}, {"Ç", "C"}, {"Ñ", "N"},
	}
	for _, par := range pares {
		t.Run(par[0], func(t *testing.T) {
			if Render(par[0], WithFont(Small)) != Render(par[1], WithFont(Small)) {
				t.Errorf("Render(%q) deveria ser igual a Render(%q)", par[0], par[1])
			}
		})
	}
}

// TestLetrasAcentuadasNaoViramEspaco garante que as letras acentuadas não caem
// no fallback de caractere não suportado.
func TestLetrasAcentuadasNaoViramEspaco(t *testing.T) {
	espaco := Render(" ", WithFont(Small))
	for _, c := range "ÁÀÂÃÄÉÊËÍÎÏÓÔÕÖÚÛÜÇÑ" {
		if Render(string(c), WithFont(Small)) == espaco {
			t.Errorf("a letra acentuada %q foi desenhada como espaço", string(c))
		}
	}
}

func TestRenderComEstilo(t *testing.T) {
	style := lipgloss.NewStyle().Bold(true)
	resultado := Render("A", WithStyle(style))
	if resultado == "" {
		t.Fatal("Render com estilo retornou string vazia")
	}
	if !strings.Contains(resultado, "\n") {
		t.Error("Render com estilo deveria preservar as quebras de linha")
	}
}

// TestRenderComEstiloPreservaLargura garante que o estilo lipgloss (códigos
// ANSI) não altera a largura de exibição do banner.
func TestRenderComEstiloPreservaLargura(t *testing.T) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF79C6")).Bold(true)
	for _, f := range allFonts {
		t.Run(f.nome, func(t *testing.T) {
			semEstilo := Render("GO", WithFont(f.font))
			comEstilo := Render("GO", WithFont(f.font), WithStyle(style))
			if bannerWidth(semEstilo) != bannerWidth(comEstilo) {
				t.Errorf("largura com estilo (%d) difere da largura sem estilo (%d)",
					bannerWidth(comEstilo), bannerWidth(semEstilo))
			}
		})
	}
}

func TestRenderMinusculasConvertidas(t *testing.T) {
	if Render("ABC", WithFont(Small)) != Render("abc", WithFont(Small)) {
		t.Error("minúsculas deveriam ser convertidas para maiúsculas")
	}
}

// TestRenderSeparadorUnico verifica que existe exatamente uma coluna de
// espaço entre dois glifos.
func TestRenderSeparadorUnico(t *testing.T) {
	um := Render("A", WithFont(Small))
	dois := Render("AA", WithFont(Small))
	linhasUm := strings.Split(um, "\n")
	linhasDois := strings.Split(dois, "\n")
	for i := range linhasUm {
		if !strings.HasPrefix(linhasDois[i], linhasUm[i]+" ") {
			t.Errorf("linha %d: esperava %q como prefixo de %q", i, linhasUm[i]+" ", linhasDois[i])
		}
	}
}

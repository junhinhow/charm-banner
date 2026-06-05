package banner_test

import (
	"fmt"

	banner "github.com/junhinhow/charm-banner"

	"charm.land/lipgloss/v2"
)

func Example() {
	result := banner.Render("GO", banner.WithFont(banner.Small))
	fmt.Println(result)
}

func ExampleWithStyle() {
	style := lipgloss.NewStyle().Bold(true)
	result := banner.Render("HI", banner.WithFont(banner.Small), banner.WithStyle(style))
	fmt.Println(result)
}

package main

import (
	tui "driptionary/internal/tui/teaModel"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	m := tui.NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

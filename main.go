package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"multiviews/views"
)

func main() {
	initialModel := views.InitialModel()

	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Erro ao iniciar o programa:", err)
		os.Exit(1)
	}
}

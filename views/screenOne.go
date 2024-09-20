package views

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type screeOneModel struct {
	textInput textinput.Model
	err       error
}

func screenOne() screeOneModel {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return screeOneModel{
		textInput: ti,
		err:       nil,
	}
}

func (m screeOneModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m screeOneModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyCtrlN:
			menu := menuModel{}
			return RootScreen().SwitchScreen(&menu)
		}
		

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m screeOneModel) View() string {
	return fmt.Sprintf(
		"What’s your favorite Digimon?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
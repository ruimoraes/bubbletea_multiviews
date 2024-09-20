package views

import (
	tea "github.com/charmbracelet/bubbletea"
   
)

type menuModel struct {
    // ...your options
}

func (m menuModel) Init() tea.Cmd {
    return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC:
            return m, tea.Quit
		case tea.KeyCtrlA:            
            return screenTwo().Update(msg)
			case tea.KeyCtrlB:            
            return screenOne().Update(msg)			
        }
    default:
        return m, nil
    }

	return m, nil 
}

func (m menuModel) View() string {
    str := "This is the first screen. Press any key to switch to the second screen."
    return str
}
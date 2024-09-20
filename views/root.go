package views

import tea "github.com/charmbracelet/bubbletea"

type rootScreenModel struct {
    model  tea.Model
}

func RootScreen() rootScreenModel {
    var rootModel tea.Model
	
	menu := menuModel{}
	rootModel = &menu

    return rootScreenModel{
        model: rootModel,
    }
}

func (m rootScreenModel) Init() tea.Cmd {
    return m.model.Init()
}

func (m rootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m.model.Update(msg)
}

func (m rootScreenModel) View() string {
    return m.model.View()
}

func (m rootScreenModel) SwitchScreen(model tea.Model) (tea.Model, tea.Cmd) {
    m.model = model
    return m.model, m.model.Init()
}
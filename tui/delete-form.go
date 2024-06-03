package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/PhotoManager/internal"
	"github.com/PhotoManager/utils"
)

var idDelete int

type MDeleteForm struct {
	input textinput.Model
}

func NewMDeleteForm() MDeleteForm {
	input := textinput.New()
	input.Prompt = "Id: "
	input.Placeholder = "1"
	input.Focus()
	input.CharLimit = 50
	input.Width = 50

	return MDeleteForm{
		input: input,
	}
}

func (m *MDeleteForm) Init() tea.Cmd { return textinput.Blink }

func (m *MDeleteForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			utils.Throw(internal.Delete(idDelete))
			return RenderOptionListUpdate(func() tea.Msg { return RenderMsg{} })
		case "ctrl+c", "q", "esc":
			return RenderOptionListUpdate(func() tea.Msg { return RenderMsg{} })
		}
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.updateInputs()
	return m, cmd
}

func (m *MDeleteForm) View() string {
	s := "Delete Photo\n\n"
	s += m.input.View()
	s += "\n\nPress Enter to delete or Esc to cancel."
	return s
}

func (m *MDeleteForm) updateInputs() {
	idDelete, _ = strconv.Atoi(m.input.Value())
}

func RenderDeleteFormView() string {
	m := NewMDeleteForm()
	m.Init()
	return m.View()
}

func RenderDeleteFormUpdate() (tea.Model, tea.Cmd) {
	m := NewMDeleteForm()
	m.Init()
	return m.Update(func() tea.Msg { return RenderMsg{} })
}

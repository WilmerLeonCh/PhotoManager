package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var formCreatePhoto = &struct {
	title string
	url   string
}{}

const (
	title = iota
	url
)

type MCreateForm struct {
	inputs      []textinput.Model
	focusCursor int
}

func NewMCreateForm() MCreateForm {
	var inputs = make([]textinput.Model, 2)
	inputs[title] = textinput.New()
	inputs[title].Prompt = "Title: "
	inputs[title].Placeholder = "lorem impsun ..."
	inputs[title].Focus()
	inputs[title].CharLimit = 50
	inputs[title].Width = 50

	inputs[url] = textinput.New()
	inputs[url].Prompt = "Url: "
	inputs[url].Placeholder = "https://..."
	inputs[url].CharLimit = 1000
	inputs[url].Width = 100

	return MCreateForm{
		inputs:      inputs,
		focusCursor: title,
	}
}

func (m *MCreateForm) Init() tea.Cmd {
	return textinput.Blink
}

func (m *MCreateForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyDown:
			m.increaseFocusCursor()
		case tea.KeyShiftTab, tea.KeyUp:
			m.decreaseFocusCursor()
		case tea.KeyEnter:
			if m.focusCursor == url && m.inputs[title].Value() != "" {
				//Saving
				return RenderOptionListUpdate(func() tea.Msg { return RenderMsg{} })
			}
			m.increaseFocusCursor()
		case tea.KeyCtrlC, tea.KeyEsc:
			formCreatePhoto = nil
			return RenderOptionListUpdate(func() tea.Msg { return RenderMsg{} })
		default:
			break
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focusCursor].Focus()
	}
	var cmds = make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	m.updateInputs()
	return m, tea.Batch(cmds...)
}

func (m *MCreateForm) View() string {
	s := "Create form \n\n"
	for i, input := range m.inputs {
		s += input.View()
		if i < len(m.inputs)-1 {
			s += "\n\n"
		}
	}
	s += "\n\nPress 'ctrl+c' or 'esc' to quit."
	return s
}

func (m *MCreateForm) increaseFocusCursor() {
	m.focusCursor = (m.focusCursor + 1) % len(m.inputs)
}

func (m *MCreateForm) decreaseFocusCursor() {
	m.focusCursor--
	if m.focusCursor < 0 {
		m.focusCursor = len(m.inputs) - 1
	}
}

func (m *MCreateForm) updateInputs() {
	if formCreatePhoto == nil {
		return
	}
	formCreatePhoto.title = m.inputs[title].Value()
	formCreatePhoto.url = m.inputs[url].Value()
}

func RenderCreateFormView() string {
	m := NewMCreateForm()
	m.Init()
	return m.View()
}

func RenderCreateFormUpdate() (tea.Model, tea.Cmd) {
	m := NewMCreateForm()
	m.Init()
	return m.Update(func() tea.Msg { return RenderMsg{} })
}

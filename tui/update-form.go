package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/PhotoManager/internal"
	"github.com/PhotoManager/tui/common"
	"github.com/PhotoManager/tui/constants"
	"github.com/PhotoManager/utils"
)

var formUpdatePhoto = &internal.MPhoto{}

const (
	idUpdate = iota
	titleUpdate
	urlUpdate
)

type MUpdateForm struct {
	inputs      []textinput.Model
	focusCursor int
}

func NewMUpdateForm() MUpdateForm {
	var inputs = make([]textinput.Model, 3)
	inputs[idUpdate] = textinput.New()
	inputs[idUpdate].Prompt = "  "
	inputs[idUpdate].Placeholder = "1, 2, 3, ..."
	inputs[idUpdate].Focus()
	inputs[idUpdate].CharLimit = 50
	inputs[idUpdate].Width = 50

	inputs[titleUpdate] = textinput.New()
	inputs[titleUpdate].Prompt = "  "
	inputs[titleUpdate].Placeholder = "lorem impsun ..."
	inputs[titleUpdate].CharLimit = 50
	inputs[titleUpdate].Width = 50

	inputs[urlUpdate] = textinput.New()
	inputs[urlUpdate].Prompt = "  "
	inputs[urlUpdate].Placeholder = "https://pexels.com/..."
	inputs[urlUpdate].CharLimit = 1000
	inputs[urlUpdate].Width = 100

	return MUpdateForm{
		inputs:      inputs,
		focusCursor: idUpdate,
	}
}

func (m *MUpdateForm) Init() tea.Cmd {
	return textinput.Blink
}

func (m *MUpdateForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.inputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDown:
			m.increaseFocusCursor()
		case tea.KeyUp:
			m.decreaseFocusCursor()
		case tea.KeyEnter:
			if m.focusCursor == len(m.inputs)-1 && m.inputs[titleCreate].Value() != "" {
				utils.Throw(internal.Update(*formUpdatePhoto))
				return RenderOptionListUpdate(func() tea.Msg { return constants.RenderMsg{} })
			}
			m.increaseFocusCursor()
		case tea.KeyCtrlC, tea.KeyEsc:
			formUpdatePhoto = nil
			return RenderOptionListUpdate(func() tea.Msg { return constants.RenderMsg{} })
		default:
			break
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focusCursor].Focus()
	}
	m.updateInputs()
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m *MUpdateForm) View() string {
	s := "\n" + common.TitleStyle.Render(" ❯ Update photo") + "\n\n"
	s += common.PromptStyle.Width(m.inputs[idUpdate].Width).Render(" Id: ") + "\n"
	s += m.inputs[idUpdate].View() + "\n"
	s += common.PromptStyle.Width(m.inputs[titleUpdate].Width).Render(" Title: ") + "\n"
	s += m.inputs[titleUpdate].View() + "\n"
	s += common.PromptStyle.Width(m.inputs[urlUpdate].Width).Render(" Url: ") + "\n"
	s += m.inputs[urlUpdate].View() + "\n\n\n"
	s += common.PlaceholderStyle.Render("[↑] up • [↓] down • [enter] submit • [esc / ctrl + c] quit")
	return s
}

func (m *MUpdateForm) increaseFocusCursor() {
	m.focusCursor = (m.focusCursor + 1) % len(m.inputs)
}

func (m *MUpdateForm) decreaseFocusCursor() {
	m.focusCursor--
	if m.focusCursor < 0 {
		m.focusCursor = len(m.inputs) - 1
	}
}

func (m *MUpdateForm) updateInputs() {
	if formUpdatePhoto == nil {
		return
	}
	if idStr, errStrConv := strconv.Atoi(m.inputs[idUpdate].Value()); errStrConv == nil {
		formUpdatePhoto.Id = idStr
	}
	formUpdatePhoto.Title = m.inputs[titleUpdate].Value()
	formUpdatePhoto.Url = m.inputs[urlUpdate].Value()
}

func RenderUpdateFormView() string {
	m := NewMUpdateForm()
	m.Init()
	return m.View()
}

func RenderUpdateFormUpdate() (tea.Model, tea.Cmd) {
	m := NewMUpdateForm()
	m.Init()
	return m.Update(func() tea.Msg { return constants.RenderMsg{} })
}

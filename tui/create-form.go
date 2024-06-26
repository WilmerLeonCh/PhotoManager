package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/PhotoManager/internal"
	"github.com/PhotoManager/notification"
	"github.com/PhotoManager/tui/common"
	"github.com/PhotoManager/tui/constants"
	"github.com/PhotoManager/utils"
)

var formCreatePhoto = &internal.MPhoto{}

const (
	titleCreate = iota
	urlCreate
)

type MCreateForm struct {
	inputs      []textinput.Model
	focusCursor int
}

func NewMCreateForm() MCreateForm {
	var inputs = make([]textinput.Model, 2)
	inputs[titleCreate] = textinput.New()
	inputs[titleCreate].Prompt = "  "
	inputs[titleCreate].Placeholder = "lorem impsun ..."
	inputs[titleCreate].Focus()
	inputs[titleCreate].CharLimit = 50
	inputs[titleCreate].Width = 50

	inputs[urlCreate] = textinput.New()
	inputs[urlCreate].Prompt = "  "
	inputs[urlCreate].Placeholder = "https://pexels.com/..."
	inputs[urlCreate].CharLimit = 1000
	inputs[urlCreate].Width = 100

	return MCreateForm{
		inputs:      inputs,
		focusCursor: titleCreate,
	}
}

func (m *MCreateForm) Init() tea.Cmd {
	return textinput.Blink
}

func (m *MCreateForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDown:
			m.increaseFocusCursor()
		case tea.KeyUp:
			m.decreaseFocusCursor()
		case tea.KeyEnter:
			if m.inputs[titleCreate].Value() == "" {
				m.focusCursor = titleCreate
				break
			}
			if m.inputs[urlCreate].Value() == "" {
				m.focusCursor = urlCreate
				break
			}
			resMPhoto, errCreate := internal.Create(*formCreatePhoto)
			notify := notification.SlackClient.NewSlackMessage(notification.MsgActionCreate)
			if errCreate != nil {
				notify.Attachments[0].Color = notification.StatusColorActionError
				notify.Attachments[0].Title = errCreate.Error()
			} else {
				notify.Attachments[0].Color = notification.StatusColorActionSuccess
				notify.Attachments[0].Title = fmt.Sprintf("success | creating photo: %d [%s]", resMPhoto.Id, resMPhoto.Url)
			}
			utils.Throw(notification.SlackClient.SendMsg(notify))
			return RenderOptionListUpdate(func() tea.Msg { return constants.RenderMsg{} })
		case tea.KeyCtrlC, tea.KeyEsc:
			formCreatePhoto = nil
			return RenderOptionListUpdate(func() tea.Msg { return constants.RenderMsg{} })
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
	s := "\n" + common.TitleStyle.Render(" ❯ Create photo") + "\n\n"
	s += common.PromptStyle.Width(m.inputs[titleCreate].Width).Render(" Title: ") + "\n"
	s += m.inputs[titleCreate].View() + "\n"
	s += common.PromptStyle.Width(m.inputs[urlCreate].Width).Render(" Url: ") + "\n"
	s += m.inputs[urlCreate].View() + "\n\n\n"
	s += common.PlaceholderStyle.Render("[↑] up • [↓] down • [enter] submit • [esc / ctrl + c] quit")
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
	formCreatePhoto.Title = m.inputs[titleCreate].Value()
	formCreatePhoto.Url = m.inputs[urlCreate].Value()
}

func RenderCreateFormView() string {
	m := NewMCreateForm()
	m.Init()
	return m.View()
}

func RenderCreateFormUpdate() (tea.Model, tea.Cmd) {
	m := NewMCreateForm()
	m.Init()
	return m.Update(func() tea.Msg { return constants.RenderMsg{} })
}

package tui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/PhotoManager/internal"
	"github.com/PhotoManager/notification"
	"github.com/PhotoManager/tui/common"
	"github.com/PhotoManager/tui/constants"
	"github.com/PhotoManager/utils"
)

var idDelete int

type MDeleteForm struct {
	input textinput.Model
}

func NewMDeleteForm() MDeleteForm {
	input := textinput.New()
	input.Placeholder = "1, 2, 3, ..."
	input.Focus()
	input.CharLimit = 5
	input.Width = 50

	return MDeleteForm{
		input: input,
	}
}

func (m *MDeleteForm) Init() tea.Cmd { return textinput.Blink }

func (m *MDeleteForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			errDelete := internal.Delete(idDelete)
			notify := notification.SlackClient.NewSlackMessage(notification.MsgActionDelete)
			if errDelete != nil {
				notify.Attachments[0].Color = notification.StatusColorActionError
				notify.Attachments[0].Title = errDelete.Error()
			} else {
				notify.Attachments[0].Color = notification.StatusColorActionSuccess
				notify.Attachments[0].Title = fmt.Sprintf("success | deleting photo: %d", idDelete)
			}
			utils.Throw(notification.SlackClient.SendMsg(notify))
			return RenderOptionListUpdate(func() tea.Msg { return constants.RenderMsg{} })
		case tea.KeyCtrlC, tea.KeyEsc:
			return RenderOptionListUpdate(func() tea.Msg { return constants.RenderMsg{} })
		}
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.updateInputs()
	return m, cmd
}

func (m *MDeleteForm) View() string {
	s := "\n" + common.TitleStyle.Render(" ❯ Delete photo") + "\n\n"
	s += common.PromptStyle.Width(m.input.Width).Render(" Id: ") + "\n"
	s += m.input.View() + "\n\n\n"
	s += common.PlaceholderStyle.Render("[enter] submit • [esc / ctrl + c] quit")
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
	return m.Update(func() tea.Msg { return constants.RenderMsg{} })
}

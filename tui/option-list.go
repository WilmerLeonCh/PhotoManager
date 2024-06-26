package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/PhotoManager/tui/common"
	"github.com/PhotoManager/utils"
)

type view int

const (
	createView view = iota
	listSearchView
	updateView
	deleteView
	optionListView
)

type tuiCursor int

var choiceActions = []string{
	tuiCursor(0): "Create",
	tuiCursor(1): "List & Search",
	tuiCursor(2): "Update",
	tuiCursor(3): "Delete",
}

type MOptionList struct {
	tuiCursor   tuiCursor
	tuiSelected view
	tuiCmd      func(tea.Msg) (tea.Model, tea.Cmd)
}

func CreateOptionList() {
	bubble := tea.NewProgram(NewMOptionList())
	utils.Must(bubble.Run())
}

func NewMOptionList() *MOptionList {
	return &MOptionList{
		tuiCursor:   tuiCursor(0),
		tuiSelected: optionListView,
		tuiCmd:      nil,
	}
}

func (m *MOptionList) Init() tea.Cmd { return nil }

func (m *MOptionList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDown:
			m.increaseViewCursor()
		case tea.KeyUp:
			m.decreaseViewCursor()
		case tea.KeyEnter:
			m.tuiSelected = view(m.tuiCursor)
			return m.handleCmd()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		default:
			break
		}
	}
	return m, nil
}

func (m *MOptionList) View() string {
	return m.handleView()
}

func (m *MOptionList) increaseViewCursor() {
	m.tuiCursor = tuiCursor((int(m.tuiCursor) + 1) % len(choiceActions))
}

func (m *MOptionList) decreaseViewCursor() {
	m.tuiCursor--
	if m.tuiCursor < 0 {
		m.tuiCursor = tuiCursor(len(choiceActions) - 1)
	}
}

func (m *MOptionList) handleView() string {
	switch m.tuiSelected {
	case optionListView:
		return m.RenderOptionListView()
	case createView:
		return RenderCreateFormView()
	case listSearchView:
		return RenderListTableView()
	case updateView:
		return RenderUpdateFormView()
	case deleteView:
		return RenderDeleteFormView()
	default:
		return ""
	}
}

func (m *MOptionList) handleCmd() (tea.Model, tea.Cmd) {
	switch m.tuiSelected {
	case optionListView:
		return m, nil
	case createView:
		return RenderCreateFormUpdate()
	case listSearchView:
		return RenderListTableUpdate()
	case updateView:
		return RenderUpdateFormUpdate()
	case deleteView:
		return RenderDeleteFormUpdate()
	default:
		return m, nil
	}
}

func (m *MOptionList) RenderOptionListView() string {
	s := "\n" + common.TitleStyle.Render(" What would you like to do?") + "\n\n"
	for i, choice := range choiceActions {
		cursor := " "
		if tuiCursor(i) == m.tuiCursor {
			cursor = "❯"
		}
		s += common.OptionListStyle.Render(fmt.Sprintf("%s %s", cursor, choice))
	}
	s += "\n\n" + common.PlaceholderStyle.Render("[↑] up • [↓] down • [enter] select • [esc / ctrl + c] quit")
	return s
}

func RenderOptionListUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	m := NewMOptionList()
	return m.Update(msg)
}

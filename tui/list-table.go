package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/PhotoManager/internal"
	"github.com/PhotoManager/utils"
)

type MListTable struct {
	table table.Model
}

func NewMListTable() MListTable {
	mMPhotos := utils.Must(internal.Read())
	tableRows := mapToTableRow(mMPhotos)
	mTable := table.New(
		table.WithColumns([]table.Column{
			{Title: "Id", Width: 4},
			{Title: "Title", Width: 60},
			{Title: "URL", Width: 50},
		}),
		table.WithRows(tableRows),
		table.WithFocused(true),
		table.WithHeight(10),
	)
	return MListTable{
		table: mTable,
	}
}

func (m MListTable) Init() tea.Cmd { return nil }

func (m MListTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return RenderOptionListUpdate(func() tea.Msg { return RenderMsg{} })
		}
	}
	m.table, _ = m.table.Update(msg)
	return m, nil
}

func (m MListTable) View() string {
	return m.table.View()
}

func mapToTableRow(mMPhotos []internal.MPhoto) (res []table.Row) {
	for _, mPhoto := range mMPhotos {
		res = append(res, table.Row{
			strconv.Itoa(mPhoto.Id),
			mPhoto.Title,
			mPhoto.Url,
		})

	}
	return
}

func RenderListTableView() string {
	m := NewMListTable()
	return m.View()
}

func RenderListTableUpdate() (tea.Model, tea.Cmd) {
	m := NewMListTable()
	return m.Update(func() tea.Msg { return RenderMsg{} })
}

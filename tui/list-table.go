package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/PhotoManager/internal"
	"github.com/PhotoManager/utils"
)

type MListTable struct {
	table     table.Model
	paginator paginator.Model
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

	mPaginator := paginator.New()
	mPaginator.Type = paginator.Dots
	mPaginator.PerPage = 10
	mPaginator.SetTotalPages(len(tableRows))

	return MListTable{
		table:     mTable,
		paginator: mPaginator,
	}
}

func (m MListTable) Init() tea.Cmd { return nil }

func (m MListTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.table.Cursor() == 0 && m.paginator.Page == 0 {
				break
			}
			if m.table.Cursor() == 0 {
				m.paginator.PrevPage()
				m.table.GotoBottom()
			}
			m.table, _ = m.table.Update(msg)
			return m, nil
		case tea.KeyDown:
			if m.table.Cursor() == m.paginator.PerPage-1 && m.paginator.Page == m.paginator.TotalPages-1 {
				break
			}
			if m.table.Cursor() == m.paginator.PerPage-1 {
				m.paginator.NextPage()
				m.table.GotoTop()
			}
			m.table, _ = m.table.Update(msg)
			return m, nil
		case tea.KeyLeft:
			m.paginator.PrevPage()
		case tea.KeyRight:
			m.paginator.NextPage()
		case tea.KeyCtrlC, tea.KeyEsc:
			return RenderOptionListUpdate(func() tea.Msg { return RenderMsg{} })
		default:
			break
		}
	}
	m.table, _ = m.table.Update(msg)
	return m, nil
}

func (m MListTable) View() string {
	s := "List & Search photos\n\n"
	s += m.table.View() + "\n\n"
	s += m.paginator.View() + "\n\n"
	s += "Press 'ctrl+c' or 'esc' to quit."
	return s
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

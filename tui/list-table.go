package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/PhotoManager/internal"
	"github.com/PhotoManager/utils"
)

const (
	defaultStart = 0
	defaultLimit = 10
)

var (
	titleSearch      string
	reloadReading    bool
	startCurrentPage int
)

type MListTable struct {
	inputTitleSearch textinput.Model
	table            table.Model
	paginator        paginator.Model
}

func NewMListTable() MListTable {
	mInput := textinput.New()
	mInput.Prompt = "Search by title: "
	mInput.Placeholder = "lorem impsun ..."
	mInput.Focus()
	mInput.CharLimit = 50
	mInput.Width = 50

	mTable := table.New(
		table.WithColumns([]table.Column{
			{Title: "Id", Width: 4},
			{Title: "Title", Width: 60},
			{Title: "URL", Width: 40},
		}),
		table.WithRows([]table.Row{}),
		table.WithFocused(true),
		table.WithHeight(defaultLimit),
	)

	mPaginator := paginator.New()
	mPaginator.Type = paginator.Dots
	mPaginator.PerPage = defaultLimit

	startCurrentPage = defaultStart
	reloadReading = true

	return MListTable{
		inputTitleSearch: mInput,
		table:            mTable,
		paginator:        mPaginator,
	}
}

func (m *MListTable) Init() tea.Cmd { return textinput.Blink }

func (m *MListTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case RenderMsg:
		return m, nil
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.table.Cursor() == defaultStart && m.paginator.Page == 0 {
				break
			}
			if m.table.Cursor() == defaultStart {
				m.paginator.PrevPage()
				m.table.GotoBottom()
				startCurrentPage = m.paginator.Page * m.paginator.PerPage
				reloadReading = true
				break
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
				startCurrentPage = m.paginator.Page * m.paginator.PerPage
				reloadReading = true
				break
			}
			m.table, _ = m.table.Update(msg)
			return m, nil
		case tea.KeyLeft:
			break
		case tea.KeyCtrlLeft:
			if m.paginator.Page > 0 {
				startCurrentPage = (m.paginator.Page - 1) * m.paginator.PerPage
			}
			m.paginator.PrevPage()
			reloadReading = true
		case tea.KeyRight:
			break
		case tea.KeyCtrlRight:
			if m.paginator.Page < m.paginator.TotalPages-1 {
				startCurrentPage = (m.paginator.Page + 1) * m.paginator.PerPage
			}
			m.paginator.NextPage()
			reloadReading = true
		case tea.KeyEnter:
			if titleSearch == m.inputTitleSearch.Value() {
				break
			}
			titleSearch = m.inputTitleSearch.Value()
			reloadReading = true
		case tea.KeyCtrlC, tea.KeyEsc:
			return RenderOptionListUpdate(func() tea.Msg { return RenderMsg{} })
		default:
			break
		}
	}
	m.inputTitleSearch, _ = m.inputTitleSearch.Update(msg)
	return m, tea.Batch(cmds...)
}

func (m *MListTable) View() string {
	if reloadReading {
		mPhotosFilter := internal.NewMPhotosFilter(
			internal.WithTitle(titleSearch),
			internal.WithStart(startCurrentPage),
			internal.WithLimit(m.paginator.PerPage),
		)
		mPhotosResponse := utils.Must(internal.Read(mPhotosFilter))
		tableRows := mapToTableRow(mPhotosResponse.Photos)
		m.table.SetRows(tableRows)
		m.paginator.SetTotalPages(mPhotosResponse.TotalCount)
		reloadReading = false
	}

	s := "List & Search photos\n\n"
	s += m.inputTitleSearch.View() + "\n\n"
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

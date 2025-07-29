package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/Vedjw/DonkeyType/state"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var titleArt string = `
▄     ▌     ▄▖      
▌▌▛▌▛▌▙▘█▌▌▌▐ ▌▌▛▌█▌
▙▘▙▌▌▌▛▖▙▖▙▌▐ ▙▌▙▌▙▖
          ▄▌  ▄▌▌                                     
`
var titleArtTrimmed = strings.TrimSpace(titleArt)

const listHeight = 15

var (
	mainTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("221")).
			MarginLeft(2).
			Align(lipgloss.Left)

	subTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("231")).
			Margin(1, 0, 0, 0)

	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("208"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type listModel struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			isQuitting = m.quitting
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
				state.SelectedLength = state.Length((m.list.Cursor()))
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m listModel) View() string {
	if m.choice != "" {
		return ""
	}
	if m.quitting {
		return quitTextStyle.Render("Come back again")
	}

	return "\n" + m.list.View()
}

var isQuitting bool

func RenderList() (bool, error) {
	items := []list.Item{
		item("short"),
		item("medium"),
		item("long"),
		item("thicc"),
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = lipgloss.JoinVertical(lipgloss.Left,
		mainTitleStyle.Render(titleArtTrimmed),
		subTitleStyle.Render("Choose Difficulty"),
	)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = mainTitleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := listModel{list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		return true, err
	}

	return isQuitting, nil
}

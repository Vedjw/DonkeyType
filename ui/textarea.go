package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Vedjw/DonkeyType/internals/words"
	"github.com/Vedjw/DonkeyType/state"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var art string = `
██████╗░░█████╗░███╗░░██╗██╗░░██╗███████╗██╗░░░██╗████████╗██╗░░░██╗██████╗░███████╗
██╔══██╗██╔══██╗████╗░██║██║░██╔╝██╔════╝╚██╗░██╔╝╚══██╔══╝╚██╗░██╔╝██╔══██╗██╔════╝
██║░░██║██║░░██║██╔██╗██║█████═╝░█████╗░░░╚████╔╝░░░░██║░░░░╚████╔╝░██████╔╝█████╗░░
██║░░██║██║░░██║██║╚████║██╔═██╗░██╔══╝░░░░╚██╔╝░░░░░██║░░░░░╚██╔╝░░██╔═══╝░██╔══╝░░
██████╔╝╚█████╔╝██║░╚███║██║░╚██╗███████╗░░░██║░░░░░░██║░░░░░░██║░░░██║░░░░░███████╗
╚═════╝░░╚════╝░╚═╝░░╚══╝╚═╝░░╚═╝╚══════╝░░░╚═╝░░░░░░╚═╝░░░░░░╚═╝░░░╚═╝░░░░░╚══════╝`

// ---------- Lipgloss Styles ----------

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("128")).
			Underline(true).
			Padding(1, 2, 0, 2) // top right bottom left

	wordsStyle = lipgloss.NewStyle().
			Italic(true).
			Bold(true).
			Foreground(lipgloss.Color("97")).
			Padding(1, 2, 1, 2)

	textareaStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("211")).
			Padding(0, 2)

	inputStyleCorrect = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("42"))
		//Padding(1, 2, 1, 2)

	inputStyleWrong = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("88"))
	//Padding(1, 2, 1, 2)
	cursorStatus = lipgloss.NewStyle().
			Bold(true)
)

// ---------- Model ----------

type target struct {
	target string
}
type output struct {
	output string
}

func (o *output) update(val string) {
	o.output = val
}

type model struct {
	width      int
	height     int
	ta         textarea.Model
	quitting   bool
	confirmed  bool
	words      *target
	useroutput *output
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = ""
	ta.Focus()
	ta.CharLimit = 2000
	ta.ShowLineNumbers = false
	ta.Prompt = ""
	ta.FocusedStyle.Base = textareaStyle
	ta.Cursor.Blink = false

	return model{
		ta: ta,
		words: &target{
			target: words.WordsSelector(state.SelectedLength),
		},
		useroutput: &output{output: ""},
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.ta.Focused() {
				m.confirmed = true
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ta.SetWidth(msg.Width - 6)
		m.ta.SetHeight(3)
	}

	m.useroutput.update(m.ta.Value())
	m.ta, cmd = m.ta.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return "Exiting...\n"
	}

	if m.confirmed {
		return fmt.Sprintf("You typed:\n%s", m.ta.Value())
	}

	header := headerStyle.Width(m.width).Render("START TYPING")
	words := wordsStyle.Width(m.width).Render(m.words.target)
	textareaView := textareaStyle.Render(m.ta.View())
	cursor := m.ta.LineInfo().CharOffset
	cursorStat := cursorStatus.Render(strconv.Itoa(cursor))

	return lipgloss.JoinVertical(lipgloss.Left, header, words, textareaView, cursorStat)
}

func checkInput(cursor int, t *target, o *output) {

}

func RenderTextarea() error {
	if _, err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Run(); err != nil {
		return err
	}
	return nil
}

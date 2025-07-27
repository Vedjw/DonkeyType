package ui

import (
	"fmt"

	"github.com/Vedjw/DonkeyType/internals/words"
	"github.com/Vedjw/DonkeyType/state"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ---------- Lipgloss Styles ----------

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("221")).
			Underline(true).
			Padding(1)

	wordsStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("137")).
			Padding(1)

	textareaStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("231")).
			Padding(1)

	inputStyleCorrect = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("42")).
				Padding(1)

	inputStyleWrong = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("160")).
			Padding(1)
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

type textareaModel struct {
	width      int
	height     int
	ta         textarea.Model
	quitting   bool
	confirmed  bool
	words      *target
	useroutput *output
}

func initialModel() textareaModel {
	ta := textarea.New()
	ta.Placeholder = ""
	ta.Focus()
	ta.CharLimit = 2000
	ta.ShowLineNumbers = false
	ta.Prompt = ""
	ta.FocusedStyle.Base = textareaStyle
	ta.Cursor.Blink = false
	m := textareaModel{
		ta: ta,
		words: &target{
			target: words.WordsSelector(state.SelectedLength),
		},
		useroutput: &output{output: ""},
	}

	state.TotalChLength = len(m.words.target)

	return m
}

func (m textareaModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m textareaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		m.ta.SetWidth(msg.Width - 4)
		m.ta.SetHeight(3)
	}

	m.useroutput.update(m.ta.Value())
	m.ta, cmd = m.ta.Update(msg)
	return m, cmd
}

func (m textareaModel) View() string {
	if m.quitting {
		return "Exiting...\n"
	}

	if m.confirmed {
		return fmt.Sprintf("You typed:\n%s", m.ta.Value())
	}

	header := headerStyle.Width(m.width).Render("START TYPING")
	words := wordsStyle.Width(m.width).Render(m.words.target)
	textareaView := m.ta.View()

	var userInputStatus string
	if len(m.useroutput.output) > 0 {
		if !isInputCorrect(m.words, m.useroutput) {
			userInputStatus = inputStyleWrong.Render("U Goofed up!")
		} else {
			userInputStatus = inputStyleCorrect.Render("Keep Going")
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, header, userInputStatus, words, textareaView)
}

func isInputCorrect(t *target, o *output) bool {
	currIndex := len(o.output) - 1

	if currIndex >= len(t.target) {
		return false
	}

	if t.target[currIndex] != o.output[currIndex] {
		state.TotalMistakes++
		return false
	}

	return true
}

func RenderTextarea() error {
	if _, err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Run(); err != nil {
		return err
	}
	return nil
}

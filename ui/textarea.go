package ui

import (
	"time"

	"github.com/Vedjw/DonkeyType/internals/words"
	"github.com/Vedjw/DonkeyType/state"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("221")).
			Padding(1)

	wordsStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("137")).
			Padding(1)

	textareaStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("231")).
			Padding(1)

	statusStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("231")).
			Padding(0, 1)

	inputStyleCorrect = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("42"))

	inputStyleWrong = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("160"))
)

type target struct {
	target string
}
type output struct {
	output string
}

func (o *output) update(val string) {
	o.output = val
	state.TotalUserInputLenght = len(o.output)
}

type textareaModel struct {
	width      int
	height     int
	ta         textarea.Model
	quitting   bool
	confirmed  bool
	words      *target
	useroutput *output
	usertime   time.Time
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
			isQuitting = m.quitting
			return m, tea.Quit
		case "enter":
			if m.ta.Focused() {
				end := time.Now()
				if m.usertime.IsZero() {
					state.TimeTaken = 0
				} else {
					state.TimeTaken = end.Sub(m.usertime)
				}
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

	// Update user's typed output
	m.useroutput.update(m.ta.Value())

	// Start timer on first keystroke
	if m.usertime.IsZero() && len(m.useroutput.output) > 0 {
		m.usertime = time.Now()
	}

	// Count mistake only once per wrong index
	if len(m.useroutput.output) > 0 {
		currIndex := len(m.useroutput.output) - 1
		if currIndex < len(m.words.target) &&
			m.words.target[currIndex] != m.useroutput.output[currIndex] &&
			!state.MistakeMap[currIndex] {

			state.TotalMistakes++
			state.MistakeMap[currIndex] = true
		}
	}

	// Update textarea
	m.ta, cmd = m.ta.Update(msg)
	return m, cmd
}

func (m textareaModel) View() string {
	if m.quitting {
		return quitTextStyle.Render("Come back again")
	}

	if m.confirmed {
		return ""
	}

	header := headerStyle.Width(m.width).Render("START TYPING")
	words := wordsStyle.Width(m.width).Render(m.words.target)
	textareaView := m.ta.View()

	status := statusStyle.Render("Status: ")
	var userInputStatus string
	userInputStatus = status
	if len(m.useroutput.output) > 0 {
		if !isInputCorrect(m.words, m.useroutput) {
			userInputStatus = status + inputStyleWrong.Render("U Goofed up!")
		} else {
			userInputStatus = status + inputStyleCorrect.Render("Keep Going")
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, header, userInputStatus, words, textareaView)
}

func isInputCorrect(t *target, o *output) bool {
	currIndex := len(o.output) - 1

	if currIndex >= len(t.target) {
		return false
	}

	return t.target[currIndex] == o.output[currIndex]
}

func RenderTextarea() (bool, error) {
	if _, err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Run(); err != nil {
		return false, err
	}
	return isQuitting, nil
}

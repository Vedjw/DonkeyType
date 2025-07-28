package ui

import (
	"fmt"

	"github.com/Vedjw/DonkeyType/state"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var resultStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("221")).
	Bold(true)

var resultContent = `
┌───────────────────────┐
│        Results        │
└───────────────────────┘

  Time Taken   : %.2f
  WPM          : %.2f
  Mistakes     : %d

  Press Enter to play again, or q to quit...
`

var playAgain bool

type Result struct {
	TotalMistakes int
	WPM           float32
	TimeTaken     float64
}

type resultModel struct {
	result Result
}

func (m resultModel) Init() tea.Cmd {
	return nil
}

func (m resultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			playAgain = false
			return m, tea.Quit
		case "enter":
			playAgain = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m resultModel) View() string {
	str := resultStyle.Render(fmt.Sprintf(resultContent, m.result.TimeTaken, m.result.WPM, m.result.TotalMistakes))
	return str
}

func calWPM() float32 {
	minutes := state.TimeTaken.Seconds() / 60.0
	return ((float32(state.TotalUserInputLenght) / 5) / float32(minutes))
}

func RenderResults() (bool, error) {
	initialModel := resultModel{
		result: Result{
			TotalMistakes: state.TotalMistakes,
			WPM:           calWPM(),
			TimeTaken:     state.TimeTaken.Seconds(),
		},
	}

	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return false, err
	}
	return playAgain, nil
}

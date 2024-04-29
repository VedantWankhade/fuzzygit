/*
	Ref: https://github.com/charmbracelet/bubbletea/blob/master/examples/textinput/main.go
*/

package input

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func Get(ch chan<- string) {
	p := tea.NewProgram(initialModel(ch))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
	ch        chan<- string
}

func initialModel(ch chan<- string) model {
	ti := textinput.New()
	ti.Placeholder = "feature/xyz"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
		ch:        ch,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.ch <- m.textInput.Value()
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Rename\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

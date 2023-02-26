package tui

import (
	"driptionary/internal/consts"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// MainModel is the root state of the app.
type MainModel struct {
	textInput textinput.Model

	isLoading  bool
	textResult string

	err error
}

// NewModel configures the initial model at runtime.
func NewModel() MainModel {
	ti := textinput.New()
	ti.Placeholder = consts.Placeholder
	ti.Focus()
	ti.Prompt = consts.Prompt

	return MainModel{
		textInput:  ti,
		textResult: "",
		err:        nil,
	}
}

// Init returns any number of tea.Cmds at runtime.
func (m MainModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles all tea.Msgs in the Bubble Tea event loop.
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	// Handle specific on keypress events.
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyEnter:
			m.isLoading = true
			val := m.textInput.Value()
			return m, HandleSearch(val)

		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We got a valid response from the API.
	case SearchCompleteMsg:
		m.textInput.Reset()
		m.textResult = string(msg)
		return m, nil

	// We got an error. Assign it in the Model, and we'll handle it in the View.
	case ErrMsg:
		m.err = msg
		return m, nil
	}

	// Update the text input.
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

// View renders a string representation of the MainModel.
func (m MainModel) View() string {
	return fmt.Sprintf("%s\n%s\n%s\n\n%s\n",
		Logo,
		m.textInput.View(),
		BuildTextInputView(m.textResult),
		Help,
	)
}

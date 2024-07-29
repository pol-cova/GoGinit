package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

// Define the model struct
type model struct {
	projectName string
	framework   string
	choices     []string
	cursor      int
}

// Initialize the model with choices
func initialModel() model {
	return model{
		choices: []string{"echo", "gin", "fiber", "martini", "chi", "mux", "gofr"},
	}
}

// Init is the initialization method
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles the user's key presses
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.framework = m.choices[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the UI
func (m model) View() string {
	s := "Choose a framework:\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPress q to quit.\n"
	return s
}

// GetUserInput runs the Bubble Tea program and returns the selected project name and framework
func GetUserInput() (string, string) {
	p := tea.NewProgram(initialModel())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
		return "", ""
	}

	// Type assertion for the final model
	m, ok := finalModel.(model)
	if !ok {
		fmt.Println("Could not assert final model")
		return "", ""
	}

	// Prompt for project name (can be enhanced with Bubble Tea if needed)
	var projectName string
	fmt.Print("Enter project name: ")
	fmt.Scanln(&projectName)

	return projectName, m.framework
}

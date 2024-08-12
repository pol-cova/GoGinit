// Update the TUI to pass the setupDB flag
package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define the model struct
type model struct {
	projectName string
	framework   string
	choices     []string
	dbChoices   []string
	cursor      int
	step        int    // To track which step we are in
	input       string // To store user input for project name
	setupDB     bool   // To store the user's choice for setting up the database
}

// Initialize the model with choices
func initialModel() model {
	return model{
		choices:   []string{"echo", "gin", "fiber", "martini", "chi", "mux", "gofr", "fuego", "default"},
		dbChoices: []string{"Yes", "No"},
		step:      0,
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
			if m.step == 0 {
				// Save the project name
				m.projectName = m.input
				if m.projectName == "" {
					fmt.Println("Project name cannot be empty.")
					return m, nil
				}
				// Move to the next step (framework selection)
				m.step = 1
				m.cursor = 0
			} else if m.step == 1 {
				// Save the selected framework
				m.framework = m.choices[m.cursor]
				// Move to the next step (DB selection)
				m.step = 2
				m.cursor = 0
			} else if m.step == 2 {
				m.setupDB = m.dbChoices[m.cursor] == "Yes"
				// Exit the TUI after completing the selection
				return m, tea.Quit
			}
		case "backspace":
			if m.step == 0 && len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			if m.step == 0 {
				m.input += msg.String()
			}
		}
	}
	return m, nil
}

// View renders the UI with enhanced styling
func (m model) View() string {
	var s string

	// Styling
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("11")).
		Render("🚀  Select a Framework")

	selectedChoiceStyle := lipgloss.NewStyle().
		Render("👉 ")

	// Define color styles for each choice
	choiceStyles := []lipgloss.Style{
		lipgloss.NewStyle().Foreground(lipgloss.Color("4")),  // echo
		lipgloss.NewStyle().Foreground(lipgloss.Color("2")),  // gin
		lipgloss.NewStyle().Foreground(lipgloss.Color("3")),  // fiber
		lipgloss.NewStyle().Foreground(lipgloss.Color("6")),  // martini
		lipgloss.NewStyle().Foreground(lipgloss.Color("5")),  // chi
		lipgloss.NewStyle().Foreground(lipgloss.Color("13")), // mux
		lipgloss.NewStyle().Foreground(lipgloss.Color("9")),  // gofr
		lipgloss.NewStyle().Foreground(lipgloss.Color("9")),  // fuego
		lipgloss.NewStyle().Foreground(lipgloss.Color("8")),  // default
	}

	if m.step == 0 {
		// Prompt for project name
		s = lipgloss.NewStyle().
			Foreground(lipgloss.Color("12")).
			Render("Enter project name: " + m.input)
	} else if m.step == 1 {
		// Framework selection
		s = headerStyle + "\n\n"
		for i, choice := range m.choices {
			if m.cursor == i {
				s += selectedChoiceStyle + choice + "\n"
			} else {
				// Apply color style for each choice
				s += choiceStyles[i].Render(choice) + "\n"
			}
		}
	} else if m.step == 2 {
		// DB setup selection
		dbHeaderStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("14")).
			Render("🗃️  Setup Database")

		selectedDBChoiceStyle := lipgloss.NewStyle().
			Render("👉 ")

		s = dbHeaderStyle + "\n\n"
		for i, choice := range m.dbChoices {
			if m.cursor == i {
				s += selectedDBChoiceStyle + choice + "\n"
			} else {
				s += lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render(choice) + "\n"
			}
		}
	}

	s += "\n" + lipgloss.NewStyle().
		Foreground(lipgloss.Color("8")).
		Render("Press q to quit.")

	// Center the entire content
	return lipgloss.NewStyle().Align(lipgloss.Center).Width(30).Render(s)
}

// GetUserInput runs the Bubble Tea program and returns the selected project name, framework, and setupDB flag
func GetUserInput() (string, string, bool) {
	p := tea.NewProgram(initialModel())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
		return "", "", false
	}

	// Type assertion for the final model
	m, ok := finalModel.(model)
	if !ok {
		fmt.Println("Could not assert final model")
		return "", "", false
	}

	return m.projectName, m.framework, m.setupDB
}

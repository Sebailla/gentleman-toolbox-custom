package cli

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gentleman-programming/gentle-ai/internal/system"
)

var (
	styleTitle = lipgloss.NewStyle().
			Bold(true).
			Background(lipgloss.Color("#00FFFF")).
			Foreground(lipgloss.Color("#000000")).
			Padding(0, 1)

	styleBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00FFFF")).
			Padding(1).
			Margin(1)

	styleMagenta = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF"))
	styleCyan    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF"))
)

type consoleModel struct {
	project string
	score   int
	status  string
}

func (m consoleModel) Init() tea.Cmd {
	return nil
}

func (m consoleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j":
			m.status = "EJECUTANDO JUDGE..."
		case "p":
			m.status = "GENERANDO PLAN..."
		}
	}
	return m, nil
}

func (m consoleModel) View() string {
	var s strings.Builder

	s.WriteString(styleTitle.Render(" GENTLEMAN AI CONSOLE "))
	s.WriteString("\n\n")

	content := fmt.Sprintf("📂 Proyecto: %s\n", styleCyan.Render(m.project))
	content += fmt.Sprintf("📊 Arquitectura (AHI): %d/100\n", m.score)
	content += fmt.Sprintf("\n⚡ Estado: %s", styleMagenta.Render(m.status))

	s.WriteString(styleBox.Render(content))
	s.WriteString("\n\n [J] Judge  [P] Plan  [S] Sync  [Q] Quit")

	return s.String()
}

// RunConsole inicia la interfaz interactiva de terminal.
func RunConsole(args []string, detection system.DetectionResult) error {
	p := tea.NewProgram(consoleModel{
		project: "gentleman-toolbox", // Debería extraerse dinámicamente
		score:   85,                  // Debería calcularse dinámicamente
		status:  "LISTO PARA LA GUERRA",
	})
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

package views

import (
	"driptionary/internal/consts"
	"github.com/charmbracelet/lipgloss"
)

var (
	Logo = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#114EE8",
		Dark:  "#E4F714",
	}).Bold(true).SetString(consts.AsciiLogo)

	Help = lipgloss.NewStyle().Faint(true).SetString(consts.Help)
)

func BuildTextInputView(res string) lipgloss.Style {
	textResult := lipgloss.NewStyle().SetString(res)
	return textResult
}

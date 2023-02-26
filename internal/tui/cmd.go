package tui

import (
	"driptionary/internal/pkg/urbndct"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleSearch(q string) tea.Cmd {
	return func() tea.Msg {
		res, err := urbndct.GetTopResult(q)
		if err != nil {
			return ErrMsg(err)
		}

		return SearchCompleteMsg(res)
	}
}

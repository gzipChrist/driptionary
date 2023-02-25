package tuiCmd

import (
	"driptionary/internal/pkg/urbndct"
	"driptionary/internal/tui/tuiMsg"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleSearch(q string) tea.Cmd {
	return func() tea.Msg {
		res, err := urbndct.GetTopResult(q)
		if err != nil {
			return tuiMsg.ErrMsg(err)
		}

		return tuiMsg.SearchCompleteMsg(res)
	}
}

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/desertfox/go-2048/game"
)

type model struct {
	game game.Game
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "up", "w":
			m.game.ProcessAction(game.Action("up"))
		case "down", "s":
			m.game.ProcessAction(game.Action("down"))
		case "left", "a":
			m.game.ProcessAction(game.Action("left"))
		case "right", "d":
			m.game.ProcessAction(game.Action("right"))

		case "ctrl+c", "q":
			return m, tea.Quit

		}
	}
	return m, nil
}

func (m model) View() string {
	if m.game.State != "playing" {
		return m.game.State
	}
	return m.game.BoardString()
}

func main() {
	g := game.Game{}
	g.Start()

	p := tea.NewProgram(model{
		game: g,
	})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

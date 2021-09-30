package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pigeonligh/ssh-board/pkg/auth"
	"github.com/pigeonligh/ssh-board/pkg/board"
	"github.com/pigeonligh/ssh-board/pkg/board/snake"
)

func main() {
	username, ok := auth.GetUser()
	if !ok {
		fmt.Println("Access denied.")
		return
	}

	boards := []board.Board{
		snake.New(),
	}

	menu := board.NewMenu(username, boards)
	prog := tea.NewProgram(menu, tea.WithAltScreen())
	if err := prog.Start(); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	if menu.IsOK() {
		menu.Play()
	}
}

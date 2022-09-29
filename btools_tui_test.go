package main

import (
	"fmt"
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestTui(t *testing.T) {

	var initModel = model{
		todos:    []string{"cleanning", "wash clothes", "write a blog"},
		selected: make(map[int]struct{}),
	}

	cmd := tea.NewProgram(initModel)
	if err := cmd.Start(); err != nil {
		fmt.Println("start failed:", err)
		os.Exit(1)
	}
}

// 它这个是根据面向对象的性质实现的终端UI库，我们需要实现一个接口
type model struct {
	todos    []string
	cursor   int
	selected map[int]struct{}
}

// 这个函数做初始化操作
func (m model) Init() tea.Cmd {
	return nil
}

// 更新UI界面
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "todo list:\n\n"

	for i, choice := range m.todos {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}

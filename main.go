package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"retrocli.svenvowe.de/retrolist"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table  table.Model
	width  int
	height int
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.table.SetWidth(m.width)

		for _, col := range m.table.Columns() {
			col.Width = m.width / len(m.table.Columns())
		}

	case tea.KeyMsg:
		switch msg.String() {

		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}

		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return ""
	}

	table := baseStyle.Render(m.table.View())
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, table)
}

func main() {
	// test: load retrolist and display as a table
	// https://github.com/charmbracelet/bubbletea/blob/main/examples/table/main.go
	fmt.Printf("Welcome to RetroList\n")

	columns := []table.Column{
		{Title: "TODO", Width: 10},
		{Title: "Quantity", Width: 2},
		{Title: "Status", Width: 2},
	}

	rows := make([]table.Row, 0)

	// test: create some test data for now
	list := retrolist.NewRetroList("Test List", "Testing")

	for i := range 10 {
		item := retrolist.NewItem(fmt.Sprintf("Item %d", i+1), uint(rand.Intn(i+1)+1))
		list.AddItem(item)
	}

	for _, item := range list.Items {
		var status string
		if item.Done {
			status = "DONE"
		} else {
			status = "TODO"
		}

		row := table.Row{item.Title, strconv.Itoa(int(item.Quantity)), status}
		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(5),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	m := model{t, 0, 0}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	// fmt.Printf("Saving RetroList '%s'\n", list.Title)
	// err := list.Save(config.DefaultFilename)
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// }
}

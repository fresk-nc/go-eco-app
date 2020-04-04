package main

import (
	"fmt"
)

type menuItemAction func() bool

type menuItem struct {
	title  string
	action menuItemAction
}

type menu struct {
	items []menuItem
}

func (m *menu) print() {
	for i, item := range m.items {
		fmt.Println(i+1, "-", item.title)
	}
}

func (m *menu) input() int {
	i := readIntFromStdin("Введите номер пункта меню: ", func(i int) bool {
		return i > 0 && i <= len(m.items)
	})

	return i - 1
}

func (m *menu) run() {
	showMenu := true

	for showMenu {
		printSeparator()
		m.print()
		printSeparator()
		i := m.input()
		clear()
		showMenu = m.items[i].action()
	}
}

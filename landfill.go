package main

import (
	"fmt"
	"strconv"
)

type landfill struct {
	address string
	active  bool
	date    string
}

func (l *landfill) toSlice() []string {
	return []string{l.address, strconv.FormatBool(l.active), l.date}
}

func (l *landfill) print() {
	status := "Активная"
	if !l.active {
		status = "Закрытая"
	}

	fmt.Println("Адресс:", l.address)
	fmt.Println("Статус:", status)
	fmt.Println("Дата обнаружения:", l.date)
}

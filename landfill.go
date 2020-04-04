package main

import "fmt"

type landfill struct {
	address string
	status  string
	date    string
}

func (l *landfill) toSlice() []string {
	return []string{l.address, l.status, l.date}
}

func (l *landfill) print() {
	fmt.Println("Адресс:", l.address)
	fmt.Println("Статус:", l.status)
	fmt.Println("Дата обнаружения:", l.date)
}

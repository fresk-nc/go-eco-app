package main

import "fmt"

type recyclingPoint struct {
	address   string
	wasteType string
}

func (r *recyclingPoint) toSlice() []string {
	return []string{r.address, r.wasteType}
}

func (r *recyclingPoint) print() {
	fmt.Println("Адресс:", r.address)
	fmt.Println("Тип отходов:", r.wasteType)
}

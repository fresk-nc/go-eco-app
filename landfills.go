package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type landfills []landfill

func newLandfillsFromFile(filename string) landfills {
	if !fileExists(filename) {
		return landfills{}
	}

	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	r := csv.NewReader(strings.NewReader(string(bs)))
	r.Comma = ';'

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	landfills := landfills{}

	for _, record := range records {
		landfill := landfill{
			address: record[0],
			status:  record[1],
			date:    record[2],
		}

		landfills = append(landfills, landfill)
	}

	return landfills
}

func addLandfill(l landfills, landfill landfill) landfills {
	return append(l, landfill)
}

func removeLandfill(l landfills, i int) landfills {
	return append(l[:i], l[i+1:]...)
}

func (l landfills) print() {
	for i, landfill := range l {
		fmt.Printf("#%v\n", i+1)
		landfill.print()
		if i != len(l)-1 {
			printSeparator()
		}
	}
}

func (l landfills) saveToFile(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	w := csv.NewWriter(file)
	w.Comma = ';'

	records := [][]string{}

	for _, landfill := range l {
		records = append(records, landfill.toSlice())
	}

	w.WriteAll(records)

	if err := w.Error(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

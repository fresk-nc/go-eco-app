package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type recyclingPoints []recyclingPoint

func newRecyclingPointsFromFile(filename string) recyclingPoints {
	if !fileExists(filename) {
		return recyclingPoints{}
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

	recyclingPoints := recyclingPoints{}

	for _, record := range records {
		recyclingPoint := recyclingPoint{
			address:   record[0],
			wasteType: record[1],
		}

		recyclingPoints = append(recyclingPoints, recyclingPoint)
	}

	return recyclingPoints
}

func addRecyclingPoint(r recyclingPoints, recyclingPoint recyclingPoint) recyclingPoints {
	return append(r, recyclingPoint)
}

func removeRecyclingPoint(r recyclingPoints, i int) recyclingPoints {
	return append(r[:i], r[i+1:]...)
}

func (r recyclingPoints) print() {
	for i, recyclingPoint := range r {
		fmt.Printf("#%v\n", i+1)
		recyclingPoint.print()
		if i != len(r)-1 {
			printSeparator()
		}
	}
}

func (r recyclingPoints) saveToFile(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	w := csv.NewWriter(file)
	w.Comma = ';'

	records := [][]string{}

	for _, recyclingPoint := range r {
		records = append(records, recyclingPoint.toSlice())
	}

	w.WriteAll(records)

	if err := w.Error(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

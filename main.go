package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	args := os.Args
	if len(args) < 3 {
		panic("usage: <program> <path> <col>")
	}
	path := args[1]
	col := args[2]
	file, err := os.Open(path)
	if err != nil {
		s := fmt.Sprintf("Error opening %s: %s", path, err)
		panic(s)
	}
	defer file.Close()

	// create a new csv reader
	reader := csv.NewReader(file)

	// read all records from csv

}

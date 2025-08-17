package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("usage: <program> <path> <col>")
	}

	path := args[1]

	if !strings.HasSuffix(path, ".csv") {
		s := fmt.Sprintf("Path must have .csv file extension, not %s", path)
		panic(s)
	}

	// col := args[2]

	file, err := os.Open(path)
	if err != nil {
		s := fmt.Sprintf("Error opening %s: %s", path, err)
		panic(s)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var table map[string][]string
	var header []string
	count := 0

	// read the file line by line
	for scanner.Scan() {
		// expect row 0 to be the header
		if count == 0 {
			count++
			headerStr := scanner.Text()
			headerSlice := strings.Split(headerStr, ",")
			header = handleRow(headerSlice)
			table = make(map[string][]string, len(headerSlice))
			for _, col := range headerSlice {
				table[col] = []string{}
			}
			continue
		}

		count++
		rowStr := scanner.Text()
		rowSlice := strings.Split(rowStr, ",")
		rowSlice = handleRow(rowSlice)

		// panic if current row's length not equal to the header
		if len(header) != len(rowSlice) {
			s := fmt.Sprintf("Expected to have %v columns but got %v", len(header), len(rowSlice))
			panic(s)
		}

		// add to map with correct order
		for i, col := range rowSlice {
			rowHeader := header[i]
			table[rowHeader] = append(table[rowHeader], col)
		}
	}

	for k, v := range table {
		fmt.Println(k)
		fmt.Println(v)
	}

	// if error is not EOF, then panic the error
	if err := scanner.Err(); err != nil && err != io.EOF {
		s := fmt.Sprintf("Error reading %s: %s", path, err)
		panic(s)
	}
}

// handleRow makes sure we handle each row in csv file correct
func handleRow(cols []string) []string {
	newCols := []string{}
	var curr string
	for _, col := range cols {
		// start of special column
		if strings.HasPrefix(col, "\"") {
			curr = col[1:]
			continue
		}

		// in the middle of a special column
		if curr != "" {
			// it's the end
			if strings.HasSuffix(col, "\"") {
				curr += col[:len(col)-1]
				newCols = append(newCols, curr)
				curr = ""
				continue
			}

			// it's the middle
			curr += col
			continue
		}

		// normal column
		newCols = append(newCols, col)
	}
	return newCols
}

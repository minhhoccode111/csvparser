package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func printUsage() {
	fmt.Println("usage: <program> <csv-file-path> <target-col>")
}

func main() {
	args := os.Args
	if len(args) < 3 {
		printUsage()
		panic("")
	}

	path := args[1]

	if !strings.HasSuffix(path, ".csv") {
		printUsage()
		s := fmt.Sprintf("Path must have .csv file extension, not %s", path)
		panic(s)
	}

	targetCol := args[2]

	file, err := os.Open(path)
	if err != nil {
		printUsage()
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
			if _, ok := table[targetCol]; !ok {
				printUsage()
				s := fmt.Sprintf("target column %s not found in headers %v", targetCol, header)
				panic(s)
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

	// if error is not EOF, then panic the error
	if err := scanner.Err(); err != nil && err != io.EOF {
		s := fmt.Sprintf("Error reading %s: %s", path, err)
		panic(s)
	}

	// calculate statistics
	targetSliceStr := table[targetCol]
	targetSlice := []float64{}
	for _, v := range targetSliceStr {
		parsedValue, err := strconv.ParseFloat(v, 64)
		if err != nil {
			s := fmt.Sprintf("Error while parsing a value to float64 in target column: %s", v)
			panic(s)
		}
		targetSlice = append(targetSlice, parsedValue)
	}
	printResult(targetSlice)
}

func printResult(s []float64) {
	sort.Float64s(s)

	var sum float64 = 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("Length: %v\n", len(s))

	minVal := s[0]
	fmt.Printf("Min Value: %v\n", minVal)
	maxVal := s[len(s)-1]
	fmt.Printf("Max Value: %v\n", maxVal)
	meanVal := sum / float64(len(s))
	fmt.Printf("Mean Value: %v\n", meanVal)

	mid := len(s) / 2

	if len(s)%2 == 1 { // odd
		fmt.Printf("Median Value: %v\n", s[mid])
	} else if len(s)%2 == 0 && len(s) > 0 { // even, but not zero-length
		fmt.Printf("Median Value: %v, %v\n", s[mid-1], s[mid])
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

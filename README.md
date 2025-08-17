# CSV Parser Golang

Use Golang to parse CSV files and use it to calculate some basic statistics

- This project is part of [An Efficient Go Study Guide
  Article](https://learncodethehardway.com/blog/36-an-efficient-go-study-guide/)
- CSV files are generated using [this
  repo](https://github.com/datablist/sample-csv-files)

NOTE: We can use the `encoding/csv` package, which will make our life easier, but for the sake of practicing, we will use `bufio` instead.

Because sometimes, we will have to handle edge cases like this, a column has `,` character in it

```csv
21,8de40AC4e6EaCa4,"Velez, Payne and Coffey",http://burton.com/,Luxembourg,Mandatory coherent synergy,1986,Wholesale,5010
```

## Concepts

- Read file
- Deal with the CSV
- `strings.SplitSeq()` to return iterator

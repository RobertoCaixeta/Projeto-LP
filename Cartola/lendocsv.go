package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvLine struct {
	Column1 string
	Column2 string
	Column3 string
}

func main() {
  // Link CSV file
	filename := "https://github.com/henriquepgomide/caRtola/tree/master/data"

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into *lines* variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	// Loop through *lines*, create data object, each piece to their respective column
	for _, line := range lines {
		data := CsvLine{
			Column1: line[0],
			Column2: line[1],
			Column3: line[2],
		}
		fmt.Println(data.Column1 + " " + data.Column2 + " " + data.Column3)
	}
}
package products

// Importing the libraries
import (
	"encoding/csv"
	"fmt"
	"index/suffixarray"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// This function reads a file and using the for loop, it takes the data line by line
// and word by word to be returned as a string.
func AccessData(searchData string, number int) string {
	content, err := ioutil.ReadFile("products/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	suffix := suffixarray.New(content)
	indexList := suffix.Lookup([]byte(searchData), -1)

	if len(indexList) == 0 {
		fmt.Println("Data is not found")
		return ""
	}

	data := string(content)
	for _, ids := range indexList {
		start := ids
		for start >= 0 && data[start] != '\n' {
			start--
		}
		end := ids
		for end < len(data) && data[end] != '\n' {
			end++
		}
		Data := string(data[start+1 : end])
		sliceData := strings.Split(Data, ",")
		return sliceData[number]
	}
	return ""
}

// It reads the file and use the for loop to read all the data and append that data in columns
// in a result variable.
func AccessColumns(rowNumber int) []string {
	file, err := os.Open("products/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	var result []string

	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if len(row) > 0 {
			result = append(result, row[rowNumber])
		}
	}
	return result
}

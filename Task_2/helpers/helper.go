package helpers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"text/tabwriter"

	"github.com/axrav/Task2/structs"
)

// ReadCsv function to read data from csv file
func ReadCsv(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.LazyQuotes = true // to handle quotes in csv file
	reader.Comma = ';'       // to handle commas in csv file
	reader.FieldsPerRecord = -1

	return reader.ReadAll()
}

// ConvertToInt function to convert string to int
func ConvertToInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}

// FormStruct function to form a slice of structs from the data read from the csv file

func FormStruct(data [][]string) (DataArray []structs.Data) {
	for _, v := range data {

		for _, v2 := range v {
			tempArr := strings.Split(v2, ",")
			if tempArr[0] == "Index" {
				continue // to skip the first row of the csv file
			}
			DataArray = append(DataArray, structs.Data{
				Index:             ConvertToInt(tempArr[0]),
				OrganizationId:    tempArr[1],
				Name:              tempArr[2],
				Website:           tempArr[3],
				Country:           tempArr[4],
				Description:       tempArr[5],
				Founded:           tempArr[6],
				Industry:          tempArr[7],
				NumberOfEmployees: ConvertToInt(tempArr[8]),
			})

		}
	}
	return DataArray
}

// Using tabwriter to display data in tabular format as manual formatting might be an issue in case of large data

func DisplayTabularData(data []structs.Data) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Index\tOrganizationId\tName\tWebsite\tCountry\tDescription\tFounded\tIndustry\tNumberOfEmployees\t")
	for _, v := range data {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t", v.Index, v.OrganizationId, v.Name, v.Website, v.Country, v.Description, v.Founded, v.Industry, v.NumberOfEmployees)
		fmt.Fprintln(w)
	}

}

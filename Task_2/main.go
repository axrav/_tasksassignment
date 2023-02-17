package main

import (
	"github.com/axrav/Task2/helpers"
)

const FileName = "data.csv" // file name

// main function to read data from csv file and display it in a tabular format
func main() {
	data, err := helpers.ReadCsv(FileName)
	if err != nil {
		panic(err)
	}
	SliceStructs := helpers.FormStruct(data) // slice of structs
	helpers.DisplayTabularData(SliceStructs) // display data in tabular format
}

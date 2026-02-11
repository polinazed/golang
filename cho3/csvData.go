package main

import(
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct{
	Name string
	Surname string
	Number string
	LastAcces string
}
var myData = []Record{}

func readCSVFile(filepath string)([][]string, error){
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err 
	}
	return lines, nil
}
func saveCSVFile(filepath string) error{
	csvfile, err := os.Create(filepath)
	if err != nil{
		return err 
	}
	defer csvfile.Close()
	csvwriter := csv.New
}
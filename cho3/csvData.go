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
	
}
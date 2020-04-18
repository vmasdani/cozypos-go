package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readCsv(fileName string, parserFunc func(record []string)) {
	fmt.Printf("Reading %s...\n", fileName)

	// Read items
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		panic("Failed to open csv file.")
	}

	// Convert char array to string
	fileStr := ""
	for _, chr := range data {
		fileStr += string(chr)
	}

	// Read csv format
	r := csv.NewReader(strings.NewReader(fileStr))

	// Read csv line by line
	lineCounter := 0
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		lineCounter++

		if lineCounter == 1 {
			fmt.Println("Header")
			fmt.Println(record)
		} else {
			// fmt.Println(record)
			parserFunc(record)
		}
	}
}

func Populate() {
	readCsv("./csv/items.csv", func(record []string) {
		name := record[0]
		desc := record[1]
		price, priceErr := strconv.Atoi(record[2])
		manufacturingPrice, manufacturingPriceErr := strconv.Atoi(record[3])

		if priceErr == nil && manufacturingPriceErr == nil {
			fmt.Printf("Name: %s, Desc: %s, Price: %d, Manuf. Price: %d\n", name, desc, price, manufacturingPrice)
		}

		newItem := Item{
			Name:               name,
			Desc:               desc,
			Price:              price,
			ManufacturingPrice: manufacturingPrice}

		db.Save(&newItem)
	})

	// Items
	// items := []*Item{
	// 	&Item{
	// 		Name:               "Sticker",
	// 		Desc:               "Sticker goceng tiga",
	// 		Price:              5000,
	// 		ManufacturingPrice: 1000},
	// 	&Item{
	// 		Name:               "Doujinshi 1",
	// 		Desc:               "Doujinshi touhou",
	// 		Price:              30000,
	// 		ManufacturingPrice: 10000}}
	//
	// for _, item := range items {
	// 	db.Save(item)
	// }
}

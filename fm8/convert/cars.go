package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type Car struct {
	CarOrdinal int
	Year       int
	Make       string
	Model      string
}

func readCars() (map[int]Car, []int) {
	carsJsonFile, err := os.Open("../cars.json")
	if err != nil {
		fmt.Println(err)
	}
	defer carsJsonFile.Close()
	fmt.Println("cars.json opened")

	byteValue, err := io.ReadAll(carsJsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var cars []Car
	json.Unmarshal(byteValue, &cars)

	carsWithKeys := make(map[int]Car)
	keys := make([]int, 0, len(cars))
	for _, car := range cars {
		_, ok := carsWithKeys[car.CarOrdinal]
		if ok {
			panic(fmt.Sprintf("double car ordinal %v\n", car.CarOrdinal))
		}
		carsWithKeys[car.CarOrdinal] = car
		keys = append(keys, car.CarOrdinal)
	}
	sort.Ints(keys)

	return carsWithKeys, keys
}

func writeCarsKeys(carsWithKeys map[int]Car, _ []int) {
	fileJson, err := os.OpenFile("../cars_keys.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer fileJson.Close()

	jsonString, err := json.MarshalIndent(carsWithKeys, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	_, err = fileJson.Write(jsonString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cars_keys.json saved")
}

func writeCarsCsv(carsWithKeys map[int]Car, keys []int) {
	fileCsv, err := os.OpenFile("../cars.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer fileCsv.Close()
	for _, key := range keys {
		fmt.Fprintf(fileCsv, "%s", fmt.Sprintf("%d,%d,%s,%s\n", carsWithKeys[key].CarOrdinal, carsWithKeys[key].Year, carsWithKeys[key].Make, carsWithKeys[key].Model))
	}
	fmt.Println("cars.csv saved")
}

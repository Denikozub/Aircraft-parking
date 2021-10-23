package tables

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func openFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return file
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func WriteParkingPlacesToFile(parkingPlaces *[]int, realParkingPlaces *[]int, inputName string, outputName string) {
	inputFile := openFile(inputName)
	defer closeFile(inputFile)
	reader := csv.NewReader(inputFile)

	outputFile, err := os.Create(outputName)
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	//we don't need to change the first record
	record, err1 := reader.Read()
	if err1 != nil {
		panic(err1)
	}
	err2 := writer.Write(record)
	if err2 != nil {
		panic(err2)
	}

	for i := 0; i < len(*parkingPlaces); i++ {
		record, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		pPlace := (*parkingPlaces)[i]
		pPlace = (*realParkingPlaces)[pPlace]
		record[11] = strconv.FormatInt(int64(pPlace), 10)
		err = writer.Write(record)
	}
}

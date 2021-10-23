package tables

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"
)

type Plane struct {

	//A - arrival, D - departure
	AD byte

	//time of arrival/departure
	dateTime time.Time

	//I - international, D - domestic
	ID byte

	//terminal number
	terminal int

	//R - regional, N - Narrow Body, W - Wide Body
	planeClass byte

	//buses required
	busesRequired int
}

type Planes struct {
	amountPlanes int
	data []Plane
}

func (pl *Planes) GetNumberOfPlanes () int {
	return pl.amountPlanes
}

func (pl *Planes) LoadData (name string) {
	file := openFile(name)
	defer closeFile(file)
	reader := csv.NewReader(file)
	pl.amountPlanes = 0

	//we don't need the first row
	_, err := reader.Read()
	if err != nil {
		fmt.Println(err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		terminalNumberStr := record[6]
		terminalNumber, err1 := strconv.Atoi(terminalNumberStr)
		if err1 != nil {
			panic(err1)
		}

		var numberSeatsStr = record[9]
		numberSeats, err2 := strconv.Atoi(numberSeatsStr)
		if err2 != nil {
			panic(err2)
		}
		var planeClass byte
		switch {
		case numberSeats > 220:
			planeClass = 'W'
		case numberSeats > 120:
			planeClass = 'N'
		default:
			planeClass = 'R'
		}

		dateTimeStr := record[2]
		dateTime, err3 := time.Parse("02.01.2006 15:04", dateTimeStr)
		if err3 != nil {
			panic(err3)
		}

		var numberPassengersStr = record[10]
		numberPassengers, err3 := strconv.Atoi(numberPassengersStr)
		if err3 != nil {
			panic(err3)
		}
		var busesRequired int
		if numberPassengers % 80 == 0 {
			busesRequired = numberPassengers / 80
		} else {
			busesRequired = numberPassengers / 80 + 1
		}

		pl.data = append(pl.data, Plane{record[1][0], dateTime,
			record[5][0], terminalNumber, planeClass, busesRequired})
		pl.amountPlanes += 1
	}
}

func (pl *Planes) GetArrDepByPlaneId (n int) byte {
	return pl.data[n].AD
}

func (pl *Planes) GetDateTimeByPlaneId (n int) time.Time {
	return pl.data[n].dateTime
}

func (pl *Planes) GetIntDomByPlaneId (n int) byte {
	return pl.data[n].ID
}
func (pl *Planes) GetTerminalByPlaneId (n int) int {
	return pl.data[n].terminal
}
func (pl *Planes) GetClassByPlaneId (n int) byte {
	return pl.data[n].planeClass
}
func (pl *Planes) GetNumBusesByPlaneId (n int) int {
	return pl.data[n].busesRequired
}

package mymodule

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type PlanesInfo interface {
	LoadData (string)
	GetNumberOfPlanes () int
	GetArrDepByPlaneId (int) byte
	GetDateTimeByPlaneId (int) time.Time
	GetIntDomByPlaneId (int) byte
	GetTerminalByPlaneId (int) int
	GetClassByPlaneId (int) byte
	GetNumBusesByPlaneId (int) int
}

type ParkingPlacesInfo interface {
	LoadData (string)
	GetNumberOfParkingPlaces () int
	GetJetBridgeArrByPlaceId (int) byte
	GetJetBridgeDepByPlaceId (int) byte
	GetBusTimeToTerminal (int, int) int
	GetTerminalAttachedByPlaceId (int) int
	GetTaxiingTimeByPlaceId (int) int
}

type PPlace struct {

	//D - only for domestic, I - only for international
	//N - unavailable
	jbArrival byte

	//D - only for domestic, I - only for international
	//N - unavailable
	jbDeparture byte

	timeToTerm []int

	attachedTerminal int

	taxiingTime int

}

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

type ParkingPlaces struct {
	amountPlaces int
	data []PPlace
}

type Planes struct {
	amountPlanes int
	data []Plane
}

func (pl *Planes) GetNumberOfPlanes () int {
	return pl.amountPlanes
}

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
		terminalNumberStr := record[5]
		terminalNumber, err1 := strconv.Atoi(terminalNumberStr)
		if err1 != nil {
			panic(err1)
		}

		var numberSeatsStr = record[8]
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

		dateTimeStr := record[1]
		dateTime, err3 := time.Parse("2006-01-02 15:04:00", dateTimeStr)
		if err3 != nil {
			panic(err3)
		}

		var numberPassengersStr = record[9]
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
		
		pl.data = append(pl.data, Plane{record[0][0], dateTime,
			record[4][0], terminalNumber, planeClass, busesRequired})
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

func (pp *ParkingPlaces) LoadData (name string) {
	file := openFile(name)
	defer closeFile(file)
	reader := csv.NewReader(file)
	pp.amountPlaces = 0

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

		timeT1Str := record[3]
		timeT1, err1 := strconv.Atoi(timeT1Str)
		if err1 != nil {
			panic(err1)
		}

		timeT2Str := record[4]
		timeT2, err1 := strconv.Atoi(timeT2Str)
		if err1 != nil {
			panic(err1)
		}

		timeT3Str := record[5]
		timeT3, err1 := strconv.Atoi(timeT3Str)
		if err1 != nil {
			panic(err1)
		}

		timeT4Str := record[6]
		timeT4, err1 := strconv.Atoi(timeT4Str)
		if err1 != nil {
			panic(err1)
		}

		timeT5Str := record[7]
		timeT5, err1 := strconv.Atoi(timeT5Str)
		if err1 != nil {
			panic(err1)
		}

		attachedTermStr := record[8]
		var attachedTerm int
		if attachedTermStr != "" {
			attachedTerm, err1 = strconv.Atoi(attachedTermStr)
			if err1 != nil {
				panic(err1)
			}
		} else {
			attachedTerm = 0
		}

		taxiingTimeStr := record[9]
		taxiingTime, err1 := strconv.Atoi(taxiingTimeStr)
		if err1 != nil {
			panic(err1)
		}

		pp.data = append(pp.data, PPlace{record[1][0], record[2][0], []int{timeT1,
			timeT2, timeT3, timeT4, timeT5}, attachedTerm, taxiingTime})
		pp.amountPlaces += 1
	}
}

func (pp *ParkingPlaces) GetNumberOfParkingPlaces () int {
	return pp.amountPlaces
}
func (pp *ParkingPlaces) GetJetBridgeArrByPlaceId (n int) byte {
	return pp.data[n].jbArrival
}
func (pp *ParkingPlaces) GetJetBridgeDepByPlaceId (n int) byte {
	return pp.data[n].jbDeparture
}
func (pp *ParkingPlaces) GetBusTimeToTerminal (n int, t int) int {
	return pp.data[n].timeToTerm[t]
}
func (pp *ParkingPlaces) GetTerminalAttachedByPlaceId (n int) int {
	return pp.data[n].attachedTerminal
}
func (pp *ParkingPlaces) GetTaxiingTimeByPlaceId (n int) int {
	return pp.data[n].taxiingTime
}
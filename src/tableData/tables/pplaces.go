package tables

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

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

type ParkingPlaces struct {
	amountPlaces int
	data []PPlace
	matchPlacesNumbers []int
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

		realPPlaceNumberstr := record[0]
		realPPlaceNumber, err1 := strconv.Atoi(realPPlaceNumberstr)
		if err1 != nil {
			panic(err1)
		}

		pp.data = append(pp.data, PPlace{record[1][0], record[2][0], []int{timeT1,
			timeT2, timeT3, timeT4, timeT5}, attachedTerm, taxiingTime})
		pp.matchPlacesNumbers = append(pp.matchPlacesNumbers, realPPlaceNumber)
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
	return pp.data[n].timeToTerm[t - 1]
}
func (pp *ParkingPlaces) GetTerminalAttachedByPlaceId (n int) int {
	return pp.data[n].attachedTerminal
}
func (pp *ParkingPlaces) GetTaxiingTimeByPlaceId (n int) int {
	return pp.data[n].taxiingTime
}
func (pp *ParkingPlaces) GetMatchParkingPlaces () *[]int {
	return &(pp.matchPlacesNumbers)
}
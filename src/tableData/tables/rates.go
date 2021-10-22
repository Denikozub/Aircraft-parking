package tables

import (
	"encoding/csv"
	"fmt"
	"strconv"
)

type Rates struct {
	busCost int
	awayCost int
	jetBridgeCost int
	taxiingCost int
}

func (r * Rates) GetBusCost () int {
	return r.busCost
}

func (r * Rates) GetAwayStandCost () int {
	return r.awayCost
}

func (r * Rates) GetJetBridgeStandCost () int {
	return r.jetBridgeCost
}

func (r * Rates) GetTaxiingCost () int {
	return r.taxiingCost
}

func (r *Rates) LoadData (name string) {
	file := openFile(name)
	defer closeFile(file)
	reader := csv.NewReader(file)

	//we don't need the first row
	_, err := reader.Read()
	if err != nil {
		fmt.Println(err)
	}

	record,_ := reader.Read()
	busCost, err1 := strconv.Atoi(record[1])
	if err1 != nil {
		panic(err1)
	}
	r.busCost = busCost

	record,_ = reader.Read()
	awayCost, err1 := strconv.Atoi(record[1])
	if err1 != nil {
		panic(err1)
	}
	r.awayCost = awayCost

	record,_ = reader.Read()
	jetBridgeCost, err1 := strconv.Atoi(record[1])
	if err1 != nil {
		panic(err1)
	}
	r.jetBridgeCost = jetBridgeCost

	record,_ = reader.Read()
	taxiingCost, err1 := strconv.Atoi(record[1])
	if err1 != nil {
		panic(err1)
	}
	r.taxiingCost = taxiingCost
}

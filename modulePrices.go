package modulePrices

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type HandlingTime interface {
	LoadData (string)
	GetJetBridgeHandlingTimeByPlaneClass (byte) int
	GetAwayHandlingTimeByPlaneClass (byte) int
}

type HandlingRates interface {
	LoadData (string)
	GetBusCost () int
	GetAwayStandCost () int
	GetJetBridgeStandCost () int
	GetTaxiingCost () int
}

type Rates struct {
	busCost int
	awayCost int
	jetBridgeCost int
	taxiingCost int
}

type HTime struct {
	jetBridgeTime map[byte]int
	awayTime map[byte]int
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

func (t *HTime) GetJetBridgeHandlingTimeByPlaneClass (c byte) int {
	return t.jetBridgeTime[c]
}

func (t *HTime) GetAwayHandlingTimeByPlaneClass (c byte) int {
	return t.awayTime[c]
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

func (t *HTime) LoadData (name string)  {
	file := openFile(name)
	defer closeFile(file)
	reader := csv.NewReader(file)

	t.awayTime = make(map[byte]int)
	t.jetBridgeTime = make(map[byte]int)

	//we don't need the first row
	_, err := reader.Read()
	if err != nil {
		fmt.Println(err)
	}

	record,_ := reader.Read()
	jetBridgeTime, err1 := strconv.Atoi(record[1])
	if err1 != nil {
		panic(err1)
	}
	awayTime, err1 := strconv.Atoi(record[2])
	if err1 != nil {
		panic(err1)
	}
	t.jetBridgeTime['R'] = jetBridgeTime
	t.awayTime['R'] = awayTime

	record,_ = reader.Read()
	jetBridgeTime, err1 = strconv.Atoi(record[1])
	if err1 != nil {
		panic(err1)
	}
	awayTime, err1 = strconv.Atoi(record[2])
	if err1 != nil {
		panic(err1)
	}
	t.jetBridgeTime['N'] = jetBridgeTime
	t.awayTime['N'] = awayTime

	record,_ = reader.Read()
	jetBridgeTime, err1 = strconv.Atoi(record[1])
	if err1 != nil {
		panic(err1)
	}
	awayTime, err1 = strconv.Atoi(record[2])
	if err1 != nil {
		panic(err1)
	}
	t.jetBridgeTime['W'] = jetBridgeTime
	t.awayTime['W'] = awayTime
}
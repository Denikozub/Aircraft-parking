package tables

import (
	"encoding/csv"
	"fmt"
	"strconv"
)

type HTime struct {
	jetBridgeTime map[byte]int
	awayTime map[byte]int
}

func (t *HTime) GetJetBridgeHandlingTimeByPlaneClass (c byte) int {
	return t.jetBridgeTime[c]
}

func (t *HTime) GetAwayHandlingTimeByPlaneClass (c byte) int {
	return t.awayTime[c]
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

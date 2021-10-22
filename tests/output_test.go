package main

import "SVO.AERO/src/tableData/tables"

func main() {
	var arr []int
	for i := 0; i < 1009; i++ {
		arr = append(arr, i)
	}
	tables.WriteParkingPlacesToFile(arr, "SVO.AERO/data/Timetable_Public.csv", "output.csv")
}
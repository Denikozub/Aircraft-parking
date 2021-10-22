package main

import "SVO.AERO/src/tableData/tables"

func main() {
	var arr []int
	for i := 0; i < 1009; i++ {
		arr = append(arr, i)
	}

	folder := "C:/Users/kozub/go/src/SVO.AERO/data/"
	tables.WriteParkingPlacesToFile(arr, folder + "Timetable_Public.csv", folder + "output.csv")
}

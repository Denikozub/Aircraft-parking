package main

import "SVO.AERO/src/tableData/tables"

func main() {
	folder := "C:/Users/kozub/go/src/SVO.AERO/data/"

	x := tables.ParkingPlaces{}
	x.LoadData(folder + "Aircraft_Stands_Private.csv")
	var arr []int
	for i := 0; i < x.GetNumberOfParkingPlaces(); i++ {
		arr = append(arr, x.GetNumberOfParkingPlaces() - i - 1)
	}
	tables.WriteParkingPlacesToFile(&arr, x.GetMatchParkingPlaces(), folder + "Timetable_private.csv", folder + "output.csv")
}

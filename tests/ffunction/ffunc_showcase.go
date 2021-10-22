package main

import (
  "fmt"
  "SVO.AERO/src/FitnessFunction/abstract"
  "SVO.AERO/src/FitnessFunction/function"
  "SVO.AERO/src/tableData/tables"
)

func printCost(arr []int, ffunc abstract.FitnessFunction) {
  fmt.Println(ffunc.CalculateServiceCost(arr))
}

func main() {
  folder := "C:/Users/kozub/go/src/SVO.AERO/data/"
  planes := tables.Planes{}
  planes.LoadData(folder + "Timetable_Public.csv")
  pplaces := tables.ParkingPlaces{}
  pplaces.LoadData(folder + "Aircraft_Stands_Public.csv")
  timeHandling := tables.HTime{}
  timeHandling.LoadData(folder + "Handling_Time_Public.csv")
  ratesHandling := tables.Rates{}
  ratesHandling.LoadData(folder + "Handling_Rates_Public.csv")

  var arr []int
	for i := 0; i < 5; i++ {
		arr = append(arr, i)
	}
  ffunc := function.Function{}
  ffunc.Initialize(&ratesHandling, &timeHandling, &pplaces, &planes)
  printCost(arr, &ffunc)
}

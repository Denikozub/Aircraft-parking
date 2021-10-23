package main

import (
  "fmt"
  "SVO.AERO/src/fitnessFunction"
  "SVO.AERO/src/tableData/tables"
  "SVO.AERO/src/tableData/abstractTables"
)

func main() {
  folder := "C:/Users/kozub/go/src/SVO.AERO/data/"

  planes := tables.Planes{}
  planes.LoadData(folder + "Timetable_private.csv")
  pplaces := tables.ParkingPlaces{}
  pplaces.LoadData(folder + "Aircraft_Stands_Private.csv")
  timeHandling := tables.HTime{}
  timeHandling.LoadData(folder + "Handling_Time_Private.csv")
  ratesHandling := tables.Rates{}
  ratesHandling.LoadData(folder + "Handling_Rates_Private.csv")
  data := abstractTables.AirportData{&ratesHandling, &timeHandling, &pplaces, &planes}

  var arr []int
	for i := 0; i < 5; i++ {
		arr = append(arr, i)
	}
  fmt.Println(fitnessFunction.CalculateServiceCost(&data, &arr))
}

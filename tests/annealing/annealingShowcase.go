package main

import (
  "SVO.AERO/src/tableData/tables"
  "SVO.AERO/src/tableData/abstractTables"
  "SVO.AERO/src/parkingDistribution/distribution"
  "SVO.AERO/src/simulatedAnnealing"
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
  dist := distribution.Solution{}

  simulatedAnnealing.Anneal(&data, &dist, 100, 1000., 0.99, 1, folder + "Timetable_private.csv", folder + "output.csv")
}

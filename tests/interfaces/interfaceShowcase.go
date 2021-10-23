package main

import (
  "fmt"
  "SVO.AERO/src/tableData/abstractTables"
  "SVO.AERO/src/tableData/tables"
)

func printPlanesNum(planes abstractTables.PlanesInfo) {
  fmt.Println(planes.GetNumberOfPlanes())
}

func printParkingNum(pplaces abstractTables.ParkingPlacesInfo) {
  fmt.Println(pplaces.GetNumberOfParkingPlaces())
}

func printRegionalJetBridgeHandlingTime(timeHandling abstractTables.HandlingTime) {
  fmt.Println(timeHandling.GetJetBridgeHandlingTimeByPlaneClass('R'))
}

func printBusCost(ratesHandling abstractTables.HandlingRates) {
  fmt.Println(ratesHandling.GetBusCost())
}

func main() {
  folder := "C:/Users/kozub/go/src/SVO.AERO/data/"

  planes := tables.Planes{}
  planes.LoadData(folder + "Timetable_private.csv")
  printPlanesNum(&planes)

  pplaces := tables.ParkingPlaces{}
  pplaces.LoadData(folder + "Aircraft_Stands_Private.csv")
  printParkingNum(&pplaces)

  timeHandling := tables.HTime{}
  timeHandling.LoadData(folder + "Handling_Time_Private.csv")
  printRegionalJetBridgeHandlingTime(&timeHandling)

  ratesHandling := tables.Rates{}
  ratesHandling.LoadData(folder + "Handling_Rates_Private.csv")
  printBusCost(&ratesHandling)
}

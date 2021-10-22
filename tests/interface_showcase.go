package main

import (
  "fmt"
  "SVO.AERO/src/tableData/abstract"
  "SVO.AERO/src/tableData/tables"
)

func printPlanesNum(planes abstract.PlanesInfo) {
  fmt.Println(planes.GetNumberOfPlanes())
}

func printParkingNum(pplaces abstract.ParkingPlacesInfo) {
  fmt.Println(pplaces.GetNumberOfParkingPlaces())
}

func printRegionalJetBridgeHandlingTime(timeHandling abstract.HandlingTime) {
  fmt.Println(timeHandling.GetJetBridgeHandlingTimeByPlaneClass('R'))
}

func printBusCost(ratesHandling abstract.HandlingRates) {
  fmt.Println(ratesHandling.GetBusCost())
}

func main() {
  folder := "C:\\Users\\kozub\\go\\src\\SVO.AERO\\data\\"

  planes := tables.Planes{}
  planes.LoadData(folder + "Timetable_Public.csv")
  printPlanesNum(&planes)

  pplaces := tables.ParkingPlaces{}
  pplaces.LoadData(folder + "Aircraft_Stands_Public.csv")
  printParkingNum(&pplaces)

  timeHandling := tables.HTime{}
  timeHandling.LoadData(folder + "Handling_Time_Public.csv")
  printRegionalJetBridgeHandlingTime(&timeHandling)

  ratesHandling := tables.Rates{}
  ratesHandling.LoadData(folder + "Handling_Rates_Public.csv")
  printBusCost(&ratesHandling)
}

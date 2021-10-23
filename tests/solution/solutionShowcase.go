package main

import (
  "fmt"
  "SVO.AERO/src/tableData/tables"
  "SVO.AERO/src/tableData/abstractTables"
  "SVO.AERO/src/parkingDistribution/distribution"
  "SVO.AERO/src/parkingDistribution/abstractDistribution"
)

func neighbour(dist abstractDistribution.Distribution) {
  // folder := "C:/Users/kozub/go/src/SVO.AERO/data/"
  newDist := &dist
  (*newDist).ChangeDistribution(dist.GetNextNeighbourDistribution())
  // dist = *newDist
  // dist.SaveOutput(folder + "Timetable_private.csv", folder + "output.csv")
  fmt.Println(dist.FitnessValue())
  fmt.Println((*newDist).FitnessValue())
}

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
  dist.Initialize(&data)
  newDist := distribution.Solution{}
  newDist.Initialize(&data)
  newDist.ChangeDistribution(dist.GetNextNeighbourDistribution())
  fmt.Println(newDist.FitnessValue(), newDist.FitnessValue() - dist.FitnessValue())
  dist.ChangeDistribution(newDist.GetDistribution())
  fmt.Println(newDist.FitnessValue(), newDist.FitnessValue() - dist.FitnessValue())
}

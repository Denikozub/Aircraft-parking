package main

import (
  "SVO.AERO/src/FitnessFunction/function"
  "SVO.AERO/src/tableData/tables"
  "SVO.AERO/src/parkingDistribution/distribution"
  "SVO.AERO/src/parkingDistribution/abstractDistribution"
)

func neighbour(dist abstractDistribution.Distribution) {
  folder := "C:/Users/kozub/go/src/SVO.AERO/data/"
  new_dist := &dist
  (*new_dist).ChangeDistribution((*new_dist).GetNextNeighbour())
  (*new_dist).SaveToOutput(folder + "Timetable_Public.csv", folder + "output.csv")
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
  ffunc := function.Function{}

  sol := distribution.Solution{}
  sol.Initialize(&ratesHandling, &timeHandling, &pplaces, &planes, &ffunc)
  neighbour(&sol)
}

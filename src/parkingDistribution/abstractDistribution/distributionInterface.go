package abstractDistribution

import (
  "SVO.AERO/src/tableData/abstract"
  "SVO.AERO/src/FitnessFunction/abstractFunction"
)

type Distribution interface {
	Initialize(abstract.HandlingRates, abstract.HandlingTime, abstract.ParkingPlacesInfo,
    abstract.PlanesInfo, abstractFunction.FitnessFunction)
  ChangeDistribution([]int)
  GetNextNeighbour() []int
  CalculateFitnessValue() int
  SaveToOutput(string, string)
}

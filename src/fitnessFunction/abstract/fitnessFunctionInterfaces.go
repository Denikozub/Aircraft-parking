package abstract

import "SVO.AERO/src/tableData/abstract"

type FitnessFunction interface {
	Initialize(abstract.HandlingRates, abstract.HandlingTime, abstract.ParkingPlacesInfo, abstract.PlanesInfo)
	CalculateServiceCost ([]int) int
}

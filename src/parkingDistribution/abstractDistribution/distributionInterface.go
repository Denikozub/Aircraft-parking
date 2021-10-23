package abstractDistribution

import (
  "SVO.AERO/src/tableData/abstractTables"
)

type Distribution interface {
	Initialize(*abstractTables.AirportData)
  ChangeDistribution([]int)
  GetNextNeighbour() []int
  CalculateFitnessValue() int
  SaveOutput(string, string)
}

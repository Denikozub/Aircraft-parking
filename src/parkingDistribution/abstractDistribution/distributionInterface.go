package abstractDistribution

import (
  "SVO.AERO/src/tableData/abstractTables"
)

type Distribution interface {
	Initialize(*abstractTables.AirportData)
  GetDistribution() []int
  ChangeDistribution([]int)
  GetNextNeighbourDistribution() []int
  FitnessValue() int
  SaveOutput(string, string)
}

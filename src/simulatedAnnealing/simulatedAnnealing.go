package simulatedAnnealing

import (
  "SVO.AERO/src/tableData/abstractTables"
  "SVO.AERO/src/parkingDistribution/abstractDistribution"
)

func anneal(data *abstractTables.AirportData, dist *abstractDistribution.Distribution) {
  // dist.Initialize(data)
  // new_dist := &dist
  // for ...
    // for ...
      // (*new_dist).ChangeDistribution(dist.GetNextNeighbour())
      // new_dist_cost := (*new_dist).CalculateFitnessValue()
      // if ...
        // dist = *new_dist
      // else
        // ...
        // if ...
          // dist = *new_dist
  // dist.SaveToOutput(...)
}

package simulatedAnnealing

import (
  "SVO.AERO/src/tableData/abstractTables"
  "SVO.AERO/src/parkingDistribution/abstractDistribution"
  "rand"
  "math"
)

func Anneal(data *abstractTables.AirportData, dist *abstractDistribution.Distribution,
            max_iters int, init_temp float, anneal float, bolzman float, inputName string, outputName string) {

  dist.Initialize(data)
  new_dist := &dist
  temp := init_temp
  var delta float
  for i := 0; i < max_iters; i++ {
    (*new_dist).ChangeDistribution(dist.GetNextNeighbourDistribution())
    delta = (*new_dist).FitnessValue() - dist.FitnessValue()
    if delta < 0 || rand.Float64() < math.Exp(- delta / (bolzman * temp)) {
      dist.ChangeDistribution((*new_dist).GetDistribution())
    }
    temp *= anneal
  }
  dist.SaveOutput(inputName, outputName)
}

// dist := distribution.Solution{}
// Anneal(&data, &dist)

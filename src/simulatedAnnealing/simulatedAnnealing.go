package simulatedAnnealing

import (
  "SVO.AERO/src/tableData/abstractTables"
  "SVO.AERO/src/parkingDistribution/abstractDistribution"
  "math/rand"
  "math"
  "fmt"
)

func Anneal(data *abstractTables.AirportData, dist abstractDistribution.Distribution,
            max_iters int, init_temp float64, anneal float64, bolzman float64, inputName string, outputName string) {

  dist.Initialize(data)
  new_dist := dist
  temp := init_temp
  var delta int
  for i := 0; i < max_iters; i++ {
    new_dist.ChangeDistribution(dist.GetNextNeighbourDistribution())
    delta = new_dist.FitnessValue() - dist.FitnessValue()
    if delta < 0 || rand.Float64() < math.Exp(- float64(delta) / (bolzman * temp)) {
      dist.ChangeDistribution(new_dist.GetDistribution())
    }
    temp *= anneal
  }
  dist.SaveOutput(inputName, outputName)
  fmt.Println(dist.FitnessValue())
}

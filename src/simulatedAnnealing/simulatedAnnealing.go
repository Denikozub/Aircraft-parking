package simulatedAnnealing

import (
  "SVO.AERO/src/tableData/abstractTables"
  "SVO.AERO/src/parkingDistribution/distribution"
  "math/rand"
  "math"
  "fmt"
)

func Anneal(data *abstractTables.AirportData, max_iters int, init_temp float64,
    anneal float64, bolzman float64, inputName string, outputName string) {

  dist := distribution.Solution{}
  dist.Initialize(data)
  newDist := distribution.Solution{}
  newDist.Initialize(data)

  temp := init_temp
  var delta int
  for i := 0; i < max_iters; i++ {
    newDist.ChangeDistribution(dist.GetNextNeighbourDistribution())
    delta = newDist.FitnessValue() - dist.FitnessValue()
    if delta < 0 || rand.Float64() < math.Exp(- float64(delta) / (bolzman * temp)) {
      dist.ChangeDistribution(newDist.GetDistribution())
    }
    temp *= anneal
    fmt.Println(temp)
  }
  dist.SaveOutput(inputName, outputName)
  fmt.Println(dist.FitnessValue())
}

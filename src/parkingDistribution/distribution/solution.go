package distribution

import (
  "SVO.AERO/src/fitnessFunction"
  "SVO.AERO/src/tableData/abstractTables"
  "SVO.AERO/src/tableData/tables"
  "math/rand"
  "math"
)

type Solution struct {
	data *abstractTables.AirportData
  parkingNumber int
  planeNumber int
  distribution []int
}

func (sol * Solution) checkValidPPlace(dist []int, plane int, pplace int) bool {
  planeAD := sol.data.PlanesInfo.GetArrDepByPlaneId(plane)
  var hasJetBridge bool
  if planeAD == 'A' {
    hasJetBridge = sol.data.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(pplace) != 'N'
  } else {
    hasJetBridge = sol.data.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(pplace) != 'N'
  }
  for i := 0; i < len(dist); i++ {
    if i == plane {
      continue
    }
    // check time intersection TODO: mnogo vsego (taxiing time everywhere, no constant / matrix)
    if planeAD == sol.data.PlanesInfo.GetArrDepByPlaneId(i) {
      if math.Abs(sol.data.PlanesInfo.GetDateTimeByPlaneId(i).Sub(sol.data.PlanesInfo.GetDateTimeByPlaneId(plane)).Minutes()) > 60 {
        continue
      }
    } else if planeAD == 'A' {
      diff := sol.data.PlanesInfo.GetDateTimeByPlaneId(i).Sub(sol.data.PlanesInfo.GetDateTimeByPlaneId(plane)).Minutes()
      if diff <= 0 || diff > 2 * 60 {
        continue
      }
    } else {
      diff := sol.data.PlanesInfo.GetDateTimeByPlaneId(plane).Sub(sol.data.PlanesInfo.GetDateTimeByPlaneId(i)).Minutes()
      if diff <= 0 || diff > 2 * 60 {
        continue
      }
    }
    // take the same place
    if dist[i] == pplace {
      return false
    }
    // check wing intersection
    if hasJetBridge && (sol.data.PlanesInfo.GetTerminalByPlaneId(i) == sol.data.PlanesInfo.GetTerminalByPlaneId(plane)) &&
      (sol.data.PlanesInfo.GetClassByPlaneId(i) == 'W') && (sol.data.PlanesInfo.GetClassByPlaneId(plane) == 'W') &&
      ((dist[i] - pplace == 1) || (dist[i] - pplace == -1)) {
        return false
    }
  }
  return true
}

func (sol * Solution) Initialize(data *abstractTables.AirportData) {
	sol.data = data
  sol.parkingNumber = sol.data.ParkingPlacesInfo.GetNumberOfParkingPlaces()
  sol.planeNumber = sol.data.PlanesInfo.GetNumberOfPlanes()
  for i := 0; i < sol.planeNumber; i++ {
    var pplace int
    j, max_tries := 0, sol.parkingNumber * 3
    for ok := true; ok; ok = !sol.checkValidPPlace(sol.distribution[:i], i, pplace) {
      pplace = rand.Intn(sol.parkingNumber)
      j++
      if j > max_tries {
        panic(1)
      }
    }
    sol.distribution = append(sol.distribution, pplace)
  }
}

func (sol * Solution) ChangeDistribution(new_dist []int) {
  if len(new_dist) != len(sol.distribution) {
    panic("Array lengths are different!")
  }
  for i := 0; i < len(sol.distribution); i++ {
    sol.distribution[i] = new_dist[i]
  }
}

func (sol * Solution) GetNextNeighbour() []int {
  var new_dist []int
  for i := 0; i < len(sol.distribution); i++ {
    new_dist = append(new_dist, sol.distribution[i])
  }
  plane := rand.Intn(sol.planeNumber)
  var pplace int
  j, max_tries := 0, sol.parkingNumber * 3
  for ok := true; ok; ok = !sol.checkValidPPlace(sol.distribution, plane, pplace) {
    pplace = rand.Intn(sol.parkingNumber)
    j++
    if j > max_tries {
      panic(1)
    }
  }
  new_dist[plane] = pplace
  return new_dist
}

func (sol * Solution) CalculateFitnessValue() int {
  return fitnessFunction.CalculateServiceCost(sol.data, sol.distribution)
}

func (sol * Solution) SaveToOutput(inputName string, outputName string) {
  tables.WriteParkingPlacesToFile(sol.distribution, inputName, outputName)
}

package distribution

import (
  "SVO.AERO/src/FitnessFunction/abstractFunction"
  "SVO.AERO/src/tableData/abstract"
  "SVO.AERO/src/tableData/tables"
  "math/rand"
  "math"
)

type Solution struct {
	handlingRates abstract.HandlingRates
	handlingTime abstract.HandlingTime
	parkingPlacesInfo  abstract.ParkingPlacesInfo
	planesInfo  abstract.PlanesInfo
  fitnessFunction abstractFunction.FitnessFunction
  distribution []int
}

func (sol * Solution) checkValidPPlace(dist []int, plane int, pplace int) bool {
  planeAD := sol.planesInfo.GetArrDepByPlaneId(plane)
  var hasJetBridge bool
  if planeAD == 'A' {
    hasJetBridge = sol.parkingPlacesInfo.GetJetBridgeArrByPlaceId(pplace) != 'N'
  } else {
    hasJetBridge = sol.parkingPlacesInfo.GetJetBridgeDepByPlaceId(pplace) != 'N'
  }
  for i := 0; i < len(dist); i++ {
    if i == plane {
      continue
    }
    // check time intersection TODO: mnogo vsego (taxiing time everywhere, no constant / matrix)
    if planeAD == sol.planesInfo.GetArrDepByPlaneId(i) {
      if math.Abs(sol.planesInfo.GetDateTimeByPlaneId(i).Sub(sol.planesInfo.GetDateTimeByPlaneId(plane)).Minutes()) > 60 {
        continue
      }
    } else if planeAD == 'A' {
      diff := sol.planesInfo.GetDateTimeByPlaneId(i).Sub(sol.planesInfo.GetDateTimeByPlaneId(plane)).Minutes()
      if diff <= 0 || diff > 2 * 60 {
        continue
      }
    } else {
      diff := sol.planesInfo.GetDateTimeByPlaneId(plane).Sub(sol.planesInfo.GetDateTimeByPlaneId(i)).Minutes()
      if diff <= 0 || diff > 2 * 60 {
        continue
      }
    }
    // take the same place
    if dist[i] == pplace {
      return false
    }
    // check wing intersection
    if hasJetBridge && (sol.planesInfo.GetTerminalByPlaneId(i) == sol.planesInfo.GetTerminalByPlaneId(plane)) &&
      (sol.planesInfo.GetClassByPlaneId(i) == 'W') && (sol.planesInfo.GetClassByPlaneId(plane) == 'W') &&
      ((dist[i] - pplace == 1) || (dist[i] - pplace == -1)) {
        return false
    }
  }
  return true
}

func (sol * Solution) Initialize(handlingRates abstract.HandlingRates,
								handlingTime abstract.HandlingTime,
								parkingPlacesInfo  abstract.ParkingPlacesInfo,
								planesInfo  abstract.PlanesInfo,
                fitnessFunction abstractFunction.FitnessFunction) {
	sol.handlingRates = handlingRates
	sol.handlingTime =  handlingTime
	sol.parkingPlacesInfo = parkingPlacesInfo
	sol.planesInfo  = planesInfo
  sol.fitnessFunction = fitnessFunction
  sol.fitnessFunction.Initialize(handlingRates, handlingTime, parkingPlacesInfo, planesInfo)
  for i := 0; i < sol.planesInfo.GetNumberOfPlanes(); i++ {
    var pplace int
    j, max_tries := 0, sol.parkingPlacesInfo.GetNumberOfParkingPlaces() * 3
    for ok := true; ok; ok = !sol.checkValidPPlace(sol.distribution[:i], i, pplace) {
      pplace = rand.Intn(sol.parkingPlacesInfo.GetNumberOfParkingPlaces())
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
    panic(1)
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
  plane := rand.Intn(sol.planesInfo.GetNumberOfPlanes())
  var pplace int
  j, max_tries := 0, sol.parkingPlacesInfo.GetNumberOfParkingPlaces() * 3
  for ok := true; ok; ok = !sol.checkValidPPlace(sol.distribution, plane, pplace) {
    pplace = rand.Intn(sol.parkingPlacesInfo.GetNumberOfParkingPlaces())
    j++
    if j > max_tries {
      panic(1)
    }
  }
  new_dist[plane] = pplace
  return new_dist
}

func (sol * Solution) CalculateFitnessValue() int {
  return sol.fitnessFunction.CalculateServiceCost(sol.distribution)
}

func (sol * Solution) SaveToOutput(inputName string, outputName string) {
  tables.WriteParkingPlacesToFile(sol.distribution, inputName, outputName)
}

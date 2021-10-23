package distribution

import (
  "SVO.AERO/src/fitnessFunction"
  "SVO.AERO/src/tableData/abstractTables"
  "SVO.AERO/src/tableData/tables"
  "math/rand"
)

type Solution struct {
	data *abstractTables.AirportData
  parkingNumber int
  planeNumber int
  distribution []int
  fitnessValue int
}

func (sol * Solution) windIntersection(plane int, pplace int, i int , iplace int) bool {
  planeAD := sol.data.PlanesInfo.GetArrDepByPlaneId(plane)
  var planeHasJetBridge bool
  if planeAD == 'A' {
    planeHasJetBridge = sol.data.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(pplace) != 'N'
  } else {
    planeHasJetBridge = sol.data.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(pplace) != 'N'
  }
  return planeHasJetBridge &&
    (sol.data.PlanesInfo.GetTerminalByPlaneId(i) == sol.data.PlanesInfo.GetTerminalByPlaneId(plane)) &&
      ((sol.data.PlanesInfo.GetClassByPlaneId(i) == 'W') && (sol.data.PlanesInfo.GetClassByPlaneId(plane) == 'W')) &&
      ((iplace - pplace == 1) || (iplace - pplace == -1))
}

func (sol * Solution) timeIntersection(plane int, pplace int, i int , iplace int) bool {
  planeAD := sol.data.PlanesInfo.GetArrDepByPlaneId(plane)
  planeData := sol.data.PlanesInfo.GetDateTimeByPlaneId(plane)
  iAD := sol.data.PlanesInfo.GetArrDepByPlaneId(i)
  iData := sol.data.PlanesInfo.GetDateTimeByPlaneId(i)
  planeTaxiing := sol.data.ParkingPlacesInfo.GetTaxiingTimeByPlaceId(pplace)
  iTaxiing := sol.data.ParkingPlacesInfo.GetTaxiingTimeByPlaceId(iplace)
  var planeHandling, iHandling int

  if planeAD == 'A' {
    if (sol.data.PlanesInfo.GetTerminalByPlaneId(plane) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(pplace)) &&
      (sol.data.PlanesInfo.GetIntDomByPlaneId(plane) == sol.data.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(pplace)) {
      planeHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(plane))
    } else {
      planeHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(plane))
    }
    if planeAD == iAD {
      if (sol.data.PlanesInfo.GetTerminalByPlaneId(i) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(iplace)) &&
          (sol.data.PlanesInfo.GetIntDomByPlaneId(i) == sol.data.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(iplace)) {
        iHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      } else {
        iHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      }
      return ((planeData.Sub(iData).Minutes() > 0) && (planeData.Sub(iData).Minutes() > float64(iTaxiing + iHandling -planeTaxiing))) ||
          ((planeData.Sub(iData).Minutes() < 0) && (planeData.Sub(iData).Minutes() > float64(planeTaxiing + planeHandling - iTaxiing)))
      } else {
      if (sol.data.PlanesInfo.GetTerminalByPlaneId(i) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(iplace)) &&
          (sol.data.PlanesInfo.GetIntDomByPlaneId(i) == sol.data.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(iplace)) {
        iHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      } else {
        iHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      }
      if  planeData.Sub(iData).Minutes() >= 0 {
        return false
      } else {
        return (iData.Sub(planeData).Minutes() <= float64(planeTaxiing + iTaxiing))  ||
          (iData.Sub(planeData).Minutes() >= float64(iTaxiing + planeTaxiing + planeHandling + iHandling))
      }
    }
  } else {
    if (sol.data.PlanesInfo.GetTerminalByPlaneId(plane) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(pplace)) &&
        (sol.data.PlanesInfo.GetIntDomByPlaneId(plane) == sol.data.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(pplace)) {
      planeHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(plane))
    } else {
      planeHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(plane))
    }
    if planeAD == iAD {
      if (sol.data.PlanesInfo.GetTerminalByPlaneId(i) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(iplace)) &&
          (sol.data.PlanesInfo.GetIntDomByPlaneId(i) == sol.data.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(iplace)) {
        iHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      } else {
        iHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      }
      return ((planeData.Sub(iData).Minutes() > 0) && (planeData.Sub(iData).Minutes() < float64(planeTaxiing + planeHandling - iTaxiing))) ||
          ((planeData.Sub(iData).Minutes() < 0) && (planeData.Sub(iData).Minutes() < float64(iTaxiing + iHandling - planeTaxiing)))
    } else {
      if (sol.data.PlanesInfo.GetTerminalByPlaneId(i) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(iplace)) &&
          (sol.data.PlanesInfo.GetIntDomByPlaneId(i) == sol.data.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(iplace)) {
        iHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      } else {
        iHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      }
      if  planeData.Sub(iData).Minutes() <= 0 {
        return false
      } else {
        return (iData.Sub(planeData).Minutes() <= float64(planeTaxiing + iTaxiing))  ||
            (iData.Sub(planeData).Minutes() >= float64(iTaxiing + planeTaxiing + planeHandling + iHandling))
      }
    }
  }
}

func (sol * Solution) checkValidPPlace(dist []int, plane int, pplace int) bool {
  for i := 0; i < len(dist); i++ {
    if i == plane {
      continue
    }
    if sol.timeIntersection(plane, pplace, i, dist[i]) {
      if dist[i] == pplace {
        return false
      }
      if sol.windIntersection(plane, pplace, i, dist[i]) {
        return false
      }
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
  sol.fitnessValue = fitnessFunction.CalculateServiceCost(sol.data, &sol.distribution)
}

func (sol * Solution) ChangeDistribution(new_dist []int) {
  if len(new_dist) != len(sol.distribution) {
    panic("Array lengths are different!")
  }
  for i := 0; i < len(sol.distribution); i++ {
    sol.distribution[i] = new_dist[i]
  }
  sol.fitnessValue = fitnessFunction.CalculateServiceCost(sol.data, &sol.distribution)
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
  // return copy
  return new_dist
}

func (sol * Solution) CalculateFitnessValue() int {
  return sol.fitnessValue
}

func (sol * Solution) SaveOutput(inputName string, outputName string) {
  tables.WriteParkingPlacesToFile(&sol.distribution,
      sol.data.ParkingPlacesInfo.GetMatchParkingPlaces(), inputName, outputName)
}

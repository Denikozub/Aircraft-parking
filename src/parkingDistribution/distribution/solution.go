package distribution

import (
  "SVO.AERO/src/fitnessFunction"
  "SVO.AERO/src/tableData/abstractTables"
  "SVO.AERO/src/tableData/tables"
  "math/rand"
  "time"
  "strconv"
  "fmt"
)

type Solution struct {
	data *abstractTables.AirportData
  parkingNumber int
  planeNumber int
  distribution []int
  fitnessValue int
}

func (sol * Solution) wingIntersection(plane int, pplace int, i int , iplace int) bool {
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
  var terminalUse bool
  if planeAD == 'A' {
    terminalUse = (sol.data.PlanesInfo.GetTerminalByPlaneId(plane) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(pplace)) &&
        (sol.data.PlanesInfo.GetIntDomByPlaneId(plane) == sol.data.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(pplace))
    if terminalUse {
      planeHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(plane))
    } else {
      planeHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(plane))
    }
    if planeAD == iAD {
      terminalUse = (sol.data.PlanesInfo.GetTerminalByPlaneId(i) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(iplace)) &&
          (sol.data.PlanesInfo.GetIntDomByPlaneId(i) == sol.data.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(iplace))
      if terminalUse {
        iHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      } else {
        iHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      }
      return !((planeData.Sub(iData).Minutes() >= 0) && (planeData.Sub(iData).Minutes() <= float64(iTaxiing + iHandling -planeTaxiing))) &&
          !((planeData.Sub(iData).Minutes() < 0) && (planeData.Sub(iData).Minutes() <= float64(planeTaxiing + planeHandling - iTaxiing)))
    } else {
      terminalUse = (sol.data.PlanesInfo.GetTerminalByPlaneId(i) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(iplace)) &&
          (sol.data.PlanesInfo.GetIntDomByPlaneId(i) == sol.data.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(iplace))
      if  terminalUse {
        iHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      } else {
        iHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      }
      return planeData.Sub(iData).Minutes() <= 0 && !(iData.Sub(planeData).Minutes() <= float64(planeTaxiing + iTaxiing)) &&
          !(iData.Sub(planeData).Minutes() >= float64(iTaxiing + planeTaxiing + planeHandling + iHandling))
    }
  } else {
    // посчитать handling plane
    terminalUse =  (sol.data.PlanesInfo.GetTerminalByPlaneId(plane) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(pplace)) &&
        (sol.data.PlanesInfo.GetIntDomByPlaneId(plane) == sol.data.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(pplace))
    if terminalUse {
      planeHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(plane))
    } else {
      planeHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(plane))
    }
    if planeAD == iAD {
      // посчитать handling i
      terminalUse = (sol.data.PlanesInfo.GetTerminalByPlaneId(i) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(iplace)) &&
          (sol.data.PlanesInfo.GetIntDomByPlaneId(i) == sol.data.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(iplace))
      if terminalUse {
        iHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      } else {
        iHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
      }
      return !((planeData.Sub(iData).Minutes() >= 0) && (planeData.Sub(iData).Minutes() <= float64(planeTaxiing + planeHandling - iTaxiing))) &&
          !((planeData.Sub(iData).Minutes() < 0) && (planeData.Sub(iData).Minutes() <= float64(iTaxiing + iHandling - planeTaxiing)))
    } else {
      // посчитать handling plane
      terminalUse = (sol.data.PlanesInfo.GetTerminalByPlaneId(i) == sol.data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(iplace)) &&
          (sol.data.PlanesInfo.GetIntDomByPlaneId(i) == sol.data.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(iplace))
      if terminalUse {
        iHandling = sol.data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
          } else {
            iHandling = sol.data.HandlingTime.GetAwayHandlingTimeByPlaneClass(sol.data.PlanesInfo.GetClassByPlaneId(i))
          }
      return planeData.Sub(iData).Minutes() > 0 && !(iData.Sub(planeData).Minutes() <= float64(planeTaxiing + iTaxiing)) &&
          !(iData.Sub(planeData).Minutes() >= float64(iTaxiing + planeTaxiing + planeHandling + iHandling))
    }
  }
}


func (sol * Solution) checkValidPPlace(dist []int, plane int, pplace int) bool {
  for i := 0; i < len(dist); i++ {
    if i == plane {
      continue
    }
    if (dist[i] == pplace || sol.wingIntersection(plane, pplace, i, dist[i])) && sol.timeIntersection(plane, pplace, i, dist[i]) {
      return false
    }
  }
  return true
}

func (sol * Solution) Initialize(data *abstractTables.AirportData) {
	sol.data = data
  sol.parkingNumber = sol.data.ParkingPlacesInfo.GetNumberOfParkingPlaces()
  sol.planeNumber = sol.data.PlanesInfo.GetNumberOfPlanes()

  rand.Seed(time.Now().UnixNano())
  var pplace int
  maxTries := sol.parkingNumber * 3

  for i := 0; i < sol.planeNumber; i++ {
    for ok, j := true, 0; ok; ok, j = !sol.checkValidPPlace(sol.distribution[:i], i, pplace), j+1 {
      pplace = rand.Intn(sol.parkingNumber)
      if j > maxTries {
        panic("Unable to init plane " + strconv.Itoa(i))
      }
    }
    sol.distribution = append(sol.distribution, pplace)
  }
  sol.fitnessValue = fitnessFunction.CalculateServiceCost(sol.data, &sol.distribution)
}

func (sol * Solution) GetDistribution() []int {
  return sol.distribution
}

func (sol * Solution) ChangeDistribution(newDist []int) {
  if len(newDist) != len(sol.distribution) {
    panic("Array lengths are different!")
  }
  copy(sol.distribution, newDist)
  sol.fitnessValue = fitnessFunction.CalculateServiceCost(sol.data, &sol.distribution)
}

func (sol * Solution) GetNextNeighbourDistribution() []int {
  newDist := make([]int, len(sol.distribution))
  copy(newDist, sol.distribution)

  rand.Seed(time.Now().UnixNano())
  var pplace, plane int
  maxPplaceTries := sol.parkingNumber * 3
  maxNeighbourTries := 50

  for k := 0; k < maxNeighbourTries; k++ {
    plane := rand.Intn(sol.planeNumber)
    limitReached := false
    for ok, j := true, 0; ok && !limitReached; ok, j = !sol.checkValidPPlace(sol.distribution, plane, pplace), j+1 {
      pplace = rand.Intn(sol.parkingNumber)
      limitReached = j >= maxPplaceTries
    }
    if !limitReached {
      break
    }
    if k == maxNeighbourTries - 1 {
      panic("Unable to get neighbout. "  + strconv.Itoa(maxNeighbourTries) + " tries failed.")
    }
    fmt.Println("Unable to alter plane " + strconv.Itoa(plane) + ". Trying next one...")
  }

  newDist[plane] = pplace
  return newDist
}

func (sol * Solution) FitnessValue() int {
  return sol.fitnessValue
}

func (sol * Solution) SaveOutput(inputName string, outputName string) {
  tables.WriteParkingPlacesToFile(&sol.distribution,
      sol.data.ParkingPlacesInfo.GetMatchParkingPlaces(), inputName, outputName)
}

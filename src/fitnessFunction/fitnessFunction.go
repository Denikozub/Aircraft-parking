package fitnessFunction

import "SVO.AERO/src/tableData/abstractTables"

func calculatePlaneCost (data *abstractTables.AirportData, plane int, place int ) int {

	// сразу считаем taxiing cost
	sum := data.HandlingRates.GetTaxiingCost() * data.ParkingPlacesInfo.GetTaxiingTimeByPlaceId(place)

	// standing cost
	planeTerminal  := data.PlanesInfo.GetTerminalByPlaneId(plane)
	planeClass, planeID := data.PlanesInfo.GetClassByPlaneId(plane), data.PlanesInfo.GetIntDomByPlaneId(plane)
	placeTerminal := data.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(place)
	busCost := data.PlanesInfo.GetNumBusesByPlaneId(plane) * data.HandlingRates.GetBusCost() *
		data.ParkingPlacesInfo.GetBusTimeToTerminal(place, planeTerminal)
	planeAD := data.PlanesInfo.GetArrDepByPlaneId(plane)
	if planeAD == 'A' {
		if data.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(place) == planeID {
			if planeTerminal == placeTerminal {
				sum += data.HandlingRates.GetJetBridgeStandCost() *
					data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(planeClass)
			} else {
				sum += data.HandlingRates.GetJetBridgeStandCost() *
					data.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass) + busCost
			}
		} else {
			sum += busCost
			if data.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(place) == 'N' {
				sum += data.HandlingRates.GetAwayStandCost() *
					data.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass)
			} else {
				sum += data.HandlingRates.GetJetBridgeStandCost() *
					data.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass)
			}
		}
	} else if planeAD == 'D' {

		if data.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(place) == planeID {
			if planeTerminal == placeTerminal {
				sum += data.HandlingRates.GetJetBridgeStandCost() *
					data.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(planeClass)
			} else {
				sum += data.HandlingRates.GetJetBridgeStandCost() *
					data.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass) + busCost
			}
		} else {
			sum += busCost
			if data.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(place) == 'N' {
				sum += data.HandlingRates.GetAwayStandCost() *
					data.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass)
			} else {
				sum += data.HandlingRates.GetJetBridgeStandCost() *
					data.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass)
			}
		}
	}
	return sum
}

func CalculateServiceCost (data *abstractTables.AirportData, planes *[]int) int {

	cost, i := 0, 0
	for ; i < len(*planes); i++ {
		cost += calculatePlaneCost(data, i, (*planes)[i])
	}
	return cost
}

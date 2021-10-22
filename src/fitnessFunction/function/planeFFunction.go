package function

import "SVO.AERO/src/tableData/abstract"

type Function struct {
	HandlingRates abstract.HandlingRates
	HandlingTime abstract.HandlingTime
	ParkingPlacesInfo  abstract.ParkingPlacesInfo
	PlanesInfo  abstract.PlanesInfo
}

func (ff * Function) Initialize( HandlingRates abstract.HandlingRates,
								HandlingTime abstract.HandlingTime,
								ParkingPlacesInfo  abstract.ParkingPlacesInfo,
								PlanesInfo  abstract.PlanesInfo) {
	ff.HandlingRates = HandlingRates
	ff.HandlingTime =  HandlingTime
	ff.ParkingPlacesInfo = ParkingPlacesInfo
	ff.PlanesInfo  = PlanesInfo
}

func (ff * Function) calculatePlaneCost (plane int, place int ) int {

	// сразу считаем taxiing cost
	sum := ff.HandlingRates.GetTaxiingCost() * ff.ParkingPlacesInfo.GetTaxiingTimeByPlaceId(place)

	// standing cost
	planeTerminal  := ff.PlanesInfo.GetTerminalByPlaneId(plane)
	planeClass, planeID := ff.PlanesInfo.GetClassByPlaneId(plane), ff.PlanesInfo.GetIntDomByPlaneId(plane)
	placeTerminal := ff.ParkingPlacesInfo.GetTerminalAttachedByPlaceId(place)
	busCost := ff.PlanesInfo.GetNumBusesByPlaneId(plane) * ff.HandlingRates.GetBusCost() *
		ff.ParkingPlacesInfo.GetBusTimeToTerminal(place, planeTerminal - 1)
	planeAD := ff.PlanesInfo.GetArrDepByPlaneId(plane)
	if planeAD == 'A' {
		if ff.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(place) == planeID {
			if planeTerminal == placeTerminal {
				sum += ff.HandlingRates.GetJetBridgeStandCost() *
					ff.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(planeClass)
			} else {
				sum += ff.HandlingRates.GetJetBridgeStandCost() *
					ff.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass) + busCost
			}
		} else {
			sum += busCost
			if ff.ParkingPlacesInfo.GetJetBridgeArrByPlaceId(place) == 'N' {
				sum += ff.HandlingRates.GetAwayStandCost() * ff.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass)
			} else {
				sum += ff.HandlingRates.GetJetBridgeStandCost() *
					ff.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass)
			}
		}
	} else if planeAD == 'D' {

		if ff.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(place) == planeID {
			if planeTerminal == placeTerminal {
				sum += ff.HandlingRates.GetJetBridgeStandCost() *
					ff.HandlingTime.GetJetBridgeHandlingTimeByPlaneClass(planeClass)
			} else {
				sum += ff.HandlingRates.GetJetBridgeStandCost() *
					ff.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass) + busCost
			}
		} else {
			sum += busCost
			if ff.ParkingPlacesInfo.GetJetBridgeDepByPlaceId(place) == 'N' {
				sum += ff.HandlingRates.GetAwayStandCost()*ff.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass)
			} else {
				sum += ff.HandlingRates.GetJetBridgeStandCost() *
					ff.HandlingTime.GetAwayHandlingTimeByPlaneClass(planeClass)
			}
		}
	}
	return sum
}

func (ff * Function) CalculateServiceCost (planes []int) int {
	cost, i := 0, 0
	for ; i < len(planes); i++ {
		cost += ff.calculatePlaneCost(i, planes[i])
	}
	return cost
}

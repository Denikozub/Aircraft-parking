package abstractTables

import "time"

type PlanesInfo interface {
	LoadData (string)
	GetNumberOfPlanes () int
	GetArrDepByPlaneId (int) byte
	GetDateTimeByPlaneId (int) time.Time
	GetIntDomByPlaneId (int) byte
	GetTerminalByPlaneId (int) int
	GetClassByPlaneId (int) byte
	GetNumBusesByPlaneId (int) int
}

type ParkingPlacesInfo interface {
	LoadData (string)
	GetNumberOfParkingPlaces () int
	GetJetBridgeArrByPlaceId (int) byte
	GetJetBridgeDepByPlaceId (int) byte
	GetBusTimeToTerminal (int, int) int
	GetTerminalAttachedByPlaceId (int) int
	GetTaxiingTimeByPlaceId (int) int
	GetMatchParkingPlaces () *[]int
}

type HandlingTime interface {
	LoadData (string)
	GetJetBridgeHandlingTimeByPlaneClass (byte) int
	GetAwayHandlingTimeByPlaneClass (byte) int
}

type HandlingRates interface {
	LoadData (string)
	GetBusCost () int
	GetAwayStandCost () int
	GetJetBridgeStandCost () int
	GetTaxiingCost () int
}

type AirportData struct {
	HandlingRates HandlingRates
	HandlingTime HandlingTime
	ParkingPlacesInfo ParkingPlacesInfo
	PlanesInfo PlanesInfo
}

package main

type StationInfo struct {
	stNm            string
	arsId           string
	stationId       string
}

type BusInfo struct {
	busRouteId      string
	busRouteNm      string
}

type UserDB struct {
	DestAddr     string      `json:"destAddr"`
	BusStop      string      `json:"busStop"`
	BusNum       string      `json:"busNum"`
	ZipName      string      `json:"zipName"`
	SpecialDay   string      `json:"specialDay"`
	StockName    string      `json:"stockName"`
	BusInfo      BusInfo     `json:"busInfo"`
	StationInfo  StationInfo `json:"stationInfo"`
}

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

type Town struct {
	LocationX     string
	LocationY     string
}

type ZipInfo struct {
	TownName      string
	TownCode      string
}

type UserDB struct {
	TimeStamp    string      `json:"timeStamp"`
	HomeAddr     string      `json:"homeAddr"`
	DestAddr     string      `json:"destAddr"`
	BusStop      string      `json:"busStop"`
	BusNum       string      `json:"busNum"`
	ZipName      string      `json:"zipName"`
	SpecialDay   string      `json:"specialDay"`
	StockName    string      `json:"stockName"`
	BusInfo      BusInfo     `json:"busInfo"`
	StationInfo  StationInfo `json:"stationInfo"`
	HomeTown     Town        `json:"homeTown"`
	DestTown     Town        `json:"destTown"`
	ZipInfo      ZipInfo     `json:"zipInfo"`
}

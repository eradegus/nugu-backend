package jsonstruct

type OpenAPIResponse struct {
        ComMsgHeader ComMsgHeader `json:"comMsgHeader"`
        MsgHeader    MsgHeader    `json:"msgHeader"`
        MsgBody      MsgBody      `json:"msgBody"`
}

type ComMsgHeader struct {
        ResponseMsgID interface{} `json:"responseMsgID"`
        ResponseTime  interface{} `json:"responseTime"`
        RequestMsgID  interface{} `json:"requestMsgID"`
        SuccessYN     interface{} `json:"successYN"`
        ReturnCode    interface{} `json:"returnCode"`
        ErrMsg        interface{} `json:"errMsg"`
}

type MsgHeader struct {
        HeaderMsg string `json:"headerMsg"`
        HeaderCd  string `json:"headerCd"`
        ItemCount int    `json:"itemCount"`
}

type MsgBody struct {
        ItemList []ItemList `json:"itemList"`
}

type ItemList struct {
        StID         string      `json:"stId"`
        StNm         string      `json:"stNm"`
        TmX          string      `json:"tmX"`
        TmY          string      `json:"tmY"`
        PosX         string      `json:"posX"`
        PosY         string      `json:"posY"`
        ArsID        string      `json:"arsId"`
        BusRouteID   string      `json:"busRouteId"`
        RtNm         string      `json:"rtNm"`
        BusRouteAbrv string      `json:"busRouteAbrv"`
        SectNm       string      `json:"sectNm"`
        GpsX         string      `json:"gpsX"`
        GpsY         string      `json:"gpsY"`
        StationTp    string      `json:"stationTp"`
        FirstTm      string      `json:"firstTm"`
        LastTm       string      `json:"lastTm"`
        Term         string      `json:"term"`
        RouteType    string      `json:"routeType"`
        NextBus      string      `json:"nextBus"`
        StaOrd       string      `json:"staOrd"`
        VehID1       string      `json:"vehId1"`
        PlainNo1     interface{} `json:"plainNo1"`
        SectOrd1     string      `json:"sectOrd1"`
        StationNm1   string      `json:"stationNm1"`
        TraTime1     string      `json:"traTime1"`
        TraSpd1      string      `json:"traSpd1"`
        IsArrive1    string      `json:"isArrive1"`
        RepTm1       string      `json:"repTm1"`
        IsLast1      string      `json:"isLast1"`
        BusType1     string      `json:"busType1"`
        VehID2       string      `json:"vehId2"`
        PlainNo2     interface{} `json:"plainNo2"`
        SectOrd2     string      `json:"sectOrd2"`
        StationNm2   string      `json:"stationNm2"`
        TraTime2     string      `json:"traTime2"`
        TraSpd2      string      `json:"traSpd2"`
        IsArrive2    string      `json:"isArrive2"`
        RepTm2       interface{} `json:"repTm2"`
        IsLast2      string      `json:"isLast2"`
        BusType2     string      `json:"busType2"`
        Adirection   string      `json:"adirection"`
        Arrmsg1      string      `json:"arrmsg1"`
        Arrmsg2      string      `json:"arrmsg2"`
        ArrmsgSec1   string      `json:"arrmsgSec1"`
        ArrmsgSec2   string      `json:"arrmsgSec2"`
        NxtStn       string      `json:"nxtStn"`
        RerdieDiv1   string      `json:"rerdieDiv1"`
        RerdieDiv2   string      `json:"rerdieDiv2"`
        RerideNum1   string      `json:"rerideNum1"`
        RerideNum2   string      `json:"rerideNum2"`
        IsFullFlag1  string      `json:"isFullFlag1"`
        IsFullFlag2  string      `json:"isFullFlag2"`
        DeTourAt     string      `json:"deTourAt"`
        Congestion   string      `json:"congestion"`
        BusRouteNm   string      `json:"busRouteNm"`
        Length       string      `json:"length"`
        StStationNm  string      `json:"stStationNm"`
        EdStationNm  string      `json:"edStationNm"`
        LastBusYn    string      `json:"lastBusYn"`
        LastBusTm    string      `json:"lastBusTm"`
        FirstBusTm   string      `json:"firstBusTm"`
        LastLowTm    string      `json:"lastLowTm"`
        FirstLowTm   string      `json:"firstLowTm"`
        CorpNm       string      `json:"corpNm"`
}


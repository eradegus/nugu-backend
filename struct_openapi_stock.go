package main

type OpenAPIResponse_stock struct {
	Response struct {
		Header struct {
			ResultCode string `json:"resultCode"`
			ResultMsg  string `json:"resultMsg"`
		} `json:"header"`
		Body struct {
			NumOfRows  int `json:"numOfRows"`
			PageNo     int `json:"pageNo"`
			TotalCount int `json:"totalCount"`
			Items      struct {
				Item []struct {
					BasDt      string `json:"basDt"`
					SrtnCd     string `json:"srtnCd"`
					IsinCd     string `json:"isinCd"`
					ItmsNm     string `json:"itmsNm"`
					MrktCtg    string `json:"mrktCtg"`
					Clpr       string `json:"clpr"`
					Vs         string `json:"vs"`
					FltRt      string `json:"fltRt"`
					Mkp        string `json:"mkp"`
					Hipr       string `json:"hipr"`
					Lopr       string `json:"lopr"`
					Trqu       string `json:"trqu"`
					TrPrc      string `json:"trPrc"`
					LstgStCnt  string `json:"lstgStCnt"`
					MrktTotAmt string `json:"mrktTotAmt"`
				} `json:"item"`
			} `json:"items"`
		} `json:"body"`
	} `json:"response"`
}
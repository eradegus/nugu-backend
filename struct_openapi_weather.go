package main

type OpenAPIResponse_weather struct {
	Response struct {
		Header struct {
			ResultCode string `json:"resultCode"`
			ResultMsg  string `json:"resultMsg"`
		} `json:"header"`
		Body struct {
			DataType string `json:"dataType"`
			Items    struct {
				Item []struct {
					BaseDate  string `json:"baseDate"`
					BaseTime  string `json:"baseTime"`
					Category  string `json:"category"`
					FcstDate  string `json:"fcstDate"`
					FcstTime  string `json:"fcstTime"`
					FcstValue string `json:"fcstValue"`
					Nx        int    `json:"nx"`
					Ny        int    `json:"ny"`
				} `json:"item"`
			} `json:"items"`
			PageNo     int `json:"pageNo"`
			NumOfRows  int `json:"numOfRows"`
			TotalCount int `json:"totalCount"`
		} `json:"body"`
	} `json:"response"`
}
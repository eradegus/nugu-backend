package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"time"
	"crypto/tls"
)

func GetStockPriceByStockName(input string) string {

	companyName := input		// ex. SK텔레콤

	// Time Calculation
	nowFullTime := time.Now()
	pastFullTime := nowFullTime.AddDate(0, 0, -1) /////////시간에 따라 조정하기
	pastDate := pastFullTime.Format("20060102")

	// Generate HTTP GET request
	api := "https://apis.data.go.kr/1160100/service/GetStockSecuritiesInfoService/getStockPriceInfo?"
	params := url.Values{}
	params.Add("serviceKey", serviceKey)
	params.Add("pageNo", "1")
	params.Add("numOfRows", "1")
	params.Add("resultType", "json")
	params.Add("basDt", pastDate)
	params.Add("itmsNm", companyName)
	fmt.Println(api + params.Encode())

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(api + params.Encode())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonData := OpenAPIResponse_stock{}
	json.Unmarshal(resData, &jsonData)
	resultPrice := jsonData.Response.Body.Items.Item[0].Clpr

	return resultPrice
}

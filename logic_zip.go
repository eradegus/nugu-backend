package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/xml"
	"bufio"
    "encoding/csv"
    "os"
	"time"
	"strconv"
	"strings"
)

func GetLocationInfoByAptName(input string) string {

	targetApt := input		// ex. "반포자이"

	// Generate HTTP GET request
	api := "http://apis.data.go.kr/1613000/AptListService2/getTotalAptList?"
	params := url.Values{}
	params.Add("serviceKey", serviceKey)
	params.Add("pageNo", "1")
	params.Add("numOfRows", "20000")
	fmt.Println(api + params.Encode())

	resp, err := http.Get(api + params.Encode())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Unmarshal xml string
	xmlData := Response_GetLocationInfoByAptName{}
	xml.Unmarshal(resData, &xmlData)

	itemLists := xmlData.Body.Items.Item
	var targetLocation string
	for _, item := range itemLists {
		if item.KaptName == targetApt {
			targetLocation = item.As2
			break
		}
	}
	return targetLocation
}

func GetCodeByZipLoca(input string) string {

	targetLoca := input

	// file open
    file, _ := os.Open("static/files/NationalRegionCodeutf8.csv")

    // csv reader generation
    rdr := csv.NewReader(bufio.NewReader(file))

    // csv read all
    rows, _ := rdr.ReadAll()

	var targetCode string

	for i, _ := range rows {
        if rows[i][3] == targetLoca {
			targetCode = rows[i][1]
			break
		}
    }
	targetCode = targetCode[:5]
	return targetCode
}

func GetZipInfoByCode(targetCode string, targetApt string) string {

	// Time Calculation
	nowFullTime := time.Now()
	pastFullTime := nowFullTime.AddDate(0, 0, -1)
	pastDate := pastFullTime.Format("20060102")
	pastMonth := pastDate[:6]
	pastDay := pastDate[6:]

	// Generate HTTP GET request
	api := "http://openapi.molit.go.kr/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcAptTradeDev?"
	params := url.Values{}
	params.Add("serviceKey", serviceKey)
	params.Add("pageNo", "1")
	params.Add("numOfRows", "100")
	params.Add("LAWD_CD", targetCode)
	params.Add("DEAL_YMD", pastMonth)

	fmt.Println(api + params.Encode())

	resp, err := http.Get(api + params.Encode())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Unmarshal json string
	xmlData := Response_Zip{}
	xml.Unmarshal(resData, &xmlData)

	itemLists := xmlData.Body.Items.Item

	resultZipInfo := ""
	for _, item := range itemLists {
		if item.Apt == targetApt && item.Day == pastDay {
			resultZipInfo += item.Deal
			break;
		}
	}

	resultZipInfo = strings.Replace(resultZipInfo, ",", "", -1)
	resultZipInfo = strings.Replace(resultZipInfo, " ", "", -1)

	var resultZipInfoNum int
	if resultZipInfo != "" {
		resultZipInfoNum, _ = strconv.Atoi(resultZipInfo)
		resultZipInfoNum = resultZipInfoNum*10000
		resultZipInfo = strconv.Itoa(resultZipInfoNum)
	}
	return resultZipInfo
}
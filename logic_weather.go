package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"time"
	"bufio"
	"encoding/csv"
	"os"
)

func getWeatherLoca (input string) (string, string) {

	targetLoca := input

	// file open
	file, _ := os.Open("static/files/NationalRegionCodeutf8.csv")

	// csv reader generation
	rdr := csv.NewReader(bufio.NewReader(file))

	// csv read all
	rows, _ := rdr.ReadAll()
	var targetX, targetY string

	// rows, row read
	for i, _ := range rows {
		if rows[i][3] == targetLoca {
			targetX = rows[i][5]
			targetY = rows[i][6]
			fmt.Println("targetX: " + targetX)
			fmt.Println("targetY: " + targetY)
			break
		}
	}
	return targetX, targetY
}

func GetWeatherInfoByTownName(input string) (weatherDesc string, temparature string, nowRain string, futureRain string) {

	curLoca := input		// ex. 관악구
	curX, curY := getWeatherLoca(curLoca)

	// Time Calculation
	nowFullTime := time.Now()
	curDate := nowFullTime.Format("20060102")
	curTime := nowFullTime.Format("15:04")
	curTime = curTime[:2] + curTime[3:]

	var nowTimeNum int
	nowTimeNum, _ = strconv.Atoi(curTime)
	nowTimeNum = nowTimeNum - (nowTimeNum%100)
	curTime = strconv.Itoa(nowTimeNum)

	// Generate HTTP GET request
	api := "http://apis.data.go.kr/1360000/VilageFcstInfoService_2.0/getVilageFcst?"
	params := url.Values{}
	params.Add("serviceKey", serviceKey)
	params.Add("pageNo", "1")
	params.Add("numOfRows", "1000")
	params.Add("dataType", "JSON")
	params.Add("base_date", curDate)
	params.Add("base_time", "0500")
	params.Add("nx", curX)
	params.Add("ny", curY)
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

	jsonData := OpenAPIResponse_weather{}
	json.Unmarshal(resData, &jsonData)

	itemList := jsonData.Response.Body.Items.Item

	for _, item := range itemList {
		if item.FcstDate == curDate && item.FcstTime == curTime {
			if item.Category == "TMP" {
				temparature = item.FcstValue
				continue
			}
			if item.Category == "SKY" {
				switch item.FcstValue{
				case "1" :
					weatherDesc = "맑고"
				case "3" :
					weatherDesc = "구름이 많고"
				case "4" :
					weatherDesc = "흐리고"
				}
				continue
			}
			if item.Category == "PTY" {
				switch item.FcstValue{
				case "1" :
					nowRain = "비가 오고 있으며"
				case "2" :
					nowRain = "눈과 비가 같이 오고 있으며"
				case "3" :
					nowRain = "눈이 오고 있으며"
				case "4" :
					nowRain = "소나기가 오고 있으며"
				case "0" :
					nowRain = " "
				}
			}
		}
	}
	var num1, num2 int
	num1, _ = strconv.Atoi(curTime)
	for _, item := range itemList {
		if item.Category == "PTY" && item.FcstDate == curDate {

			num2, _ = strconv.Atoi(item.FcstTime)

			if num2 > num1 {
				switch item.FcstValue{
				case "1" :
					futureRain = "비가 올 예정입니다."
				case "2" :
					futureRain = "눈과 비가 같이 올 예정입니다."
				case "3" :
					futureRain = "눈이 올 예정입니다."
				case "4" :
					futureRain = "소나기가 올 예정입니다."
				case "0" :
					futureRain = " "
				}
			}
		}
	}

	return weatherDesc, temparature, nowRain, futureRain
}


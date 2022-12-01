package main

import (
	"encoding/xml"
)

type Response_GetLocationInfoByAptName struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		ResultCode string `xml:"resultCode"`
		ResultMsg  string `xml:"resultMsg"`
	} `xml:"header"`
	Body struct {
		Text  string `xml:",chardata"`
		Items struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text     string `xml:",chardata"`
				As1      string `xml:"as1"`
				As2      string `xml:"as2"`
				As3      string `xml:"as3"`
				As4      string `xml:"as4"`
				BjdCode  string `xml:"bjdCode"`
				KaptCode string `xml:"kaptCode"`
				KaptName string `xml:"kaptName"`
			} `xml:"item"`
		} `xml:"items"`
		NumOfRows  string `xml:"numOfRows"`
		PageNo     string `xml:"pageNo"`
		TotalCount string `xml:"totalCount"`
	} `xml:"body"`
}

type Response_Zip struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		ResultCode string `xml:"resultCode"`
		ResultMsg  string `xml:"resultMsg"`
	} `xml:"header"`
	Body struct {
		Text  string `xml:",chardata"`
		Items struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text       string `xml:",chardata"`
				Deal       string `xml:"거래금액"`
				거래유형       string `xml:"거래유형"`
				건축년도       string `xml:"건축년도"`
				년          string `xml:"년"`
				도로명        string `xml:"도로명"`
				도로명건물본번호코드 string `xml:"도로명건물본번호코드"`
				도로명건물부번호코드 string `xml:"도로명건물부번호코드"`
				도로명시군구코드   string `xml:"도로명시군구코드"`
				도로명일련번호코드  string `xml:"도로명일련번호코드"`
				도로명지상지하코드  string `xml:"도로명지상지하코드"`
				도로명코드      string `xml:"도로명코드"`
				법정동        string `xml:"법정동"`
				법정동본번코드    string `xml:"법정동본번코드"`
				법정동부번코드    string `xml:"법정동부번코드"`
				법정동시군구코드   string `xml:"법정동시군구코드"`
				법정동읍면동코드   string `xml:"법정동읍면동코드"`
				법정동지번코드    string `xml:"법정동지번코드"`
				Apt        string `xml:"아파트"`
				Month          string `xml:"월"`
				Day          string `xml:"일"`
				일련번호       string `xml:"일련번호"`
				전용면적       string `xml:"전용면적"`
				중개사소재지     string `xml:"중개사소재지"`
				지번         string `xml:"지번"`
				지역코드       string `xml:"지역코드"`
				층          string `xml:"층"`
				해제사유발생일    string `xml:"해제사유발생일"`
				해제여부       string `xml:"해제여부"`
			} `xml:"item"`
		} `xml:"items"`
		NumOfRows  string `xml:"numOfRows"`
		PageNo     string `xml:"pageNo"`
		TotalCount string `xml:"totalCount"`
	} `xml:"body"`
}
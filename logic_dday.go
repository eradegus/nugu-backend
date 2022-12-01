package main

import (
	"strconv"
	"strings"
	"time"
)

func GetDDayInfoByDate(input string) string {

	slice := strings.Split(input, "-")		// input ex. 2022-01-01
	yy, _ := strconv.Atoi(slice[0])
	mm, _ := strconv.Atoi(slice[1])
	dd, _ := strconv.Atoi(slice[2])
	dday := Date(yy, mm, dd)

	now := time.Now()
	today := Date(now.Year(), int(now.Month()), now.Day())

	if !today.After(dday) {
		return ""
	}

	elapsedDays := today.Sub(dday).Hours() / 24
	return CheckAnniversary(int(elapsedDays))
}

func Date(year, month, day int) time.Time {
	location, _ := time.LoadLocation("Asia/Seoul")
	timeInUTC := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return timeInUTC.In(location)
}

func CheckAnniversary(days int) string {
	if days <= 1000 {
		modHundred := days % 100
		if modHundred == 0 {
			return strconv.Itoa(days / 100) + "00일이에요 축하해요."
		}

		gapHundred := 100 - modHundred
		if gapHundred <= 7 {
			return strconv.Itoa(days / 100 + 1) + "00일까지 " + strconv.Itoa(gapHundred) + "일 남았어요."
		}
	}

	modYear := days % 365
	if modYear == 0 {
		return strconv.Itoa(days / 365) + "주년이에요 축하해요."
	}

	gapYear := 365 - modYear
	if gapYear <= 7 {
		return strconv.Itoa(days / 365 + 1) + "주년까지 " + strconv.Itoa(gapYear) + "일 남았어요."
	}

	return ""
}

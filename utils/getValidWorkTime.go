package utils

import (
	"fmt"
	"math"
	"time"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func IsHoliday(ti time.Time) bool {
	festivalDayList := []string{}
	weekendDayWorkList := []string{}
	// 2022 statutory holiday From gov like http://www.gov.cn/zhengce/content/2021-10/25/content_5644835.htm  or third-party like http://timor.tech/api/holiday/year
	festivalDay2022 := []string{"2022-10-01", "2022-10-03", "2022-05-04", "2022-05-03", "2022-05-02", "2022-05-01", "2022-10-05", "2022-10-04", "2022-04-30", "2022-10-02", "2022-10-06", "2022-09-10", "2022-09-11", "2022-09-12", "2022-06-03", "2022-06-04", "2022-06-05", "2022-01-31", "2022-02-06", "2022-02-04", "2022-02-05", "2022-02-02", "2022-02-03", "2022-10-07", "2022-02-01", "2022-04-04", "2022-04-05", "2022-04-03", "2022-01-03", "2022-01-02", "2022-01-01"}
	weekendDayWork2022 := []string{"2022-05-07", "2022-01-30", "2022-10-09", "2022-10-08", "2022-04-24", "2022-04-02", "2022-01-29"}
	// All holiday sum
	festivalDayList = append(festivalDayList, festivalDay2022...)
	weekendDayWorkList = append(weekendDayWorkList, weekendDayWork2022...)

	// get the day of week number
	dayOfWeekInt := int(ti.Weekday())
	yearMonthDayStr := string(ti.Format("2006-01-02"))

	// check if yearMonthDayStr in festivalDayList,must be holiday;if yearMonthDayStr in weekendDayWorkList,must be workDay;
	// except these, check the day of week to decide whether it is holiday
	if contains(festivalDayList, yearMonthDayStr) {
		//fmt.Println("festivalDay warning:", yearMonthDayStr)
		return true
	} else if contains(weekendDayWorkList, yearMonthDayStr) {
		//fmt.Println("weekendDay(work for festivalDay) warning:", yearMonthDayStr)
		return false
	} else if dayOfWeekInt == 6 || dayOfWeekInt == 0 {
		//fmt.Println("normal weekend warning:", yearMonthDayStr)
		// sunday corresponding value is 0, not 7
		return true
	}

	//fmt.Println("normal work day:", yearMonthDayStr)
	return false
}

func GetValidWorkingTimeDuration(startTime time.Time, endTime time.Time) (validRepairedMinutes int, holidayList []string, err error) {
	if startTime.IsZero() || endTime.IsZero() {
		return 0, nil, fmt.Errorf("startTime or endTime is nil,check it")
	}
	if startTime.After(endTime) {
		return 0, nil, fmt.Errorf("startTime cann't after endTime,check it")
	}

	// Duration between startTime and endTime Format: [ PrefixDurationTo24 ] + [ Middle1 + ... + MiddleN ] + [ SuffixDurationFrom24 ]
	// MiddleDays could be null, or could be Mix with holidays and workdays; PrefixDurationTo24/SuffixDurationFrom24 alse maybe holiday or workday
	cursorTime := startTime
	for {
		if math.Floor(endTime.Sub(startTime).Hours()) < 24 {
			if !IsHoliday(startTime) {
				validRepairedMinutes = (int(endTime.Unix() - startTime.Unix())) / 60
			} else {
				yearMonthDayStr := string(startTime.Format("2006-01-02"))
				holidayList = append(holidayList, yearMonthDayStr)
			}
			break
		}

		cursorTime = cursorTime.AddDate(0, 0, 1)
		y1, m1, d1 := cursorTime.Date()
		y2, m2, d2 := endTime.Date()
		if y1 == y2 && m1 == m2 && d1 == d2 {
			if !IsHoliday(startTime) {
				firstDayWorkTimeMinutes := (24-startTime.Hour()-1)*60 + (60 - startTime.Minute())
				validRepairedMinutes += firstDayWorkTimeMinutes
			} else {
				yearMonthDayStr := string(startTime.Format("2006-01-02"))
				holidayList = append(holidayList, yearMonthDayStr)
			}
			if !IsHoliday(endTime) {
				// workTimeMinutes := (endTime.Hour()*60 + endTime.Minute()) * 60
				lastDayWorkTimeMinutes := endTime.Hour()*60 + endTime.Minute()
				validRepairedMinutes += lastDayWorkTimeMinutes
			} else {
				yearMonthDayStr := string(endTime.Format("2006-01-02"))
				holidayList = append(holidayList, yearMonthDayStr)
			}
			break
		} else {
			if !IsHoliday(cursorTime) {
				middleDayWorkTimeMinutes := 24 * 60
				validRepairedMinutes += middleDayWorkTimeMinutes
			} else {
				yearMonthDayStr := string(cursorTime.Format("2006-01-02"))
				holidayList = append(holidayList, yearMonthDayStr)
				continue
			}
		}
	}
	return validRepairedMinutes, holidayList, nil
}


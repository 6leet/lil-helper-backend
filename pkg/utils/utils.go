package utils

import (
	"fmt"
	"lil-helper-backend/config"
	"strconv"
	"time"
)

func ParseTime(timestr string) time.Time {
	year, _ := strconv.Atoi(timestr[0:4])
	fmt.Println(year)
	month, _ := strconv.Atoi(timestr[5:7])
	fmt.Println(month)
	day, _ := strconv.Atoi(timestr[8:10])
	fmt.Println(day)
	hour := config.Config.Mission.Updateat
	fmt.Println(hour)
	return time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.UTC).Add(time.Hour * -8)
}

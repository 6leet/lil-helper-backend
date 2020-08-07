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
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Add(time.Hour * -8)
}

func ParseTimeLocation(t *time.Time) {
	location, _ := time.LoadLocation("Local")
	*t = t.In(location)
}

func ExpToLevel(exp int) uint {
	maxLevel := config.Config.Mission.Maxlevel
	levelThres := config.Config.Mission.Levelsexp

	for l := 0; l < maxLevel-1; l++ {
		if levelThres[l] <= exp && exp < levelThres[l+1] {
			return uint(l)
		}
	}
	return 0
}

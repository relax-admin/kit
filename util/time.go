package util

import "time"

func ZoneChina(timeParam time.Time) (newZone time.Time) {
	return Zone(timeParam, "Asia/Shanghai")
}

func Zone(timeParam time.Time, location string) (newZone time.Time) {
	loc, _ := time.LoadLocation(location)
	return timeParam.UTC().In(loc)
}

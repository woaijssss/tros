package utils

import (
	"fmt"
	"math"
)

const EarthRadius = 6371.0 // 地球半径（km）

func FormatLocation2String(longitude, latitude float64) string {
	return fmt.Sprintf("%f,%f", longitude, latitude)
}

type DistanceOption struct {
	Longitude float64
	Latitude  float64
}

func Distance(poi1, poi2 *DistanceOption) float64 {
	lat1 := poi1.Latitude * math.Pi / 180.0
	lng1 := poi1.Longitude * math.Pi / 180.0
	lat2 := poi2.Latitude * math.Pi / 180.0
	lng2 := poi2.Longitude * math.Pi / 180.0

	d := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(lng2-lng1))
	return math.Round(EarthRadius*d*100) / 100 // 四舍五入
}

package utils

import "fmt"

func FormatLocation2String(longitude, latitude float64) string {
	return fmt.Sprintf("%f,%f", longitude, latitude)
}

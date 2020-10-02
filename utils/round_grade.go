package utils

import "math"

// GradeRounder round to nearest value
func GradeRounder(val float64) float64 {
	return math.Round(val*10) / 10
}

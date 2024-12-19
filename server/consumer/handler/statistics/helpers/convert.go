package helpers

import (
	"fmt"
	"strconv"
)

func DivisionFloat(n1 float64, division int) float64{
	p := fmt.Sprintf("%.2f", (n1/float64(division)))
	f32, _ := strconv.ParseFloat(p, 64)
	return float64(f32)
}

func SubtractionFloat(n1 float64, n2 float64) float64 {
	p := fmt.Sprintf("%.2f", (n1 - n2))
	f32, _ := strconv.ParseFloat(p, 64)
	return float64(f32)
}
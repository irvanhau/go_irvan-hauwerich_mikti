package models

import (
	"fmt"
	"math"
	"strconv"
)

type Identity struct {
	Berat  float64
	Tinggi float64
}

type BMIInterface interface {
	HitungBMI() float64
}

func (i *Identity) HitungBMI() float64 {
	powTinggi := math.Pow(i.Tinggi, 2)
	bmi := i.Berat / powTinggi

	bmiString := fmt.Sprintf("%.2f", bmi)
	bmiFloat, _ := strconv.ParseFloat(bmiString, 64)

	return bmiFloat
}

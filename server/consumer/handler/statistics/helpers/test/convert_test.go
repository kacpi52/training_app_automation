package test

import (
	helpers_statictics "myInternal/consumer/handler/statistics/helpers"
	"testing"
)

func TestDivisionFloat(t *testing.T){
	number1 := 77.77
	number2 := 6

	divisionFloat := helpers_statictics.DivisionFloat(number1, number2)
	if divisionFloat != 12.96{
		t.Errorf("error DivisionFloat return is not 12.96")
	}

}

func TestSubtractionFloat(t *testing.T){
	number1 := 77.77
	number2 := 66.66

	divisionFloat := helpers_statictics.SubtractionFloat(number1, number2)
	if divisionFloat != 11.11{
		t.Errorf("error SubtractionFloat return is not 11.11")
	}
}
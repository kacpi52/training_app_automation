package helpers

import (
	"fmt"
	statistics_data "myInternal/consumer/data/statistics"
	"os"
)

func CreateFileStatistic(data []statistics_data.Create) error{

	file, err := os.Create("dane.txt")
	if err != nil{
		return fmt.Errorf("file not create %v", err)
	}
	defer file.Close()

	for _, v := range data{

		var down string = fmt.Sprintf("Waga w doł: %.2f kg \n", v.DownWeight)
		if v.DownWeight > 0{
			down = fmt.Sprintf("Waga w doł: +%.2f kg \n", v.DownWeight)
		}

		line := fmt.Sprintf("Tydzień: %d \n", v.Week) +
		fmt.Sprintf("Startowa Waga: %.2f \n", v.StartWeight) +
		fmt.Sprintf("Waga końcowa: %.2f \n", v.EndWeight) +
		down +
		fmt.Sprintf("Suma kilogramów: %.2f \n", v.SumKg) +
		fmt.Sprintf("Średnia waga: %.2f \n", v.AvgKg) +
		fmt.Sprintf("Suma fitatu: %.2f \n", v.SumKcal)

		for _, t := range v.Training{
			for _, k := range t.Data{
				line += fmt.Sprintf("Typ treningu: %v, liczba treningów: %d, suma spalonych kalorii: %d \n", k.Type, k.Currecnt, k.SumKcal)
			}

		}

		line += fmt.Sprintf("Suma czasu treningów: %v \n\n", v.Training)

		_, err := file.WriteString(line)
		if err != nil{
			return fmt.Errorf("not wirte to file %v", err)
		}
	}

	return nil
}

package main

import (
	"fmt"
	"tugasgo/challenge2/soal1/models"
)

func main() {
	lumba := models.Score{}
	koala := models.Score{}

	fmt.Print("Masukan 3 Score Tim Lumba-Lumba : ")
	fmt.Scanln(&lumba.Score1, &lumba.Score2, &lumba.Score3)

	fmt.Print("Masukan 3 Score Tim Koala : ")
	fmt.Scanln(&koala.Score1, &koala.Score2, &koala.Score3)

	lumbaScore := lumba.CalculateScore()
	koalaScore := koala.CalculateScore()

	fmt.Println("Score Lumba-Lumba : ", lumbaScore)
	fmt.Println("Score Koala : ", koalaScore)
	fmt.Println(models.DecideTheWinner(lumbaScore, koalaScore))
}

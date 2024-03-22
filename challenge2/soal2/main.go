package main

import (
	"fmt"
	"tugasgo/challenge2/soal2/models"
)

func main() {
	var markHigherBMI bool
	mark := models.Identity{}
	john := models.Identity{}

	fmt.Print("Tulis Berat dan Tinggi Mark : ")
	fmt.Scanln(&mark.Berat, &mark.Tinggi)

	fmt.Print("Tulis Berat dan Tinggi John : ")
	fmt.Scanln(&john.Berat, &john.Tinggi)

	markBMI := mark.HitungBMI()
	johnBMI := john.HitungBMI()

	if markBMI > johnBMI {
		markHigherBMI = true
	}

	fmt.Println("BMI Mark : ", markBMI)
	fmt.Println("BMI John : ", johnBMI)
	fmt.Println("Apakah Mark Memiliki BMI lebih tinggi dari John : ", markHigherBMI)
}

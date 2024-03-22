package models

import (
	"errors"
	"fmt"
	"strconv"
)

type Score struct {
	Score1, Score2, Score3 float64
}

type ScoreInterface interface {
	CalculateScore() float64
}

func (s *Score) CalculateScore() float64 {
	allScore := s.Score1 + s.Score2 + s.Score3

	calculate := allScore / 3

	avgString := fmt.Sprintf("%.2f", calculate)
	avgFloat, _ := strconv.ParseFloat(avgString, 64)

	return avgFloat
}

func DecideTheWinner(lumba, koala float64) string {
	if lumba > koala {
		if err := checkScore(lumba); err != nil {
			return err.Error()
		}

		return "Lumba-Lumba adalah Pemenang nya"
	} else if koala > lumba {
		if err := checkScore(koala); err != nil {
			return err.Error()
		}

		return "Koala adalah Pemenang nya"
	} else if lumba == koala {
		if err := checkDraw(lumba, koala); err != nil {
			return err.Error()
		}

		return "Seri"
	}

	return ""
}

func checkScore(score float64) error {
	if score < 100 {
		return errors.New("Tidak Mencapai Skor Minimum")
	}

	return nil
}

func checkDraw(lumba, koala float64) error {
	if lumba < 100 && koala < 100 && lumba != koala {
		return errors.New("Tidak ada Tim yang memenangkan Trofi")
	}

	return nil
}

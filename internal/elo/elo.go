package elo

import "math"

// Probability calculates the chances for rating1 to beat rating2
func Probability(rating1 float64, rating2 float64) float64 {
	exponent := (rating2 - rating1) / 400.0
	probability := 1.0 / (1.0 + math.Pow(10, exponent))

	return math.Round(probability*100) / 100
}

// Ranking calculates the new rankings of player 1 and player 2. If winner is
// true, then player 1 won, else player 2 won. K is the constant, normally 30.
func Ranking(rating1 float64, rating2 float64, winner bool, K int) (float64, float64) {
	p1 := Probability(rating1, rating2)
	p2 := Probability(rating2, rating1)
	var r1 float64
	var r2 float64

	if winner {
		r1 = rating1 + float64(K)*(1.0-p1)
		r2 = rating2 + float64(K)*(0.0-p2)
	} else {
		r1 = rating1 + float64(K)*(0.0-p1)
		r2 = rating2 + float64(K)*(1.0-p2)
	}

	return r1, r2
}

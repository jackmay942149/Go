package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Min(xs []float64) (min float64) {
	for i, v := range xs {
		if (i == 0) || (v < min) {
			min = v
		}
	}
	return
}

// Calculates the largest value in slice xs
func Max(xs []float64) (max float64) {
	for i, v := range xs {
		if (i == 0) || (v > max) {
			max = v
		}
	}
	return
}

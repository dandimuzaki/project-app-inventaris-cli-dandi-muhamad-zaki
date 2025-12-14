package utils

func CalculateNetValue(price float64, year int) float64 {
	for i := 1; i <= year; i++ {
		price *= 0.8
	}
	return price
}
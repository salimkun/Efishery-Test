package util

import "sort"

func GetMedian(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)

	sort.Float64s(dataCopy)

	var median float64
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}

	return median
}

func GetAvg(data []float64) float64 {
	var total float64
	for _, i := range data {
		total = total + i
	}
	total = total / float64(len(data))
	return total
}

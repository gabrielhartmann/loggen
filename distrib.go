package generator

import (
	"fmt"
	rand "math/rand"
)

// Mean and standard deviation should be denominated in seconds
func GetNormalDistribution(mean float64, stdDev float64, count int) []float64 {
	list := []float64{}

	for i := 0; i < count; i++ {
		sample := rand.NormFloat64()*stdDev + mean
		list = append(list, sample)
	}

	return list
}

func GetBucketCount(input []float64, bucketStart float64, bucketEnd float64) int {
	count := 0
	for _, e := range input {
		if bucketStart <= e && e < bucketEnd {
			count++
		}
	}

	return count
}

// print +/- 3 standard deviations
func PrintDistribution(input []float64, mean float64, stdDev float64, bucketSize float64) {
	stdDevCount := 3.0
	for curr := mean - stdDev*stdDevCount; curr < mean+stdDev*stdDevCount; curr += bucketSize {
		bucketStart := curr
		bucketEnd := curr + bucketSize
		count := GetBucketCount(input, bucketStart, bucketEnd)
		fmt.Printf("[%v, %v] = %v\n", bucketStart, bucketEnd, count)
	}
}

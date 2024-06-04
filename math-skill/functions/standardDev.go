package functions

import (
	"math"
)

func StandardDev(variance []float64) float64 {
    if len(variance) == 0 {
        return 0
    }
    sum := 0.0
    for _, v := range variance {
        sum += v
    }
    mean := sum / float64(len(variance))
    sumSquaredDiffs := 0.0
    for _, v := range variance {
        diff := v - mean
        sumSquaredDiffs += diff * diff
    }
    return math.Sqrt(sumSquaredDiffs / float64(len(variance)))
}

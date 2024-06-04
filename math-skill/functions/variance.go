package functions

func Variance(data []float64, average float64) float64 {
  var variance float64
  for _, value := range data {
    variance += (value - average) * (value - average)
  }
  return variance / float64(len(data))
}

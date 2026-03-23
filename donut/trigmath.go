package donut

import "math"

var (
	sinTable2 []float64
	cosTable2 []float64
	sinTable4 []float64
	cosTable4 []float64
	sinTable7 []float64
	cosTable7 []float64
)

func init() {
	sinTable2 = make([]float64, 314)
	cosTable2 = make([]float64, 314)
	sinTable4 = make([]float64, 157)
	cosTable4 = make([]float64, 157)
	sinTable7 = make([]float64, 89)
	cosTable7 = make([]float64, 89)

	for i := range sinTable2 {
		sinTable2[i] = math.Sin(float64(i) * 0.02)
		cosTable2[i] = math.Cos(float64(i) * 0.02)
	}

	for i := range sinTable4 {
		sinTable4[i] = math.Sin(float64(i) * 0.04)
		cosTable4[i] = math.Cos(float64(i) * 0.04)
	}

	for i := range sinTable7 {
		sinTable7[i] = math.Sin(float64(i) * 0.07)
		cosTable7[i] = math.Cos(float64(i) * 0.07)
	}
}

package np

type Float64OneDArray struct {
	arr  []float64
	step *float64
	err  error
}

type Float64TwoDArray struct {
	arr [][]float64
	err error
}

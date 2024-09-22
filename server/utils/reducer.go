package utils

func Float64Reducer(arr []float64, op func(float64, float64) float64) float64 {
	var result = arr[0]
	for _, v := range arr[1:] {
		result = op(result, v)
	}
	return result
}

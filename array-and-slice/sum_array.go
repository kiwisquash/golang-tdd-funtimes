package array_and_slice

func SumArray(arr []int) int {
	sum := 0
	for _, s := range arr {
		sum +=s
	}
	return sum
}
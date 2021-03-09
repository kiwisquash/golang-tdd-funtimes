package array_and_slice

func SumArray(arr []int) int {
	sum := 0
	for _, s := range arr {
		sum +=s
	}
	return sum
}

func SumAll(collection ...[]int) []int {
	var output []int
	for _, c := range collection {
		output = append(output, SumArray(c))
	}
	return output
}

func SumTail(collection ...[]int) []int {
	var tailSums []int
	for _, c := range collection {
		if len(c) > 0 {
			tailSums = append(tailSums, SumArray(c[1:]))
		} else {
			tailSums = append(tailSums, 0)
		}
	}
	return tailSums
}
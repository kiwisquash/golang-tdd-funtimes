package iteration

// Repeat repeats string n times
func Repeat(r string, n int) string {

	var output string

	for i:=0 ; i < n; i ++ {
		output += r
	}

	return output
}
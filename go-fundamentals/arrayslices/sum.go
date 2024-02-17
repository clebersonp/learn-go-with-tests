package arrayslices

// Sum - range lets you iterate over an array.
// On each iteration, range returns two values - the index and the value.
// We are choosing to ignore the index value by using _ blank identifier
// Array size is encoded in its type. If you try to pass an [4]int into a function that expects [5]int, it won't compile.
func Sum(numbers [5]int) (result int) {
	for _, num := range numbers {
		result += num
	}
	return
}

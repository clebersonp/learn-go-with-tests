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

// SumSlice - we don't user numbers or '...' operator when we declare a slice, we give an empty size
func SumSlice(numbers []int) (result int) {
	for _, num := range numbers {
		result += num
	}
	return
}

// SumAll - Go can let you write 'variadic function' that can take a variable number of arguments.
func SumAll(slicesToSum ...[]int) []int {
	// make allows you to create a slice with a starting capacity of the len
	// of the slicesToSum we need to work through
	numOfSlices := len(slicesToSum)
	sums := make([]int, numOfSlices)
	for i, slice := range slicesToSum {
		sums[i] = SumSlice(slice)
	}
	return sums
}

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

// SumAllWithAppend - Go can let you write 'variadic function' that can take a variable number of arguments.
func SumAllWithAppend(slicesToSum ...[]int) []int {
	var sums []int
	for _, slice := range slicesToSum {
		// In this implementation, we are worrying less about capacity.
		// We start with an empty slice sums and append to it the result of Sum as we work through the varargs.
		sums = append(sums, SumSlice(slice))
	}
	return sums
}

func SumAllTails(slicesToSumTails ...[]int) []int {
	var sumTails []int
	for _, slice := range slicesToSumTails {
		if len(slice) == 0 {
			sumTails = append(sumTails, 0)
		} else {
			// slicing a slice
			tails := slice[1:]
			sumTails = append(sumTails, SumSlice(tails))
		}
	}
	return sumTails
}

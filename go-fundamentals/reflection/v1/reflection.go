package v1

func walk(x interface{}, fn func(input string)) {
	fn("some words here!")
}

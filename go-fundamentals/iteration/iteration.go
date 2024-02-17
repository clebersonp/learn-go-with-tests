package iteration

func Repeat(character string, times int) string {
	var repeated string
	for times > 0 {
		repeated += character
		times--
	}
	if repeated == "" {
		repeated = character
	}
	return repeated
}

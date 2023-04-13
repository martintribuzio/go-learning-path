package iteration

// Repeat repeats the given character 5 times
func Repeat(character string) string {
	var repeated string
	// Additional variants of the for loop: https://gobyexample.com/for
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}

// RepeatNTimes repeats the given character n times
func RepeatNTimes(character string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}

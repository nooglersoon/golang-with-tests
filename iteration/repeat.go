package iteration

func Repeat(character string, iteration int) string {
	// Define a variable with var following by its name and type
	var repeated string
	if iteration == 0 {
		return character
	} else if iteration < 0 {
		return "The function can't support negative iteration"
	}

	// We don't need the curly braces when using for loop
	for i := 0; i < iteration; i++ {
		repeated += character
	}
	return repeated
}

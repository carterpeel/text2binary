package util


// ToBinary returns the binary representation of n.
func ToBinary(n int) (bin int) {
	// Initialize the counter
	c := 1
	// Initialize the remainder
	r := 0

	// While n is not 0, continue looping
	for n != 0 {
		// Set r to the remainder of dividing n by two
		// using the modulus operator
		r = n%2

		// Add the remainder multiplied by the value of
		// the counter to the binary result
		bin += r * c

		// Divide n by two
		n = n/2

		// Multiply the counter by 10
		c *= 10
	}

	// Return the binary result
	return bin
}
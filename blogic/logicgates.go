package blogic

//And takes 2 bits as input and returns 1 if they are both 1,
//or 0 otherwise
func And(a, b int) int {
	if a == 1 && b == 1 {
		return 1
	}
	return 0
}

//Or takes 2 bits as input and returns 1 as long as the input
//bits are not both 0
func Or(a, b int) int {
	if a == 1 || b == 1 {
		return 1
	}
	return 0
}

//Not reverses a bit
func Not(a int) int {
	if a != 1 {
		return 1
	}
	return 0
}

//Xor is an exclusive or gate, built by combining several
//simpler gates. It returns a 1 if ONLY one of the inputs
//is one, or a zero if they are both 0 or both 1
func Xor(a, b int) int {
	x := Not(And(a, b))
	y := Or(a, b)
	return And(x, y)
}

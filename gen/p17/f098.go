package p17

// F098_0 transforms x for stage 0 of file 98.
func F098_0(x int) int {
	y := F097_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 98
}

// F098_1 transforms x for stage 1 of file 98.
func F098_1(x int) int {
	y := F097_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 98
}

// F098_2 transforms x for stage 2 of file 98.
func F098_2(x int) int {
	y := F097_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 98
}

// F098_3 transforms x for stage 3 of file 98.
func F098_3(x int) int {
	y := F097_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 98
}

// F098_4 transforms x for stage 4 of file 98.
func F098_4(x int) int {
	y := F097_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 98
}

// F098_5 transforms x for stage 5 of file 98.
func F098_5(x int) int {
	y := F097_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 98
}

// F098_6 transforms x for stage 6 of file 98.
func F098_6(x int) int {
	y := F097_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 98
}

// F098_7 transforms x for stage 7 of file 98.
func F098_7(x int) int {
	y := F097_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 98
}


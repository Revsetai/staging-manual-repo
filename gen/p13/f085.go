package p13

// F085_0 transforms x for stage 0 of file 85.
func F085_0(x int) int {
	y := F084_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 85
}

// F085_1 transforms x for stage 1 of file 85.
func F085_1(x int) int {
	y := F084_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 85
}

// F085_2 transforms x for stage 2 of file 85.
func F085_2(x int) int {
	y := F084_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 85
}

// F085_3 transforms x for stage 3 of file 85.
func F085_3(x int) int {
	y := F084_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 85
}

// F085_4 transforms x for stage 4 of file 85.
func F085_4(x int) int {
	y := F084_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 85
}

// F085_5 transforms x for stage 5 of file 85.
func F085_5(x int) int {
	y := F084_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 85
}

// F085_6 transforms x for stage 6 of file 85.
func F085_6(x int) int {
	y := F084_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 85
}

// F085_7 transforms x for stage 7 of file 85.
func F085_7(x int) int {
	y := F084_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 85
}


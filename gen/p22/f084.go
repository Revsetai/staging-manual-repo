package p22

// F084_0 transforms x for stage 0 of file 84.
func F084_0(x int) int {
	y := F083_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 84
}

// F084_1 transforms x for stage 1 of file 84.
func F084_1(x int) int {
	y := F083_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 84
}

// F084_2 transforms x for stage 2 of file 84.
func F084_2(x int) int {
	y := F083_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 84
}

// F084_3 transforms x for stage 3 of file 84.
func F084_3(x int) int {
	y := F083_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 84
}

// F084_4 transforms x for stage 4 of file 84.
func F084_4(x int) int {
	y := F083_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 84
}

// F084_5 transforms x for stage 5 of file 84.
func F084_5(x int) int {
	y := F083_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 84
}

// F084_6 transforms x for stage 6 of file 84.
func F084_6(x int) int {
	y := F083_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 84
}

// F084_7 transforms x for stage 7 of file 84.
func F084_7(x int) int {
	y := F083_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 84
}


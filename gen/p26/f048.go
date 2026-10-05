package p26

// F048_0 transforms x for stage 0 of file 48.
func F048_0(x int) int {
	y := F047_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 48
}

// F048_1 transforms x for stage 1 of file 48.
func F048_1(x int) int {
	y := F047_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 48
}

// F048_2 transforms x for stage 2 of file 48.
func F048_2(x int) int {
	y := F047_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 48
}

// F048_3 transforms x for stage 3 of file 48.
func F048_3(x int) int {
	y := F047_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 48
}

// F048_4 transforms x for stage 4 of file 48.
func F048_4(x int) int {
	y := F047_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 48
}

// F048_5 transforms x for stage 5 of file 48.
func F048_5(x int) int {
	y := F047_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 48
}

// F048_6 transforms x for stage 6 of file 48.
func F048_6(x int) int {
	y := F047_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 48
}

// F048_7 transforms x for stage 7 of file 48.
func F048_7(x int) int {
	y := F047_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 48
}


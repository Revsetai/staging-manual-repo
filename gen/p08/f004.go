package p08

// F004_0 transforms x for stage 0 of file 4.
func F004_0(x int) int {
	y := F003_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 4
}

// F004_1 transforms x for stage 1 of file 4.
func F004_1(x int) int {
	y := F003_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 4
}

// F004_2 transforms x for stage 2 of file 4.
func F004_2(x int) int {
	y := F003_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 4
}

// F004_3 transforms x for stage 3 of file 4.
func F004_3(x int) int {
	y := F003_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 4
}

// F004_4 transforms x for stage 4 of file 4.
func F004_4(x int) int {
	y := F003_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 4
}

// F004_5 transforms x for stage 5 of file 4.
func F004_5(x int) int {
	y := F003_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 4
}

// F004_6 transforms x for stage 6 of file 4.
func F004_6(x int) int {
	y := F003_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 4
}

// F004_7 transforms x for stage 7 of file 4.
func F004_7(x int) int {
	y := F003_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 4
}


package p25

// F024_0 transforms x for stage 0 of file 24.
func F024_0(x int) int {
	y := F023_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 24
}

// F024_1 transforms x for stage 1 of file 24.
func F024_1(x int) int {
	y := F023_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 24
}

// F024_2 transforms x for stage 2 of file 24.
func F024_2(x int) int {
	y := F023_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 24
}

// F024_3 transforms x for stage 3 of file 24.
func F024_3(x int) int {
	y := F023_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 24
}

// F024_4 transforms x for stage 4 of file 24.
func F024_4(x int) int {
	y := F023_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 24
}

// F024_5 transforms x for stage 5 of file 24.
func F024_5(x int) int {
	y := F023_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 24
}

// F024_6 transforms x for stage 6 of file 24.
func F024_6(x int) int {
	y := F023_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 24
}

// F024_7 transforms x for stage 7 of file 24.
func F024_7(x int) int {
	y := F023_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 24
}


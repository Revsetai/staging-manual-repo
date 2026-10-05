package p04

// F019_0 transforms x for stage 0 of file 19.
func F019_0(x int) int {
	y := F018_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 19
}

// F019_1 transforms x for stage 1 of file 19.
func F019_1(x int) int {
	y := F018_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 19
}

// F019_2 transforms x for stage 2 of file 19.
func F019_2(x int) int {
	y := F018_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 19
}

// F019_3 transforms x for stage 3 of file 19.
func F019_3(x int) int {
	y := F018_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 19
}

// F019_4 transforms x for stage 4 of file 19.
func F019_4(x int) int {
	y := F018_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 19
}

// F019_5 transforms x for stage 5 of file 19.
func F019_5(x int) int {
	y := F018_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 19
}

// F019_6 transforms x for stage 6 of file 19.
func F019_6(x int) int {
	y := F018_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 19
}

// F019_7 transforms x for stage 7 of file 19.
func F019_7(x int) int {
	y := F018_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 19
}


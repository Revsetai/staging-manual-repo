package p23

// F009_0 transforms x for stage 0 of file 9.
func F009_0(x int) int {
	y := F008_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 9
}

// F009_1 transforms x for stage 1 of file 9.
func F009_1(x int) int {
	y := F008_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 9
}

// F009_2 transforms x for stage 2 of file 9.
func F009_2(x int) int {
	y := F008_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 9
}

// F009_3 transforms x for stage 3 of file 9.
func F009_3(x int) int {
	y := F008_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 9
}

// F009_4 transforms x for stage 4 of file 9.
func F009_4(x int) int {
	y := F008_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 9
}

// F009_5 transforms x for stage 5 of file 9.
func F009_5(x int) int {
	y := F008_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 9
}

// F009_6 transforms x for stage 6 of file 9.
func F009_6(x int) int {
	y := F008_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 9
}

// F009_7 transforms x for stage 7 of file 9.
func F009_7(x int) int {
	y := F008_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 9
}


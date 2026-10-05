package p03

// F036_0 transforms x for stage 0 of file 36.
func F036_0(x int) int {
	y := F035_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 36
}

// F036_1 transforms x for stage 1 of file 36.
func F036_1(x int) int {
	y := F035_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 36
}

// F036_2 transforms x for stage 2 of file 36.
func F036_2(x int) int {
	y := F035_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 36
}

// F036_3 transforms x for stage 3 of file 36.
func F036_3(x int) int {
	y := F035_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 36
}

// F036_4 transforms x for stage 4 of file 36.
func F036_4(x int) int {
	y := F035_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 36
}

// F036_5 transforms x for stage 5 of file 36.
func F036_5(x int) int {
	y := F035_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 36
}

// F036_6 transforms x for stage 6 of file 36.
func F036_6(x int) int {
	y := F035_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 36
}

// F036_7 transforms x for stage 7 of file 36.
func F036_7(x int) int {
	y := F035_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 36
}


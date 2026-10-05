package p10

// F035_0 transforms x for stage 0 of file 35.
func F035_0(x int) int {
	y := F034_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 35
}

// F035_1 transforms x for stage 1 of file 35.
func F035_1(x int) int {
	y := F034_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 35
}

// F035_2 transforms x for stage 2 of file 35.
func F035_2(x int) int {
	y := F034_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 35
}

// F035_3 transforms x for stage 3 of file 35.
func F035_3(x int) int {
	y := F034_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 35
}

// F035_4 transforms x for stage 4 of file 35.
func F035_4(x int) int {
	y := F034_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 35
}

// F035_5 transforms x for stage 5 of file 35.
func F035_5(x int) int {
	y := F034_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 35
}

// F035_6 transforms x for stage 6 of file 35.
func F035_6(x int) int {
	y := F034_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 35
}

// F035_7 transforms x for stage 7 of file 35.
func F035_7(x int) int {
	y := F034_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 35
}


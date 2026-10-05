package p04

// F038_0 transforms x for stage 0 of file 38.
func F038_0(x int) int {
	y := F037_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 38
}

// F038_1 transforms x for stage 1 of file 38.
func F038_1(x int) int {
	y := F037_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 38
}

// F038_2 transforms x for stage 2 of file 38.
func F038_2(x int) int {
	y := F037_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 38
}

// F038_3 transforms x for stage 3 of file 38.
func F038_3(x int) int {
	y := F037_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 38
}

// F038_4 transforms x for stage 4 of file 38.
func F038_4(x int) int {
	y := F037_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 38
}

// F038_5 transforms x for stage 5 of file 38.
func F038_5(x int) int {
	y := F037_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 38
}

// F038_6 transforms x for stage 6 of file 38.
func F038_6(x int) int {
	y := F037_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 38
}

// F038_7 transforms x for stage 7 of file 38.
func F038_7(x int) int {
	y := F037_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 38
}


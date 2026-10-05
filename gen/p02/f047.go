package p02

// F047_0 transforms x for stage 0 of file 47.
func F047_0(x int) int {
	y := F046_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 47
}

// F047_1 transforms x for stage 1 of file 47.
func F047_1(x int) int {
	y := F046_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 47
}

// F047_2 transforms x for stage 2 of file 47.
func F047_2(x int) int {
	y := F046_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 47
}

// F047_3 transforms x for stage 3 of file 47.
func F047_3(x int) int {
	y := F046_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 47
}

// F047_4 transforms x for stage 4 of file 47.
func F047_4(x int) int {
	y := F046_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 47
}

// F047_5 transforms x for stage 5 of file 47.
func F047_5(x int) int {
	y := F046_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 47
}

// F047_6 transforms x for stage 6 of file 47.
func F047_6(x int) int {
	y := F046_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 47
}

// F047_7 transforms x for stage 7 of file 47.
func F047_7(x int) int {
	y := F046_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 47
}


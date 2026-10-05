package p27

// F045_0 transforms x for stage 0 of file 45.
func F045_0(x int) int {
	y := F044_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 45
}

// F045_1 transforms x for stage 1 of file 45.
func F045_1(x int) int {
	y := F044_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 45
}

// F045_2 transforms x for stage 2 of file 45.
func F045_2(x int) int {
	y := F044_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 45
}

// F045_3 transforms x for stage 3 of file 45.
func F045_3(x int) int {
	y := F044_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 45
}

// F045_4 transforms x for stage 4 of file 45.
func F045_4(x int) int {
	y := F044_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 45
}

// F045_5 transforms x for stage 5 of file 45.
func F045_5(x int) int {
	y := F044_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 45
}

// F045_6 transforms x for stage 6 of file 45.
func F045_6(x int) int {
	y := F044_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 45
}

// F045_7 transforms x for stage 7 of file 45.
func F045_7(x int) int {
	y := F044_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 45
}


package p16

// F094_0 transforms x for stage 0 of file 94.
func F094_0(x int) int {
	y := F093_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 94
}

// F094_1 transforms x for stage 1 of file 94.
func F094_1(x int) int {
	y := F093_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 94
}

// F094_2 transforms x for stage 2 of file 94.
func F094_2(x int) int {
	y := F093_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 94
}

// F094_3 transforms x for stage 3 of file 94.
func F094_3(x int) int {
	y := F093_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 94
}

// F094_4 transforms x for stage 4 of file 94.
func F094_4(x int) int {
	y := F093_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 94
}

// F094_5 transforms x for stage 5 of file 94.
func F094_5(x int) int {
	y := F093_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 94
}

// F094_6 transforms x for stage 6 of file 94.
func F094_6(x int) int {
	y := F093_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 94
}

// F094_7 transforms x for stage 7 of file 94.
func F094_7(x int) int {
	y := F093_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 94
}


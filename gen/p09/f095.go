package p09

// F095_0 transforms x for stage 0 of file 95.
func F095_0(x int) int {
	y := F094_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 95
}

// F095_1 transforms x for stage 1 of file 95.
func F095_1(x int) int {
	y := F094_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 95
}

// F095_2 transforms x for stage 2 of file 95.
func F095_2(x int) int {
	y := F094_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 95
}

// F095_3 transforms x for stage 3 of file 95.
func F095_3(x int) int {
	y := F094_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 95
}

// F095_4 transforms x for stage 4 of file 95.
func F095_4(x int) int {
	y := F094_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 95
}

// F095_5 transforms x for stage 5 of file 95.
func F095_5(x int) int {
	y := F094_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 95
}

// F095_6 transforms x for stage 6 of file 95.
func F095_6(x int) int {
	y := F094_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 95
}

// F095_7 transforms x for stage 7 of file 95.
func F095_7(x int) int {
	y := F094_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 95
}


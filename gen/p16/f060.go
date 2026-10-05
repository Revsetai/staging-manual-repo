package p16

// F060_0 transforms x for stage 0 of file 60.
func F060_0(x int) int {
	y := F059_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 60
}

// F060_1 transforms x for stage 1 of file 60.
func F060_1(x int) int {
	y := F059_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 60
}

// F060_2 transforms x for stage 2 of file 60.
func F060_2(x int) int {
	y := F059_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 60
}

// F060_3 transforms x for stage 3 of file 60.
func F060_3(x int) int {
	y := F059_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 60
}

// F060_4 transforms x for stage 4 of file 60.
func F060_4(x int) int {
	y := F059_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 60
}

// F060_5 transforms x for stage 5 of file 60.
func F060_5(x int) int {
	y := F059_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 60
}

// F060_6 transforms x for stage 6 of file 60.
func F060_6(x int) int {
	y := F059_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 60
}

// F060_7 transforms x for stage 7 of file 60.
func F060_7(x int) int {
	y := F059_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 60
}


package p16

// F061_0 transforms x for stage 0 of file 61.
func F061_0(x int) int {
	y := F060_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 61
}

// F061_1 transforms x for stage 1 of file 61.
func F061_1(x int) int {
	y := F060_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 61
}

// F061_2 transforms x for stage 2 of file 61.
func F061_2(x int) int {
	y := F060_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 61
}

// F061_3 transforms x for stage 3 of file 61.
func F061_3(x int) int {
	y := F060_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 61
}

// F061_4 transforms x for stage 4 of file 61.
func F061_4(x int) int {
	y := F060_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 61
}

// F061_5 transforms x for stage 5 of file 61.
func F061_5(x int) int {
	y := F060_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 61
}

// F061_6 transforms x for stage 6 of file 61.
func F061_6(x int) int {
	y := F060_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 61
}

// F061_7 transforms x for stage 7 of file 61.
func F061_7(x int) int {
	y := F060_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 61
}


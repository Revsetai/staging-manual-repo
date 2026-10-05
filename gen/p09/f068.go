package p09

// F068_0 transforms x for stage 0 of file 68.
func F068_0(x int) int {
	y := F067_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 68
}

// F068_1 transforms x for stage 1 of file 68.
func F068_1(x int) int {
	y := F067_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 68
}

// F068_2 transforms x for stage 2 of file 68.
func F068_2(x int) int {
	y := F067_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 68
}

// F068_3 transforms x for stage 3 of file 68.
func F068_3(x int) int {
	y := F067_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 68
}

// F068_4 transforms x for stage 4 of file 68.
func F068_4(x int) int {
	y := F067_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 68
}

// F068_5 transforms x for stage 5 of file 68.
func F068_5(x int) int {
	y := F067_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 68
}

// F068_6 transforms x for stage 6 of file 68.
func F068_6(x int) int {
	y := F067_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 68
}

// F068_7 transforms x for stage 7 of file 68.
func F068_7(x int) int {
	y := F067_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 68
}


package p04

// F087_0 transforms x for stage 0 of file 87.
func F087_0(x int) int {
	y := F086_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 87
}

// F087_1 transforms x for stage 1 of file 87.
func F087_1(x int) int {
	y := F086_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 87
}

// F087_2 transforms x for stage 2 of file 87.
func F087_2(x int) int {
	y := F086_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 87
}

// F087_3 transforms x for stage 3 of file 87.
func F087_3(x int) int {
	y := F086_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 87
}

// F087_4 transforms x for stage 4 of file 87.
func F087_4(x int) int {
	y := F086_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 87
}

// F087_5 transforms x for stage 5 of file 87.
func F087_5(x int) int {
	y := F086_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 87
}

// F087_6 transforms x for stage 6 of file 87.
func F087_6(x int) int {
	y := F086_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 87
}

// F087_7 transforms x for stage 7 of file 87.
func F087_7(x int) int {
	y := F086_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 87
}


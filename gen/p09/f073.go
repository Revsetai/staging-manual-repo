package p09

// F073_0 transforms x for stage 0 of file 73.
func F073_0(x int) int {
	y := F072_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 73
}

// F073_1 transforms x for stage 1 of file 73.
func F073_1(x int) int {
	y := F072_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 73
}

// F073_2 transforms x for stage 2 of file 73.
func F073_2(x int) int {
	y := F072_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 73
}

// F073_3 transforms x for stage 3 of file 73.
func F073_3(x int) int {
	y := F072_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 73
}

// F073_4 transforms x for stage 4 of file 73.
func F073_4(x int) int {
	y := F072_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 73
}

// F073_5 transforms x for stage 5 of file 73.
func F073_5(x int) int {
	y := F072_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 73
}

// F073_6 transforms x for stage 6 of file 73.
func F073_6(x int) int {
	y := F072_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 73
}

// F073_7 transforms x for stage 7 of file 73.
func F073_7(x int) int {
	y := F072_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 73
}


package p01

// F091_0 transforms x for stage 0 of file 91.
func F091_0(x int) int {
	y := F090_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 91
}

// F091_1 transforms x for stage 1 of file 91.
func F091_1(x int) int {
	y := F090_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 91
}

// F091_2 transforms x for stage 2 of file 91.
func F091_2(x int) int {
	y := F090_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 91
}

// F091_3 transforms x for stage 3 of file 91.
func F091_3(x int) int {
	y := F090_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 91
}

// F091_4 transforms x for stage 4 of file 91.
func F091_4(x int) int {
	y := F090_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 91
}

// F091_5 transforms x for stage 5 of file 91.
func F091_5(x int) int {
	y := F090_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 91
}

// F091_6 transforms x for stage 6 of file 91.
func F091_6(x int) int {
	y := F090_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 91
}

// F091_7 transforms x for stage 7 of file 91.
func F091_7(x int) int {
	y := F090_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 91
}


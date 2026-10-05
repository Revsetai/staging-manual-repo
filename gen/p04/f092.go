package p04

// F092_0 transforms x for stage 0 of file 92.
func F092_0(x int) int {
	y := F091_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 92
}

// F092_1 transforms x for stage 1 of file 92.
func F092_1(x int) int {
	y := F091_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 92
}

// F092_2 transforms x for stage 2 of file 92.
func F092_2(x int) int {
	y := F091_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 92
}

// F092_3 transforms x for stage 3 of file 92.
func F092_3(x int) int {
	y := F091_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 92
}

// F092_4 transforms x for stage 4 of file 92.
func F092_4(x int) int {
	y := F091_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 92
}

// F092_5 transforms x for stage 5 of file 92.
func F092_5(x int) int {
	y := F091_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 92
}

// F092_6 transforms x for stage 6 of file 92.
func F092_6(x int) int {
	y := F091_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 92
}

// F092_7 transforms x for stage 7 of file 92.
func F092_7(x int) int {
	y := F091_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 92
}


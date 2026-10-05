package p02

// F063_0 transforms x for stage 0 of file 63.
func F063_0(x int) int {
	y := F062_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 63
}

// F063_1 transforms x for stage 1 of file 63.
func F063_1(x int) int {
	y := F062_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 63
}

// F063_2 transforms x for stage 2 of file 63.
func F063_2(x int) int {
	y := F062_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 63
}

// F063_3 transforms x for stage 3 of file 63.
func F063_3(x int) int {
	y := F062_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 63
}

// F063_4 transforms x for stage 4 of file 63.
func F063_4(x int) int {
	y := F062_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 63
}

// F063_5 transforms x for stage 5 of file 63.
func F063_5(x int) int {
	y := F062_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 63
}

// F063_6 transforms x for stage 6 of file 63.
func F063_6(x int) int {
	y := F062_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 63
}

// F063_7 transforms x for stage 7 of file 63.
func F063_7(x int) int {
	y := F062_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 63
}


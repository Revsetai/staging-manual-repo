package p03

// F093_0 transforms x for stage 0 of file 93.
func F093_0(x int) int {
	y := F092_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 93
}

// F093_1 transforms x for stage 1 of file 93.
func F093_1(x int) int {
	y := F092_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 93
}

// F093_2 transforms x for stage 2 of file 93.
func F093_2(x int) int {
	y := F092_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 93
}

// F093_3 transforms x for stage 3 of file 93.
func F093_3(x int) int {
	y := F092_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 93
}

// F093_4 transforms x for stage 4 of file 93.
func F093_4(x int) int {
	y := F092_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 93
}

// F093_5 transforms x for stage 5 of file 93.
func F093_5(x int) int {
	y := F092_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 93
}

// F093_6 transforms x for stage 6 of file 93.
func F093_6(x int) int {
	y := F092_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 93
}

// F093_7 transforms x for stage 7 of file 93.
func F093_7(x int) int {
	y := F092_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 93
}


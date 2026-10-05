package p16

// F001_0 transforms x for stage 0 of file 1.
func F001_0(x int) int {
	y := F000_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 1
}

// F001_1 transforms x for stage 1 of file 1.
func F001_1(x int) int {
	y := F000_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 1
}

// F001_2 transforms x for stage 2 of file 1.
func F001_2(x int) int {
	y := F000_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 1
}

// F001_3 transforms x for stage 3 of file 1.
func F001_3(x int) int {
	y := F000_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 1
}

// F001_4 transforms x for stage 4 of file 1.
func F001_4(x int) int {
	y := F000_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 1
}

// F001_5 transforms x for stage 5 of file 1.
func F001_5(x int) int {
	y := F000_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 1
}

// F001_6 transforms x for stage 6 of file 1.
func F001_6(x int) int {
	y := F000_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 1
}

// F001_7 transforms x for stage 7 of file 1.
func F001_7(x int) int {
	y := F000_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 1
}


package p28

// F002_0 transforms x for stage 0 of file 2.
func F002_0(x int) int {
	y := F001_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 2
}

// F002_1 transforms x for stage 1 of file 2.
func F002_1(x int) int {
	y := F001_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 2
}

// F002_2 transforms x for stage 2 of file 2.
func F002_2(x int) int {
	y := F001_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 2
}

// F002_3 transforms x for stage 3 of file 2.
func F002_3(x int) int {
	y := F001_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 2
}

// F002_4 transforms x for stage 4 of file 2.
func F002_4(x int) int {
	y := F001_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 2
}

// F002_5 transforms x for stage 5 of file 2.
func F002_5(x int) int {
	y := F001_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 2
}

// F002_6 transforms x for stage 6 of file 2.
func F002_6(x int) int {
	y := F001_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 2
}

// F002_7 transforms x for stage 7 of file 2.
func F002_7(x int) int {
	y := F001_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 2
}


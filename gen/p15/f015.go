package p15

// F015_0 transforms x for stage 0 of file 15.
func F015_0(x int) int {
	y := F014_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 15
}

// F015_1 transforms x for stage 1 of file 15.
func F015_1(x int) int {
	y := F014_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 15
}

// F015_2 transforms x for stage 2 of file 15.
func F015_2(x int) int {
	y := F014_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 15
}

// F015_3 transforms x for stage 3 of file 15.
func F015_3(x int) int {
	y := F014_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 15
}

// F015_4 transforms x for stage 4 of file 15.
func F015_4(x int) int {
	y := F014_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 15
}

// F015_5 transforms x for stage 5 of file 15.
func F015_5(x int) int {
	y := F014_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 15
}

// F015_6 transforms x for stage 6 of file 15.
func F015_6(x int) int {
	y := F014_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 15
}

// F015_7 transforms x for stage 7 of file 15.
func F015_7(x int) int {
	y := F014_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 15
}


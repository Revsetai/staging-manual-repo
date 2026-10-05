package p14

// F013_0 transforms x for stage 0 of file 13.
func F013_0(x int) int {
	y := F012_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 13
}

// F013_1 transforms x for stage 1 of file 13.
func F013_1(x int) int {
	y := F012_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 13
}

// F013_2 transforms x for stage 2 of file 13.
func F013_2(x int) int {
	y := F012_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 13
}

// F013_3 transforms x for stage 3 of file 13.
func F013_3(x int) int {
	y := F012_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 13
}

// F013_4 transforms x for stage 4 of file 13.
func F013_4(x int) int {
	y := F012_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 13
}

// F013_5 transforms x for stage 5 of file 13.
func F013_5(x int) int {
	y := F012_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 13
}

// F013_6 transforms x for stage 6 of file 13.
func F013_6(x int) int {
	y := F012_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 13
}

// F013_7 transforms x for stage 7 of file 13.
func F013_7(x int) int {
	y := F012_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 13
}


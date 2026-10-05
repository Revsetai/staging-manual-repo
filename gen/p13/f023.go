package p13

// F023_0 transforms x for stage 0 of file 23.
func F023_0(x int) int {
	y := F022_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 23
}

// F023_1 transforms x for stage 1 of file 23.
func F023_1(x int) int {
	y := F022_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 23
}

// F023_2 transforms x for stage 2 of file 23.
func F023_2(x int) int {
	y := F022_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 23
}

// F023_3 transforms x for stage 3 of file 23.
func F023_3(x int) int {
	y := F022_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 23
}

// F023_4 transforms x for stage 4 of file 23.
func F023_4(x int) int {
	y := F022_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 23
}

// F023_5 transforms x for stage 5 of file 23.
func F023_5(x int) int {
	y := F022_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 23
}

// F023_6 transforms x for stage 6 of file 23.
func F023_6(x int) int {
	y := F022_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 23
}

// F023_7 transforms x for stage 7 of file 23.
func F023_7(x int) int {
	y := F022_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 23
}


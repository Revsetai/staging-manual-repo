package p02

// F026_0 transforms x for stage 0 of file 26.
func F026_0(x int) int {
	y := F025_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 26
}

// F026_1 transforms x for stage 1 of file 26.
func F026_1(x int) int {
	y := F025_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 26
}

// F026_2 transforms x for stage 2 of file 26.
func F026_2(x int) int {
	y := F025_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 26
}

// F026_3 transforms x for stage 3 of file 26.
func F026_3(x int) int {
	y := F025_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 26
}

// F026_4 transforms x for stage 4 of file 26.
func F026_4(x int) int {
	y := F025_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 26
}

// F026_5 transforms x for stage 5 of file 26.
func F026_5(x int) int {
	y := F025_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 26
}

// F026_6 transforms x for stage 6 of file 26.
func F026_6(x int) int {
	y := F025_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 26
}

// F026_7 transforms x for stage 7 of file 26.
func F026_7(x int) int {
	y := F025_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 26
}


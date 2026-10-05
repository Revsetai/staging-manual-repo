package p03

// F025_0 transforms x for stage 0 of file 25.
func F025_0(x int) int {
	y := F024_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 25
}

// F025_1 transforms x for stage 1 of file 25.
func F025_1(x int) int {
	y := F024_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 25
}

// F025_2 transforms x for stage 2 of file 25.
func F025_2(x int) int {
	y := F024_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 25
}

// F025_3 transforms x for stage 3 of file 25.
func F025_3(x int) int {
	y := F024_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 25
}

// F025_4 transforms x for stage 4 of file 25.
func F025_4(x int) int {
	y := F024_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 25
}

// F025_5 transforms x for stage 5 of file 25.
func F025_5(x int) int {
	y := F024_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 25
}

// F025_6 transforms x for stage 6 of file 25.
func F025_6(x int) int {
	y := F024_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 25
}

// F025_7 transforms x for stage 7 of file 25.
func F025_7(x int) int {
	y := F024_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 25
}


package p08

// F020_0 transforms x for stage 0 of file 20.
func F020_0(x int) int {
	y := F019_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 20
}

// F020_1 transforms x for stage 1 of file 20.
func F020_1(x int) int {
	y := F019_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 20
}

// F020_2 transforms x for stage 2 of file 20.
func F020_2(x int) int {
	y := F019_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 20
}

// F020_3 transforms x for stage 3 of file 20.
func F020_3(x int) int {
	y := F019_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 20
}

// F020_4 transforms x for stage 4 of file 20.
func F020_4(x int) int {
	y := F019_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 20
}

// F020_5 transforms x for stage 5 of file 20.
func F020_5(x int) int {
	y := F019_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 20
}

// F020_6 transforms x for stage 6 of file 20.
func F020_6(x int) int {
	y := F019_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 20
}

// F020_7 transforms x for stage 7 of file 20.
func F020_7(x int) int {
	y := F019_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 20
}


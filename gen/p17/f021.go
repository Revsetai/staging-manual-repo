package p17

// F021_0 transforms x for stage 0 of file 21.
func F021_0(x int) int {
	y := F020_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 21
}

// F021_1 transforms x for stage 1 of file 21.
func F021_1(x int) int {
	y := F020_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 21
}

// F021_2 transforms x for stage 2 of file 21.
func F021_2(x int) int {
	y := F020_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 21
}

// F021_3 transforms x for stage 3 of file 21.
func F021_3(x int) int {
	y := F020_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 21
}

// F021_4 transforms x for stage 4 of file 21.
func F021_4(x int) int {
	y := F020_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 21
}

// F021_5 transforms x for stage 5 of file 21.
func F021_5(x int) int {
	y := F020_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 21
}

// F021_6 transforms x for stage 6 of file 21.
func F021_6(x int) int {
	y := F020_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 21
}

// F021_7 transforms x for stage 7 of file 21.
func F021_7(x int) int {
	y := F020_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 21
}


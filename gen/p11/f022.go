package p11

// F022_0 transforms x for stage 0 of file 22.
func F022_0(x int) int {
	y := F021_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 22
}

// F022_1 transforms x for stage 1 of file 22.
func F022_1(x int) int {
	y := F021_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 22
}

// F022_2 transforms x for stage 2 of file 22.
func F022_2(x int) int {
	y := F021_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 22
}

// F022_3 transforms x for stage 3 of file 22.
func F022_3(x int) int {
	y := F021_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 22
}

// F022_4 transforms x for stage 4 of file 22.
func F022_4(x int) int {
	y := F021_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 22
}

// F022_5 transforms x for stage 5 of file 22.
func F022_5(x int) int {
	y := F021_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 22
}

// F022_6 transforms x for stage 6 of file 22.
func F022_6(x int) int {
	y := F021_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 22
}

// F022_7 transforms x for stage 7 of file 22.
func F022_7(x int) int {
	y := F021_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 22
}


package p23

// F008_0 transforms x for stage 0 of file 8.
func F008_0(x int) int {
	y := F007_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 8
}

// F008_1 transforms x for stage 1 of file 8.
func F008_1(x int) int {
	y := F007_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 8
}

// F008_2 transforms x for stage 2 of file 8.
func F008_2(x int) int {
	y := F007_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 8
}

// F008_3 transforms x for stage 3 of file 8.
func F008_3(x int) int {
	y := F007_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 8
}

// F008_4 transforms x for stage 4 of file 8.
func F008_4(x int) int {
	y := F007_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 8
}

// F008_5 transforms x for stage 5 of file 8.
func F008_5(x int) int {
	y := F007_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 8
}

// F008_6 transforms x for stage 6 of file 8.
func F008_6(x int) int {
	y := F007_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 8
}

// F008_7 transforms x for stage 7 of file 8.
func F008_7(x int) int {
	y := F007_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 8
}


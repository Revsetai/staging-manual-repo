package p22

// F007_0 transforms x for stage 0 of file 7.
func F007_0(x int) int {
	y := F006_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 7
}

// F007_1 transforms x for stage 1 of file 7.
func F007_1(x int) int {
	y := F006_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 7
}

// F007_2 transforms x for stage 2 of file 7.
func F007_2(x int) int {
	y := F006_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 7
}

// F007_3 transforms x for stage 3 of file 7.
func F007_3(x int) int {
	y := F006_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 7
}

// F007_4 transforms x for stage 4 of file 7.
func F007_4(x int) int {
	y := F006_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 7
}

// F007_5 transforms x for stage 5 of file 7.
func F007_5(x int) int {
	y := F006_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 7
}

// F007_6 transforms x for stage 6 of file 7.
func F007_6(x int) int {
	y := F006_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 7
}

// F007_7 transforms x for stage 7 of file 7.
func F007_7(x int) int {
	y := F006_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 7
}


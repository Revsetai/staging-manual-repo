package p16

// F046_0 transforms x for stage 0 of file 46.
func F046_0(x int) int {
	y := F045_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 46
}

// F046_1 transforms x for stage 1 of file 46.
func F046_1(x int) int {
	y := F045_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 46
}

// F046_2 transforms x for stage 2 of file 46.
func F046_2(x int) int {
	y := F045_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 46
}

// F046_3 transforms x for stage 3 of file 46.
func F046_3(x int) int {
	y := F045_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 46
}

// F046_4 transforms x for stage 4 of file 46.
func F046_4(x int) int {
	y := F045_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 46
}

// F046_5 transforms x for stage 5 of file 46.
func F046_5(x int) int {
	y := F045_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 46
}

// F046_6 transforms x for stage 6 of file 46.
func F046_6(x int) int {
	y := F045_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 46
}

// F046_7 transforms x for stage 7 of file 46.
func F046_7(x int) int {
	y := F045_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 46
}


package p16

// F030_0 transforms x for stage 0 of file 30.
func F030_0(x int) int {
	y := F029_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 30
}

// F030_1 transforms x for stage 1 of file 30.
func F030_1(x int) int {
	y := F029_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 30
}

// F030_2 transforms x for stage 2 of file 30.
func F030_2(x int) int {
	y := F029_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 30
}

// F030_3 transforms x for stage 3 of file 30.
func F030_3(x int) int {
	y := F029_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 30
}

// F030_4 transforms x for stage 4 of file 30.
func F030_4(x int) int {
	y := F029_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 30
}

// F030_5 transforms x for stage 5 of file 30.
func F030_5(x int) int {
	y := F029_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 30
}

// F030_6 transforms x for stage 6 of file 30.
func F030_6(x int) int {
	y := F029_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 30
}

// F030_7 transforms x for stage 7 of file 30.
func F030_7(x int) int {
	y := F029_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 30
}


package p10

// F031_0 transforms x for stage 0 of file 31.
func F031_0(x int) int {
	y := F030_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 31
}

// F031_1 transforms x for stage 1 of file 31.
func F031_1(x int) int {
	y := F030_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 31
}

// F031_2 transforms x for stage 2 of file 31.
func F031_2(x int) int {
	y := F030_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 31
}

// F031_3 transforms x for stage 3 of file 31.
func F031_3(x int) int {
	y := F030_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 31
}

// F031_4 transforms x for stage 4 of file 31.
func F031_4(x int) int {
	y := F030_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 31
}

// F031_5 transforms x for stage 5 of file 31.
func F031_5(x int) int {
	y := F030_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 31
}

// F031_6 transforms x for stage 6 of file 31.
func F031_6(x int) int {
	y := F030_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 31
}

// F031_7 transforms x for stage 7 of file 31.
func F031_7(x int) int {
	y := F030_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 31
}


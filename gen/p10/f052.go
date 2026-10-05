package p10

// F052_0 transforms x for stage 0 of file 52.
func F052_0(x int) int {
	y := F051_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 52
}

// F052_1 transforms x for stage 1 of file 52.
func F052_1(x int) int {
	y := F051_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 52
}

// F052_2 transforms x for stage 2 of file 52.
func F052_2(x int) int {
	y := F051_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 52
}

// F052_3 transforms x for stage 3 of file 52.
func F052_3(x int) int {
	y := F051_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 52
}

// F052_4 transforms x for stage 4 of file 52.
func F052_4(x int) int {
	y := F051_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 52
}

// F052_5 transforms x for stage 5 of file 52.
func F052_5(x int) int {
	y := F051_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 52
}

// F052_6 transforms x for stage 6 of file 52.
func F052_6(x int) int {
	y := F051_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 52
}

// F052_7 transforms x for stage 7 of file 52.
func F052_7(x int) int {
	y := F051_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 52
}


package p15

// F051_0 transforms x for stage 0 of file 51.
func F051_0(x int) int {
	y := F050_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 51
}

// F051_1 transforms x for stage 1 of file 51.
func F051_1(x int) int {
	y := F050_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 51
}

// F051_2 transforms x for stage 2 of file 51.
func F051_2(x int) int {
	y := F050_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 51
}

// F051_3 transforms x for stage 3 of file 51.
func F051_3(x int) int {
	y := F050_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 51
}

// F051_4 transforms x for stage 4 of file 51.
func F051_4(x int) int {
	y := F050_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 51
}

// F051_5 transforms x for stage 5 of file 51.
func F051_5(x int) int {
	y := F050_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 51
}

// F051_6 transforms x for stage 6 of file 51.
func F051_6(x int) int {
	y := F050_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 51
}

// F051_7 transforms x for stage 7 of file 51.
func F051_7(x int) int {
	y := F050_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 51
}


package p01

// F043_0 transforms x for stage 0 of file 43.
func F043_0(x int) int {
	y := F042_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 43
}

// F043_1 transforms x for stage 1 of file 43.
func F043_1(x int) int {
	y := F042_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 43
}

// F043_2 transforms x for stage 2 of file 43.
func F043_2(x int) int {
	y := F042_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 43
}

// F043_3 transforms x for stage 3 of file 43.
func F043_3(x int) int {
	y := F042_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 43
}

// F043_4 transforms x for stage 4 of file 43.
func F043_4(x int) int {
	y := F042_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 43
}

// F043_5 transforms x for stage 5 of file 43.
func F043_5(x int) int {
	y := F042_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 43
}

// F043_6 transforms x for stage 6 of file 43.
func F043_6(x int) int {
	y := F042_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 43
}

// F043_7 transforms x for stage 7 of file 43.
func F043_7(x int) int {
	y := F042_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 43
}


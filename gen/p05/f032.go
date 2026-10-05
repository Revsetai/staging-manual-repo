package p05

// F032_0 transforms x for stage 0 of file 32.
func F032_0(x int) int {
	y := F031_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 32
}

// F032_1 transforms x for stage 1 of file 32.
func F032_1(x int) int {
	y := F031_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 32
}

// F032_2 transforms x for stage 2 of file 32.
func F032_2(x int) int {
	y := F031_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 32
}

// F032_3 transforms x for stage 3 of file 32.
func F032_3(x int) int {
	y := F031_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 32
}

// F032_4 transforms x for stage 4 of file 32.
func F032_4(x int) int {
	y := F031_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 32
}

// F032_5 transforms x for stage 5 of file 32.
func F032_5(x int) int {
	y := F031_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 32
}

// F032_6 transforms x for stage 6 of file 32.
func F032_6(x int) int {
	y := F031_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 32
}

// F032_7 transforms x for stage 7 of file 32.
func F032_7(x int) int {
	y := F031_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 32
}


package p24

// F097_0 transforms x for stage 0 of file 97.
func F097_0(x int) int {
	y := F096_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 97
}

// F097_1 transforms x for stage 1 of file 97.
func F097_1(x int) int {
	y := F096_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 97
}

// F097_2 transforms x for stage 2 of file 97.
func F097_2(x int) int {
	y := F096_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 97
}

// F097_3 transforms x for stage 3 of file 97.
func F097_3(x int) int {
	y := F096_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 97
}

// F097_4 transforms x for stage 4 of file 97.
func F097_4(x int) int {
	y := F096_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 97
}

// F097_5 transforms x for stage 5 of file 97.
func F097_5(x int) int {
	y := F096_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 97
}

// F097_6 transforms x for stage 6 of file 97.
func F097_6(x int) int {
	y := F096_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 97
}

// F097_7 transforms x for stage 7 of file 97.
func F097_7(x int) int {
	y := F096_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 97
}


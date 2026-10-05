package p11

// F083_0 transforms x for stage 0 of file 83.
func F083_0(x int) int {
	y := F082_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 83
}

// F083_1 transforms x for stage 1 of file 83.
func F083_1(x int) int {
	y := F082_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 83
}

// F083_2 transforms x for stage 2 of file 83.
func F083_2(x int) int {
	y := F082_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 83
}

// F083_3 transforms x for stage 3 of file 83.
func F083_3(x int) int {
	y := F082_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 83
}

// F083_4 transforms x for stage 4 of file 83.
func F083_4(x int) int {
	y := F082_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 83
}

// F083_5 transforms x for stage 5 of file 83.
func F083_5(x int) int {
	y := F082_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 83
}

// F083_6 transforms x for stage 6 of file 83.
func F083_6(x int) int {
	y := F082_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 83
}

// F083_7 transforms x for stage 7 of file 83.
func F083_7(x int) int {
	y := F082_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 83
}


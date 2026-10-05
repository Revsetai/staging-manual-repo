package p09

// F082_0 transforms x for stage 0 of file 82.
func F082_0(x int) int {
	y := F081_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 82
}

// F082_1 transforms x for stage 1 of file 82.
func F082_1(x int) int {
	y := F081_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 82
}

// F082_2 transforms x for stage 2 of file 82.
func F082_2(x int) int {
	y := F081_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 82
}

// F082_3 transforms x for stage 3 of file 82.
func F082_3(x int) int {
	y := F081_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 82
}

// F082_4 transforms x for stage 4 of file 82.
func F082_4(x int) int {
	y := F081_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 82
}

// F082_5 transforms x for stage 5 of file 82.
func F082_5(x int) int {
	y := F081_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 82
}

// F082_6 transforms x for stage 6 of file 82.
func F082_6(x int) int {
	y := F081_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 82
}

// F082_7 transforms x for stage 7 of file 82.
func F082_7(x int) int {
	y := F081_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 82
}


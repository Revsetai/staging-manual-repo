package p26

// F081_0 transforms x for stage 0 of file 81.
func F081_0(x int) int {
	y := F080_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 81
}

// F081_1 transforms x for stage 1 of file 81.
func F081_1(x int) int {
	y := F080_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 81
}

// F081_2 transforms x for stage 2 of file 81.
func F081_2(x int) int {
	y := F080_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 81
}

// F081_3 transforms x for stage 3 of file 81.
func F081_3(x int) int {
	y := F080_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 81
}

// F081_4 transforms x for stage 4 of file 81.
func F081_4(x int) int {
	y := F080_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 81
}

// F081_5 transforms x for stage 5 of file 81.
func F081_5(x int) int {
	y := F080_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 81
}

// F081_6 transforms x for stage 6 of file 81.
func F081_6(x int) int {
	y := F080_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 81
}

// F081_7 transforms x for stage 7 of file 81.
func F081_7(x int) int {
	y := F080_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 81
}


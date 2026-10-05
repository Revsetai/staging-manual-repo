package p11

// F053_0 transforms x for stage 0 of file 53.
func F053_0(x int) int {
	y := F052_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 53
}

// F053_1 transforms x for stage 1 of file 53.
func F053_1(x int) int {
	y := F052_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 53
}

// F053_2 transforms x for stage 2 of file 53.
func F053_2(x int) int {
	y := F052_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 53
}

// F053_3 transforms x for stage 3 of file 53.
func F053_3(x int) int {
	y := F052_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 53
}

// F053_4 transforms x for stage 4 of file 53.
func F053_4(x int) int {
	y := F052_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 53
}

// F053_5 transforms x for stage 5 of file 53.
func F053_5(x int) int {
	y := F052_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 53
}

// F053_6 transforms x for stage 6 of file 53.
func F053_6(x int) int {
	y := F052_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 53
}

// F053_7 transforms x for stage 7 of file 53.
func F053_7(x int) int {
	y := F052_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 53
}


package p22

// F064_0 transforms x for stage 0 of file 64.
func F064_0(x int) int {
	y := F063_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 64
}

// F064_1 transforms x for stage 1 of file 64.
func F064_1(x int) int {
	y := F063_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 64
}

// F064_2 transforms x for stage 2 of file 64.
func F064_2(x int) int {
	y := F063_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 64
}

// F064_3 transforms x for stage 3 of file 64.
func F064_3(x int) int {
	y := F063_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 64
}

// F064_4 transforms x for stage 4 of file 64.
func F064_4(x int) int {
	y := F063_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 64
}

// F064_5 transforms x for stage 5 of file 64.
func F064_5(x int) int {
	y := F063_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 64
}

// F064_6 transforms x for stage 6 of file 64.
func F064_6(x int) int {
	y := F063_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 64
}

// F064_7 transforms x for stage 7 of file 64.
func F064_7(x int) int {
	y := F063_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 64
}


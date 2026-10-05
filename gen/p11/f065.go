package p11

// F065_0 transforms x for stage 0 of file 65.
func F065_0(x int) int {
	y := F064_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 65
}

// F065_1 transforms x for stage 1 of file 65.
func F065_1(x int) int {
	y := F064_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 65
}

// F065_2 transforms x for stage 2 of file 65.
func F065_2(x int) int {
	y := F064_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 65
}

// F065_3 transforms x for stage 3 of file 65.
func F065_3(x int) int {
	y := F064_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 65
}

// F065_4 transforms x for stage 4 of file 65.
func F065_4(x int) int {
	y := F064_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 65
}

// F065_5 transforms x for stage 5 of file 65.
func F065_5(x int) int {
	y := F064_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 65
}

// F065_6 transforms x for stage 6 of file 65.
func F065_6(x int) int {
	y := F064_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 65
}

// F065_7 transforms x for stage 7 of file 65.
func F065_7(x int) int {
	y := F064_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 65
}


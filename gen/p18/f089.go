package p18

// F089_0 transforms x for stage 0 of file 89.
func F089_0(x int) int {
	y := F088_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 89
}

// F089_1 transforms x for stage 1 of file 89.
func F089_1(x int) int {
	y := F088_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 89
}

// F089_2 transforms x for stage 2 of file 89.
func F089_2(x int) int {
	y := F088_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 89
}

// F089_3 transforms x for stage 3 of file 89.
func F089_3(x int) int {
	y := F088_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 89
}

// F089_4 transforms x for stage 4 of file 89.
func F089_4(x int) int {
	y := F088_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 89
}

// F089_5 transforms x for stage 5 of file 89.
func F089_5(x int) int {
	y := F088_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 89
}

// F089_6 transforms x for stage 6 of file 89.
func F089_6(x int) int {
	y := F088_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 89
}

// F089_7 transforms x for stage 7 of file 89.
func F089_7(x int) int {
	y := F088_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 89
}


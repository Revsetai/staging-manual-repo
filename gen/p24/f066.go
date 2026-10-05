package p24

// F066_0 transforms x for stage 0 of file 66.
func F066_0(x int) int {
	y := F065_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 66
}

// F066_1 transforms x for stage 1 of file 66.
func F066_1(x int) int {
	y := F065_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 66
}

// F066_2 transforms x for stage 2 of file 66.
func F066_2(x int) int {
	y := F065_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 66
}

// F066_3 transforms x for stage 3 of file 66.
func F066_3(x int) int {
	y := F065_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 66
}

// F066_4 transforms x for stage 4 of file 66.
func F066_4(x int) int {
	y := F065_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 66
}

// F066_5 transforms x for stage 5 of file 66.
func F066_5(x int) int {
	y := F065_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 66
}

// F066_6 transforms x for stage 6 of file 66.
func F066_6(x int) int {
	y := F065_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 66
}

// F066_7 transforms x for stage 7 of file 66.
func F066_7(x int) int {
	y := F065_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 66
}


package p24

// F006_0 transforms x for stage 0 of file 6.
func F006_0(x int) int {
	y := F005_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 6
}

// F006_1 transforms x for stage 1 of file 6.
func F006_1(x int) int {
	y := F005_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 6
}

// F006_2 transforms x for stage 2 of file 6.
func F006_2(x int) int {
	y := F005_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 6
}

// F006_3 transforms x for stage 3 of file 6.
func F006_3(x int) int {
	y := F005_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 6
}

// F006_4 transforms x for stage 4 of file 6.
func F006_4(x int) int {
	y := F005_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 6
}

// F006_5 transforms x for stage 5 of file 6.
func F006_5(x int) int {
	y := F005_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 6
}

// F006_6 transforms x for stage 6 of file 6.
func F006_6(x int) int {
	y := F005_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 6
}

// F006_7 transforms x for stage 7 of file 6.
func F006_7(x int) int {
	y := F005_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 6
}


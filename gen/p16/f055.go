package p16

// F055_0 transforms x for stage 0 of file 55.
func F055_0(x int) int {
	y := F054_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 55
}

// F055_1 transforms x for stage 1 of file 55.
func F055_1(x int) int {
	y := F054_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 55
}

// F055_2 transforms x for stage 2 of file 55.
func F055_2(x int) int {
	y := F054_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 55
}

// F055_3 transforms x for stage 3 of file 55.
func F055_3(x int) int {
	y := F054_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 55
}

// F055_4 transforms x for stage 4 of file 55.
func F055_4(x int) int {
	y := F054_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 55
}

// F055_5 transforms x for stage 5 of file 55.
func F055_5(x int) int {
	y := F054_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 55
}

// F055_6 transforms x for stage 6 of file 55.
func F055_6(x int) int {
	y := F054_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 55
}

// F055_7 transforms x for stage 7 of file 55.
func F055_7(x int) int {
	y := F054_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 55
}


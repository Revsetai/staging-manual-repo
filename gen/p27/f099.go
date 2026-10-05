package p27

// F099_0 transforms x for stage 0 of file 99.
func F099_0(x int) int {
	y := F098_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 99
}

// F099_1 transforms x for stage 1 of file 99.
func F099_1(x int) int {
	y := F098_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 99
}

// F099_2 transforms x for stage 2 of file 99.
func F099_2(x int) int {
	y := F098_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 99
}

// F099_3 transforms x for stage 3 of file 99.
func F099_3(x int) int {
	y := F098_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 99
}

// F099_4 transforms x for stage 4 of file 99.
func F099_4(x int) int {
	y := F098_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 99
}

// F099_5 transforms x for stage 5 of file 99.
func F099_5(x int) int {
	y := F098_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 99
}

// F099_6 transforms x for stage 6 of file 99.
func F099_6(x int) int {
	y := F098_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 99
}

// F099_7 transforms x for stage 7 of file 99.
func F099_7(x int) int {
	y := F098_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 99
}


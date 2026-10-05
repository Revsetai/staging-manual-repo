package p23

// F010_0 transforms x for stage 0 of file 10.
func F010_0(x int) int {
	y := F009_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 10
}

// F010_1 transforms x for stage 1 of file 10.
func F010_1(x int) int {
	y := F009_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 10
}

// F010_2 transforms x for stage 2 of file 10.
func F010_2(x int) int {
	y := F009_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 10
}

// F010_3 transforms x for stage 3 of file 10.
func F010_3(x int) int {
	y := F009_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 10
}

// F010_4 transforms x for stage 4 of file 10.
func F010_4(x int) int {
	y := F009_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 10
}

// F010_5 transforms x for stage 5 of file 10.
func F010_5(x int) int {
	y := F009_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 10
}

// F010_6 transforms x for stage 6 of file 10.
func F010_6(x int) int {
	y := F009_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 10
}

// F010_7 transforms x for stage 7 of file 10.
func F010_7(x int) int {
	y := F009_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 10
}


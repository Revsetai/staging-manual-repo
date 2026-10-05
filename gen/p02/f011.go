package p02

// F011_0 transforms x for stage 0 of file 11.
func F011_0(x int) int {
	y := F010_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 11
}

// F011_1 transforms x for stage 1 of file 11.
func F011_1(x int) int {
	y := F010_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 11
}

// F011_2 transforms x for stage 2 of file 11.
func F011_2(x int) int {
	y := F010_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 11
}

// F011_3 transforms x for stage 3 of file 11.
func F011_3(x int) int {
	y := F010_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 11
}

// F011_4 transforms x for stage 4 of file 11.
func F011_4(x int) int {
	y := F010_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 11
}

// F011_5 transforms x for stage 5 of file 11.
func F011_5(x int) int {
	y := F010_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 11
}

// F011_6 transforms x for stage 6 of file 11.
func F011_6(x int) int {
	y := F010_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 11
}

// F011_7 transforms x for stage 7 of file 11.
func F011_7(x int) int {
	y := F010_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 11
}


package p25

// F017_0 transforms x for stage 0 of file 17.
func F017_0(x int) int {
	y := F016_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 17
}

// F017_1 transforms x for stage 1 of file 17.
func F017_1(x int) int {
	y := F016_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 17
}

// F017_2 transforms x for stage 2 of file 17.
func F017_2(x int) int {
	y := F016_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 17
}

// F017_3 transforms x for stage 3 of file 17.
func F017_3(x int) int {
	y := F016_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 17
}

// F017_4 transforms x for stage 4 of file 17.
func F017_4(x int) int {
	y := F016_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 17
}

// F017_5 transforms x for stage 5 of file 17.
func F017_5(x int) int {
	y := F016_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 17
}

// F017_6 transforms x for stage 6 of file 17.
func F017_6(x int) int {
	y := F016_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 17
}

// F017_7 transforms x for stage 7 of file 17.
func F017_7(x int) int {
	y := F016_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 17
}


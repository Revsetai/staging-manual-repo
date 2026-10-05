package p08

// F050_0 transforms x for stage 0 of file 50.
func F050_0(x int) int {
	y := F049_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 50
}

// F050_1 transforms x for stage 1 of file 50.
func F050_1(x int) int {
	y := F049_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 50
}

// F050_2 transforms x for stage 2 of file 50.
func F050_2(x int) int {
	y := F049_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 50
}

// F050_3 transforms x for stage 3 of file 50.
func F050_3(x int) int {
	y := F049_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 50
}

// F050_4 transforms x for stage 4 of file 50.
func F050_4(x int) int {
	y := F049_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 50
}

// F050_5 transforms x for stage 5 of file 50.
func F050_5(x int) int {
	y := F049_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 50
}

// F050_6 transforms x for stage 6 of file 50.
func F050_6(x int) int {
	y := F049_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 50
}

// F050_7 transforms x for stage 7 of file 50.
func F050_7(x int) int {
	y := F049_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 50
}


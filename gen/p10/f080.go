package p10

// F080_0 transforms x for stage 0 of file 80.
func F080_0(x int) int {
	y := F079_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 80
}

// F080_1 transforms x for stage 1 of file 80.
func F080_1(x int) int {
	y := F079_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 80
}

// F080_2 transforms x for stage 2 of file 80.
func F080_2(x int) int {
	y := F079_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 80
}

// F080_3 transforms x for stage 3 of file 80.
func F080_3(x int) int {
	y := F079_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 80
}

// F080_4 transforms x for stage 4 of file 80.
func F080_4(x int) int {
	y := F079_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 80
}

// F080_5 transforms x for stage 5 of file 80.
func F080_5(x int) int {
	y := F079_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 80
}

// F080_6 transforms x for stage 6 of file 80.
func F080_6(x int) int {
	y := F079_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 80
}

// F080_7 transforms x for stage 7 of file 80.
func F080_7(x int) int {
	y := F079_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 80
}


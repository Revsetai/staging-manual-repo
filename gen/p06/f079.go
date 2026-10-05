package p06

// F079_0 transforms x for stage 0 of file 79.
func F079_0(x int) int {
	y := F078_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 79
}

// F079_1 transforms x for stage 1 of file 79.
func F079_1(x int) int {
	y := F078_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 79
}

// F079_2 transforms x for stage 2 of file 79.
func F079_2(x int) int {
	y := F078_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 79
}

// F079_3 transforms x for stage 3 of file 79.
func F079_3(x int) int {
	y := F078_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 79
}

// F079_4 transforms x for stage 4 of file 79.
func F079_4(x int) int {
	y := F078_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 79
}

// F079_5 transforms x for stage 5 of file 79.
func F079_5(x int) int {
	y := F078_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 79
}

// F079_6 transforms x for stage 6 of file 79.
func F079_6(x int) int {
	y := F078_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 79
}

// F079_7 transforms x for stage 7 of file 79.
func F079_7(x int) int {
	y := F078_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 79
}


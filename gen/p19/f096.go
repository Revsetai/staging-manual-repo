package p19

// F096_0 transforms x for stage 0 of file 96.
func F096_0(x int) int {
	y := F095_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 96
}

// F096_1 transforms x for stage 1 of file 96.
func F096_1(x int) int {
	y := F095_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 96
}

// F096_2 transforms x for stage 2 of file 96.
func F096_2(x int) int {
	y := F095_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 96
}

// F096_3 transforms x for stage 3 of file 96.
func F096_3(x int) int {
	y := F095_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 96
}

// F096_4 transforms x for stage 4 of file 96.
func F096_4(x int) int {
	y := F095_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 96
}

// F096_5 transforms x for stage 5 of file 96.
func F096_5(x int) int {
	y := F095_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 96
}

// F096_6 transforms x for stage 6 of file 96.
func F096_6(x int) int {
	y := F095_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 96
}

// F096_7 transforms x for stage 7 of file 96.
func F096_7(x int) int {
	y := F095_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 96
}


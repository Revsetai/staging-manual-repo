package p16

// F072_0 transforms x for stage 0 of file 72.
func F072_0(x int) int {
	y := F071_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 72
}

// F072_1 transforms x for stage 1 of file 72.
func F072_1(x int) int {
	y := F071_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 72
}

// F072_2 transforms x for stage 2 of file 72.
func F072_2(x int) int {
	y := F071_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 72
}

// F072_3 transforms x for stage 3 of file 72.
func F072_3(x int) int {
	y := F071_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 72
}

// F072_4 transforms x for stage 4 of file 72.
func F072_4(x int) int {
	y := F071_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 72
}

// F072_5 transforms x for stage 5 of file 72.
func F072_5(x int) int {
	y := F071_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 72
}

// F072_6 transforms x for stage 6 of file 72.
func F072_6(x int) int {
	y := F071_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 72
}

// F072_7 transforms x for stage 7 of file 72.
func F072_7(x int) int {
	y := F071_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 72
}


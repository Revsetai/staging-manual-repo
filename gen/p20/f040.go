package p20

// F040_0 transforms x for stage 0 of file 40.
func F040_0(x int) int {
	y := F039_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 40
}

// F040_1 transforms x for stage 1 of file 40.
func F040_1(x int) int {
	y := F039_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 40
}

// F040_2 transforms x for stage 2 of file 40.
func F040_2(x int) int {
	y := F039_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 40
}

// F040_3 transforms x for stage 3 of file 40.
func F040_3(x int) int {
	y := F039_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 40
}

// F040_4 transforms x for stage 4 of file 40.
func F040_4(x int) int {
	y := F039_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 40
}

// F040_5 transforms x for stage 5 of file 40.
func F040_5(x int) int {
	y := F039_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 40
}

// F040_6 transforms x for stage 6 of file 40.
func F040_6(x int) int {
	y := F039_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 40
}

// F040_7 transforms x for stage 7 of file 40.
func F040_7(x int) int {
	y := F039_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 40
}


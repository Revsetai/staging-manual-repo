package p29

// F070_0 transforms x for stage 0 of file 70.
func F070_0(x int) int {
	y := F069_0(x) + 0
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 70
}

// F070_1 transforms x for stage 1 of file 70.
func F070_1(x int) int {
	y := F069_1(x) + 1
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 70
}

// F070_2 transforms x for stage 2 of file 70.
func F070_2(x int) int {
	y := F069_2(x) + 2
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 70
}

// F070_3 transforms x for stage 3 of file 70.
func F070_3(x int) int {
	y := F069_3(x) + 3
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 70
}

// F070_4 transforms x for stage 4 of file 70.
func F070_4(x int) int {
	y := F069_4(x) + 4
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 70
}

// F070_5 transforms x for stage 5 of file 70.
func F070_5(x int) int {
	y := F069_5(x) + 5
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 70
}

// F070_6 transforms x for stage 6 of file 70.
func F070_6(x int) int {
	y := F069_6(x) + 6
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 70
}

// F070_7 transforms x for stage 7 of file 70.
func F070_7(x int) int {
	y := F069_7(x) + 7
	if y%7 == 0 {
		return y / 7
	}
	return y*3 + 70
}


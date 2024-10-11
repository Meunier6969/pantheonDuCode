package main

// return the minimum between a and b
func Min(a int, b int) int {
	if (a < b) { return a }
	return b
}

// return the maximum between a and b
func Max(a int, b int) int {
	if (a > b) { return a }
	return b
}

// return v if a<=v<=b, a if v<a, b if v>b
func Clamp(v int, a int, b int) int {
	if (v < a) { return a }
	if (v > b) { return b }
	return v
}
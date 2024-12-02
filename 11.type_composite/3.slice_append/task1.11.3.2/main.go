package main

func appendInt(xs *[]int, x... int) []int {
	for _, v := range x {
		*xs = append(*xs, v)
	}
	return *xs
}
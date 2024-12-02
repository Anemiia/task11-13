package main

func MaxDifference(arr []int) int {
	if len(arr) == 0 || len(arr) == 1 {
		return 0
	}

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[len(arr)-1] - arr[0]
}
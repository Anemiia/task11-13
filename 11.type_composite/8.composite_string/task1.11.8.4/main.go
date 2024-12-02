package main

func concatStrings(xs ...string) string {
	var str string
	for _, i := range xs {
		str += i
	}

	return str
}
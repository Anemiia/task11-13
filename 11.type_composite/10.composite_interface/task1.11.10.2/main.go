package main

import "fmt"

var Operate = func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}

var Concat = func(xs ...interface{}) interface{} {
	var str string
	for _, x := range xs {
		str += x.(string)
	}
	return str
}

var Sum = func(xs ...interface{}) interface{} {
	var sum float64
	var sumInt int
	for _, x := range xs {
		switch v := x.(type) {
		case int:
			sumInt += v
		case float64:
			sum += v
		}
	}
	if sumInt != 0 {
		return sumInt
	}
	return sum
}
func main() {
	fmt.Println(Operate(Concat, "Hello, ", "World!"))  // Вывод: "Hello, World!"
	fmt.Println(Operate(Sum, 1, 2, 3, 4, 5))           // Вывод: 15
	fmt.Println(Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5)) // Вывод: 16.5
}
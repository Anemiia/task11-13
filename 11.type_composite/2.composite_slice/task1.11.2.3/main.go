package main

import "fmt"

//Свойства XOR:
//( a \oplus a = 0 ) (любое число XOR с самим собой дает 0)
//( a \oplus 0 = a ) (любое число XOR с 0 остается тем же числом)
// если биты одинаковые, то результат 0, иначе 1

func bitwiseXOR(n, res int) int {
	return n ^ res
}

func findSingleNumber(numbers []int) int {
	res := 0
	for _, n := range numbers {
		res = bitwiseXOR(n, res)
	}
	return res
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	singleNumber := findSingleNumber(numbers)
	fmt.Println(singleNumber) // 5
}
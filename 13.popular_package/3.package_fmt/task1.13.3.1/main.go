package main

import "fmt"

func generateMathString(operands []int, operator string) string {
	uslovie := fmt.Sprintf("%d", operands[0])
	result := operands[0]

	for i := 1; i < len(operands); i++ {
		uslovie += fmt.Sprintf(" %s %d", operator, operands[i])
		result += operands[i]
	}

	return fmt.Sprintf("%s = %d", uslovie, result )
}

// Пример результата выполнения программы:
func main() {
	fmt.Println(generateMathString([]int{2, 4, 6}, "+")) // "2 + 4 + 6 = 12"
}
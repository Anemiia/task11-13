package main

func getType(i interface{}) string {
	switch i.(type) {
	case nil:
		return "Пустой интерфейс"
	case int:
		return "int"
	case string:
		return "string"
	case []int:
		return "[]int"
	default:
		return "unknown"
	}
}

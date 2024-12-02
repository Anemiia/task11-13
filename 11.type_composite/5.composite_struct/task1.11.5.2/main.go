package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"strings"
	"time"
)

// Пример кода
type Animal struct {
	Type   string
	Name   string
	Age    int
}

func getAnimals() []Animal {
	gofakeit.Seed(time.Now().UnixNano())

	srez := make([]Animal, 3)
	srez1 := []string{"Бублик", "Шарик", "Кузя"}
	for i := 0; i < 3; i++ {
		srez[i] = Animal {
			Type: gofakeit.AnimalType(),
			Name: srez1[i],
			Age: gofakeit.Number(1, 20),
		}
	}

	return srez
}

func preparePrint(animals []Animal) string {
	var stroka strings.Builder

	for _, animal := range animals {
		stroka.WriteString(fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d\n", animal.Type, animal.Name, animal.Age))
	}
	return stroka.String()
}

func main() {
	animals := getAnimals()
	result := preparePrint(animals)
	fmt.Println(result)
}
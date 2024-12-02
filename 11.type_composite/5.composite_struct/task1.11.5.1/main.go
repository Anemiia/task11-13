package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"strings"
	"time"
)

type User struct {
	Name string
	Age int
}

func getUsers() []User {
	gofakeit.Seed(time.Now().UnixNano())

	srez := make([]User, 10)

	for i := 0; i < 10; i++ {
		srez[i] = User {
			Name: gofakeit.Name(),
			Age: rand.Intn(43) + 18,
		}
	}
	return srez
}

func preparePrint(asd []User) string {
	var stroka strings.Builder

	for _, user := range asd {
		stroka.WriteString(fmt.Sprintf("Имя: %s, Возраст: %d\n", user.Name, user.Age))
	}

	return stroka.String()
}

func main() {
	users := getUsers()
	result := preparePrint(users)
	fmt.Println(result)
}
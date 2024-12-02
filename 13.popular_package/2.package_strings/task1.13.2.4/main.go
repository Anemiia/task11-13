package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func generateActivationKey() string {
	massiv := make([]string, 4)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			massiv[i] += string(rune(rand.Intn(26) + 65))
		}
	}

	return strings.Join(massiv, "-")

}

func main() {
	activationKey := generateActivationKey()
	fmt.Println(activationKey) // UQNI-NYSI-ZVYB-ZEFQ
}

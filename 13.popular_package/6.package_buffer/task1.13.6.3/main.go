package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func getDataString(b *bytes.Buffer) string {
	line, _ := bufio.NewReader(b).ReadString('\n')
	return line
}



func main() {
	// Create a new buffer
	buffer := bytes.NewBufferString("Hello, World!")

	// Call the getDataString function
	result := getDataString(buffer)

	// Check if the result matches the expected output
	expected := "Hello, World!"
	fmt.Println(result)
	if result != expected {
		panic(fmt.Sprintf("Expected %s, but got %s", expected, result))
	}
}

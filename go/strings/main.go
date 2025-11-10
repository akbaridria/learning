package main

import (
	"fmt"
	"strings"
)

func main() {
	var BuilderString strings.Builder

	SlicesString := []string{"Hello", "World", "From", "Go!"}

	JoinedString := strings.Join(SlicesString, " ")

	_, err := BuilderString.WriteString(JoinedString)
	if err != nil {
		fmt.Println("failed to insert a string to builder")
	}

	for i := range 1000 {
		_, err := BuilderString.WriteString(JoinedString)
		if err != nil {
			fmt.Println("failed to insert string to builder", i)
		}
	}

	result := BuilderString.String()

	if !strings.Contains(result, "Hello World") {
		fmt.Println("do not contain hellow world")
	}

	preview := result

	if len(preview) > 100 {
		preview = preview[:100]
	}
	fmt.Println(preview)
}

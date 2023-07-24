package main

import (
	"fmt"
	"os"
)

func reorderNames(first_name string, last_name string, language string) (first string, last string) {
	switch language {

	case "vn":
		{
			first = last_name
			last = first_name
		}
	case "us":
		{
			first = first_name
			last = last_name
		}
	default:
		{
			panic("Language is not supported")
		}
	}

	return
}

func main() {
	args := os.Args

	first_name, last_name, language := args[1], args[2], args[3]

	first, last := reorderNames(first_name, last_name, language)

	fmt.Printf("Output: %s %s\n", first, last)

}

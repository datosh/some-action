package main

import (
	"fmt"
	"os"
	"strconv"

	someaction "github.com/datosh/some-action"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("calculator <int> <int>")
		return
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%d", someaction.Add(a, b))
}

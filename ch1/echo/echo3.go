package echo

import (
	"fmt"
	"os"
	"strings"
)

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("------print the index and value once per line------")
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

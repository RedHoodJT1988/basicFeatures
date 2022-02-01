package main

import (
	"fmt"
	"sort"
)


func main() {
	names := [3]string {"Alice", "Charlie", "Bob"}
	secondPosition := &names[1]
	fmt.Println(*secondPosition)
	secondName := names[1]

	fmt.Println(secondName)
	sort.Strings(names[:])
	fmt.Println(*secondPosition)
	fmt.Println(secondName)

}

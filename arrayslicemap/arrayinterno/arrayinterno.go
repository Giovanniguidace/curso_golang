package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	// neste momento, a alteração irá refletir para os 2 slices, portanto, fazem acesso a um unico array.
	s1[0] = 7
	fmt.Println(s1, s2)

}

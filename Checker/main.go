package main

import "fmt"

type s struct {
	a int32
	b int32
	c int32
}

func main() {
	var ap s = s{1, 1, 1}

	ba := ap

	ba.a += 1

	fmt.Println(ap, ba)

}

package main

import "fmt"

type fbfunc func() int

func fb() fbfunc {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func main() {
	var fbvar fbfunc = fb()
	i := 0
	for {
		fmt.Println(fbvar())
		i++
		if i > 20 {
			break
		}
	}
}

package main

import "fmt"

func main() {
	demo("hello")
}

func demo(s string) (n int) {
	i := 0
	defer fmt.Println("i =",i)
	defer func ()  {
		fmt.Println("i = ",i, "s = ", s, "num = ", n)
	}()
	i+=1
	s += s
	n = len(s)
	return n
}
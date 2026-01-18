package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan bool)
	go func(){
		ch1 <- 1
	}()
	go odd(ch1, ch2, done)
	go even(ch1, ch2, done)

	<- done
	<- done
}

func odd(ch1, ch2 chan int, done chan bool) {
	for {
		select {
		case i := <-ch1: 
			fmt.Println("form odd = ", i)
			ch2 <- i+1
			if i == 99 {
				done <- true 
			}
		}
	}
}

func even(ch1, ch2 chan int, done chan bool) {
	for {
		select {
		case i := <- ch2:
			fmt.Println("form even = ",i)
			ch1 <- i+1
			if i == 100 {
				done <- true
			}
		}
	}
}


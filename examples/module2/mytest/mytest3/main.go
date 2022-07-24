package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string, 10)

	for i := 0; i < 2; i++ {
		go func(i int) {
			for {
				produce(c, i)
			}
		}(i)
	}

	for j := 0; j < 4; j++ {
		go func(j int) {
			for {
				consume(c, j)
			}
		}(j)
	}

	time.Sleep(time.Second * 10)
}

func consume(c <-chan string, i int) {
	s := fmt.Sprintf("Thread-%d", i)
	select {
	case result := <-c:
		fmt.Printf("%s consumer %s has consume %s one data\n", time.Now().String(), s, result)
		time.Sleep(time.Second)
	default:
		fmt.Printf("%s consumer %s no data process, wait \n", time.Now().String(), s)
		time.Sleep(time.Second * 2)
	}

}

func produce(c chan<- string, i int) {
	s := fmt.Sprintf("Thread-%d", i)
	fmt.Printf("%s producer %s has produce one data\n", time.Now().String(), s)
	c <- s
	time.Sleep(time.Second * 3)
}

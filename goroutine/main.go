package main

import (
	"fmt"
	"time"
)

func main() {
	// #3.2 Goroutine
	// https://nomadcoders.co/go-for-beginners/lectures/1521
	// Run sequentially
	sexyCount("dorbae", 1)
	sexyCount("Swim", 1)
	// Run
	go sexyCount("dorbae", 1)
	sexyCount("Swim", 1)

	// #3.3 Channels
	// https://nomadcoders.co/go-for-beginners/lectures/1522
	c := make(chan bool)
	people := [2]string{"dorbae", "swim"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println("Waiting for", i)
		result := <-c
		fmt.Println(result)
	}

}

func sexyCount(persion string, loop int) {
	for i := 0; i < loop; i++ {
		fmt.Println(persion, "is sexy", i)
		time.Sleep(time.Second * 1)
	}
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println("[DEBUG]", person, "is done")
	c <- true
}

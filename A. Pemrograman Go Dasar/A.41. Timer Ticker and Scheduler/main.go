package main

import (
	"fmt"
	"os"
	"time"
)

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}

func main() {
	fmt.Println("start")
	time.Sleep(time.Second * 1)
	fmt.Println("after 2 seconds")

	for i := 0; i < 3; i++ {
		fmt.Println("Hello !!")
		time.Sleep(1 * time.Second)
	}

	var timerr = time.NewTimer(1 * time.Second)
	fmt.Println("start")
	<-timerr.C
	fmt.Println("finish")

	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	fmt.Println("start")
	// <-ch
	fmt.Println("finish")

	<-time.After(4 * time.Second)
	fmt.Println("expired")

	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second) // wait for 10 seconds
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hello !!", t)
		}
	}

	var timeout = 5
	ch = make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("what is 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}
}

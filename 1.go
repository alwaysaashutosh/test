package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	t := 1

	_, _ = s.Every(t).Second().SingletonMode().Do(numbers)

	_, _ = s.Every(t).Second().SingletonMode().Do(alphabets)

	s.StartBlocking()
}

func numbers() {

	fmt.Println("Hello Numbers from 1 to 100")
	// for i := 1; i <= 10000; i++ {
	// 	// time.Sleep(2 * time.Millisecond)
	// 	fmt.Printf("%d ", i)
	// }

}

func alphabets() {
	fmt.Println("Hello Alphabets from A to z")
	// for i := 'a'; i <= 'e'; i++ {
	// 	// time.Sleep(3 * time.Millisecond)
	// 	fmt.Printf("%c ", i)
	// }
}

// Top 5 Golang Interview Questions with Code Examples
package main

import (
	"fmt"
	"time"
)

// 1. What are the key features of Golang?
// Definition: Golang, or Go, is a statically typed, compiled language known for its simplicity, concurrency support, fast compilation, and garbage collection.
func question1() {
	fmt.Println("Q1: Hello, Go!")
}

// 2. What is a goroutine?
// Definition: A goroutine is a lightweight thread managed by the Go runtime for concurrent programming.
func question2() {
	say := func(s string) {
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s)
		}
	}
	go say("world")
	say("hello")
}

// 3. What is a channel in Go?
// Definition: A channel is a communication mechanism that allows goroutines to communicate with each other and synchronize their execution.
func question3() {
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()
	msg := <-messages
	fmt.Println("Q3:", msg)
}

// 4. What are Go interfaces?
// Definition: Interfaces in Go are a way to specify the behavior of an object. An interface type is defined as a set of method signatures.
type Speaker interface {
	Speak() string
}

type Person struct{}

func (p Person) Speak() string {
	return "Hello!"
}

func question4() {
	var s Speaker = Person{}
	fmt.Println("Q4:", s.Speak())
}

// 5. What is the difference between buffered and unbuffered channels?
// Definition: Unbuffered channels block the sender until the receiver receives. Buffered channels allow sending without waiting until the buffer is full.
func question5() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Q5:", <-ch)
	fmt.Println("Q5:", <-ch)
}

func main() {
	// Uncomment the function call below to run a specific question:
	// question1()
	question2()
	// question3()
	// question4()
	// question5()
}

package main

import "fmt"

func say(texo string, c chan<- string) {
	c <- texo
}

func main() {

	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)
	fmt.Println(<-c)

}

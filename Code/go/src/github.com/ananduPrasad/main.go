package main

import "fmt"

func main() {
	Greeting()
	advanceGreeting("Anandu")
}

func Greeting() {
	fmt.Println("Hey there")
}

func advanceGreeting(name string) {
	fmt.Print("hello ,%s!\n", name)

}

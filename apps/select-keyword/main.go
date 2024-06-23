package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type message struct {
	text   string
	number int
}

func main() {
	msg := make(chan message)
	exit := make(chan bool)

	go readInput(msg, exit)
	for {
		select {
		case m := <-msg:
			validateInputAndGetNumber(&m)
			calculate(m)
			fmt.Print("˲Give me a number or exit: ")
		case <-exit:
			fmt.Println("˲Exiting program...")
			os.Exit(1)
		}
	}
}

func validateInputAndGetNumber(input *message) {
	number, err := strconv.Atoi(input.text)
	if err != nil {
		fmt.Printf("˲Invalid input: %v --- TRY AGAIN YOU CAN !\n", err)
		return
	}
	input.number = number
}

func calculate(input message) {
	if input.number != 0 {
		for i := 0; i <= 10; i++ {
			fmt.Printf("→ %11v x %v = %-7v\n", input.number, i, input.number*i)
		}
	}
}

func readInput(msg chan<- message, exit chan<- bool) {
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Print("˲Give me a number or exit: ")
	for scanner.Scan() {
		inputMsg := scanner.Text()
		if inputMsg == "exit" {
			exit <- true
		} else {
			msg <- message{text: inputMsg}
		}
	}
}

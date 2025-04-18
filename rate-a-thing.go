package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var rating int64
	for {
	  fmt.Println("Please select what to do: ")
	  fmt.Println("1.show rating of a thing")
	  fmt.Println("2.rate a thing")
	  fmt.Println("3.exit")
		choice:=userInput(reader,'\n')
		command:= returnInt(choice)
		
		if command == 1 {
			fmt.Println("rating of a thing: ",rating)
		} else if command == 2 {
			fmt.Printf("rate a thing: ",)
			rate := userInput(reader,'\n')
			rateInt := returnInt(rate)
			rating = rating + rateInt
		} else if command == 3 {
			fmt.Println("exiting program....")
			os.Exit(0)
		} else {
			fmt.Println("invalid choice")
		}
	}
}

func returnInt(s string) int64 {
	integer,_ :=strconv.ParseInt(strings.TrimSpace(s),10,64)
	return integer
}

func userInput(reader *bufio.Reader,breakChar byte) (string) {
	input,err:=reader.ReadString(breakChar)
	if err != nil {
		fmt.Println("error",err)
		os.Exit(1)
	}
	return input
}

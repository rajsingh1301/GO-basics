package main

import (
	"bufio"
	"fmt"
	"os"
)


func  main()  {
   // var input string
	// fmt.Println("enter the rating for the pizza")
	// fmt.Scanln(&input)
	// fmt.Printf("the rating you entered is %v", input)
	reader := bufio.NewReader(os.Stdin) // this is used to read the input from the user and it is more efficient than fmt.Scanln 
	fmt.Println("enter the rating for the pizza")

	input , err  := reader.ReadString('\n') // it is like try and catch block in other languages, it will return the input and the error if there is any error
	fmt.Printf("the rating you entered is %v and the error is %v", input, err)
	fmt.Printf("the type of the input is %T", input)

}

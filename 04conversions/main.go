package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
  fmt.Println("welcome to pizza place")
    fmt.Println("rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin);
	 
	input, _ := reader.ReadString('\n');
	fmt.Printf("the rating you entered is %v", input)
	  numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // convert string to float64
	  if err != nil {
		  fmt.Println("error parsing float")
		  return
	  }else {
	  fmt.Printf("the rating you entered is %v", numrating + 1)
	  }


}

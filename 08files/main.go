package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
func main(){
      fmt.Println("enter the file:")
	  content := "this is the content of the file"

	  //create a file
	  file , err := os.Create("test.txt")
	  if(err != nil){
		panic(err)
	  }
	  //write to the file 
	  length , err := io.WriteString(file , content)
	  if(err!=nil){
		panic(err)
	  }
	  fmt.Printf("length is %v", length)

	  displayContent("test.txt")
	}

	func displayContent(fileName string){
		// read the file 
		data , err := ioutil.ReadFile(fileName)
		if(err != nil){
			panic(err)
		}
		fmt.Print(string(data))
	}
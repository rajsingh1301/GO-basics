package main

import (
	"fmt"

	"io/ioutil"
	"net/http"
)

const url = "https://api.ipify.org"
func main(){

	// make a web request to the url and print the response
fmt.Println(" lco web request")

resp , err := http.Get(url)
if err != nil {
	panic(err)	

}
fmt.Printf("%T",resp)
 defer resp.Body.Close() // close the response body to free up resources

dataBytes , err := ioutil.ReadAll(resp.Body)

if(err != nil){
	panic(err)		

}
content := string(dataBytes)
fmt.Println(content)	
}
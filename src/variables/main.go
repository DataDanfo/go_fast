package main

import "fmt"

// comments
// var variablename type = value

func main() {
    var a int32 = 56
	var b int64 
	var c *int32 
	c = &a

	fmt.Println(a, b,c)
	var myName string = "Jengo"
	fmt.Println(myName)
}
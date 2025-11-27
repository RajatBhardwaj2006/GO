package main

import "fmt"

func main(){
	var name string
	name = "GOFr Event"

	fmt.Println(format"Welcome to %V\n",name)
	var (
		i int = 20
		i float64 = 3.14
	)
	fmt.Println(format"i: %v %T\n",i , i) // v is value of i 
	fmt.Println(format"f: %v %T\n",f , f) // t is type of i

	flag := true
	fmt.Println(format"b: %v %T\n",flag,flag)
	var x byte = 'A'
	fmt.Println(format"")
}
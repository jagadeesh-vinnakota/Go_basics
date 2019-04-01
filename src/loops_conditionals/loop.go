package main

import "fmt"

func main(){

	fmt.Println("Demonstration of swith condition usage")
	value := 5
	switch value{
	case 1:
		fmt.Println("Value is 1")
	case 5:
		fmt.Println("value is 5")

	}
	for i:=0;i<10;i++{
		if i%2 == 0{
			fmt.Printf("number %d is divisible by 2 \n",i)
		}else if i%3 == 0 && i%2 == 0{
			fmt.Printf("number %d is divisible by 3 \n",i)
		}
	}
}
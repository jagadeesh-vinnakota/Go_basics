package main

import (
	"fmt"
	"packages/test_pkg"
)

func init(){
	fmt.Println("Initializatin in driver program")
}
func main(){
	area, perimeter := test_pkg.Calculate_Result(4)
	fmt.Printf("Area %d, perimeter %d",area, perimeter)
}
package test_pkg
import "fmt"
func init(){
fmt.Println("Package Initialization")
}
func Calculate_Result(side int) (int,int){
	area := side * side
	permiter := 4 * side
	return area, permiter
}

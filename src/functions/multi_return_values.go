package main
import "fmt"

func calculate_result(side int) (int,int){
	area := side * side
	permiter := 4 * side
	return area, permiter
}

func main(){

	area, perimeter := calculate_result(4)
	fmt.Printf("result of square: area %d, perimeter %d", area, perimeter)
}
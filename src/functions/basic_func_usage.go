package main
import "fmt"
func main(){
	fmt.Println("Bill is: ",calculateBill(10,5))
}
func calculateBill(price int, count int) int{
	return price*count
}
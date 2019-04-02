package main



import "fmt"



func main() {  

    var age int // variable declaration

    var weight = 155

    var caption,hobby int = 3,4

    fmt.Println("my age is", age)

    fmt.Println("weight: ",weight)

    fmt.Println("caption and hobby is: ",caption," hobby is ",hobby)



    //declare varibales of different types at same time

    var(

        name = "jaga"

        fitness = 10

    )

    fmt.Println("values are: ",name,fitness)



    //short hand notation

    short1,short2 := 5,6

    fmt.Println("short hand notation: ",short1,short2)

    fmt.Printf("type of short1 is %T \n", short1)





    //constants

    const JERSEY_NUMBER = 24

    fmt.Println("my jersey number: ",JERSEY_NUMBER)

}
package main

import "fmt"

func display_oned_array(scores []int){
		for index, value := range scores{
		fmt.Printf("%d index value is %d \n",index,value)
	}
}

func display_twod_array(scores [3][2]int){

	for outer_index,value := range scores{
		for inner_index,val := range value{
			fmt.Printf("Value at [%d][%d] is %d \n",outer_index,inner_index,val)
		}
	}
}
func main(){

	fmt.Println("Usage of arrays demonstration")
	
	//Basice array declaration
	var names[3] int
	fmt.Println("Array declaration with default vaues")
	display_oned_array(names[:])

	//Array declaration with values
	fmt.Println("Array declaration with assigned values")
	var values[3] int
	values[0] = 24
	values[1] = 99
	values[2] = 1
	display_oned_array(values[:])

	//Short hand Array declaration along with initialization 
	fmt.Println("Short hand array declation")
	scores := [3]int{99, 23, 11}
	display_oned_array(scores[:])



	//Multi dimensional arrays
	fmt.Println("Values in multi dimensional arrays are: ")
	multi_array := [3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	display_twod_array(multi_array)

	//Slice creation
	fmt.Println("Short hand slice declation")
	scores_array := [5]int{99, 23, 11,22,33}
	slice_array := scores_array[:]
	display_oned_array(slice_array)
}
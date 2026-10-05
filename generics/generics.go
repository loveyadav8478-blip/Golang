package main

import "fmt"


func printSlice[T comparable](nums []T){
	for _, v := range nums{
		fmt.Print(v," ")
	}
	println()
}

// func printStringSlice(nums []string){
// 	for _, v := range nums{
// 		fmt.Println(v)
// 	}
// }


type stack[T any] struct{
	element[]T
}
func main(){
	printSlice([]int{1,2,3,4})
	printSlice([]string{"Hi","this","side","GOLANG"})
	printSlice([]bool{false,true,false,false})
	myStack := stack[string]{
		element: []string{"12","3","46","6"},
	}
	fmt.Println(myStack)
}
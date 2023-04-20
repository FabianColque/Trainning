package main

import "fmt"

func main(){
	channel := make(chan int)

	//var list []int
	//go setList(&list)
	go setList(channel)

	for v := range channel{
		fmt.Println(v)
	}
}

func setList(channel chan int){
	for i := 0; i < 100; i++{
		channel <- i 
	}
	close(channel)
}

/*func setList(list *[]int){
	for i := 0; i < 100; i++{
		(*list) = append((*list), i); 
	}
}*/
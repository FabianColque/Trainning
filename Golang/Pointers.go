package main

import ("fmt")

func main(){
	x := 5;
	y := &x
	*y = 10;
	fmt.Println(x, *y);
	fmt.Println(&x, &y)
	PrintValues(&x, y)
}

func PrintValues(x *int, y *int){
	fmt.Println("Print values")
	fmt.Println(x, y)
	//fmt.Println(&x, &y)
}
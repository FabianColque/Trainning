package main

import ("fmt")

func main(){
	list1 := []int{1,2}
	sum := (list1[0] + list1[1])
	fmt.Println(sum)	

	list2 := []int{2,8,3,10,5,4,7,9,1,6}
	listfirst5 := list2[:5]
	listlast5  := list2[5:]
	sumfirst5 := 0
	for i := 0; i < len(listfirst5); i++{
		sumfirst5 += listfirst5[i]
	}

	sumlast5 := 0
	for i := 0; i < len(listlast5); i++{
		sumlast5 += listlast5[i]
	}

	fmt.Println(sumfirst5)
	fmt.Println(sumlast5)
}
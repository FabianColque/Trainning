package main

import "fmt"

func main(){
	slice := []int {5, 1, 2, 3}
	slice2 := []string{"a", "e", "f", "b"}

	newInts := reverse(slice)
	fmt.Println(newInts)

	newStrings := reverse(slice2)
	fmt.Println(newStrings)
}

type constraintCustom interface{
	int | string
}

func reverse[T constraintCustom](slice []T) []T{
	newInts := make([]T, len(slice))

	nLen := len(slice)-1
	for i := 0; i < nLen+1; i++ {
		newInts[nLen-i] = slice[i]
	}
	return newInts
}

/*func reverse[T int | string](slice []T) []T{
	newInts := make([]T, len(slice))

	nLen := len(slice)-1
	for i := 0; i < nLen+1; i++ {
		newInts[nLen-i] = slice[i]
	}
	return newInts
}*/
package main

import "fmt"

func main(){
	var lista []interface{}

	lista = append(lista, 2)
	lista = append(lista, true)
	lista = append(lista, "Fabian")
	lista = append(lista, 12.4)

	fmt.Println(lista)

}
package main

import ("fmt"
        "math")

func main(){
	lista := []int{1,2,3,4,4}
	fmt.Print("MyList: ")
	for i := 0; i < len(lista); i++{
		fmt.Print(" ", lista[i])
	}
	fmt.Println()

	lista = append(lista, 12)
	fmt.Print("MyList: ")
	for i := 0; i < len(lista); i++{
		fmt.Print(" ", lista[i])
	}
	fmt.Println()

	lista2 := make([]int, 1)
	fmt.Printf("%T \n", lista2)

	lista3 := []int{2,4,6,8,10}
	sublista := make([]int, 0)
	for i := 0; i < len(lista3); i++ {
		if math.Mod(float64(i), 2.0) == 0.0 {
			sublista = append(sublista, lista3[i])	
		}
	}

	fmt.Print("MyList3: ")
	for i := 0; i < len(sublista); i++{
		fmt.Print(" ", sublista[i])
	}
	fmt.Println()

	var listaAll = []int{1,2,3,4,5,6,7,8,9}
	firstlist := listaAll[:3]
	terceritem := listaAll[4:]
	lastItem := listaAll[len(listaAll)-1:]

	fmt.Println(firstlist)
	fmt.Println(terceritem)
	fmt.Println(lastItem)


	mymap := map[string]int{"sp": 1000, "cg":900}
	mymap["sp"] = 2000
	mymap["rj"] = 1212
	fmt.Println(mymap)

	value, exist := mymap["sp"]
	fmt.Println(value, exist)

	value, exist = mymap["sp2"]
	fmt.Println(value, exist)

	for key, val := range mymap{
		fmt.Println(key, val)
	}

	delete(mymap, "sp")
	fmt.Println(mymap)
}
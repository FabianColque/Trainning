//Podemos utilizar https://go.dev/play/

package main

import ("fmt"
        "reflect")

func main(){
	texto := "Teste"
	//texto = 10
	//number := 12
	//flag := true
	//float1 := 12.3
	fmt.Println("Hello " + texto)

	num1 := 10
	num2 := 20
	result := num1 * num2
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))

	//Constantes
	//const tax = 10
	//tax = 2

	var numero int8 = 2
	fmt.Println(numero)

	var entero int = -1
	var ocho int8
	var nose uint8
	ocho = int8(entero)
	nose = uint8(entero)
	//nose = -1

	fmt.Println(ocho)
	fmt.Println(nose)

	

}
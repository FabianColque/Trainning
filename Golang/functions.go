package main

import ("fmt")

func main(){
	teste()
	fmt.Println(soma(234,244))
	fmt.Println(operations(234,244))
	fmt.Println(operationsNamed(234,244))
}

func init(){
	fmt.Println("apareceu?")
}

func teste(){
	fmt.Println("Holi")
}

func soma(num1 int, num2 int) int{
	result := num1 + num2
	return result
}

func operations(num1 int, num2 int) (int, int, int, int) {
	soma := num1 + num2
	resta := num1 - num2
	multi := num1 * num2
	divi := num1 / num2
	return soma, resta, multi, divi
}

//named
func operationsNamed(num1 int, num2 int) (soma int, resta int, multi int, divi int) {
	soma = num1 + num2
	resta = num1 - num2
	multi = num1 * num2
	divi = num1 / num2
	return
}
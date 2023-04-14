package main

import ("fmt")

func main(){
	//IF
	salario := 2000.00
	var bonus float64

	bonus = salario

	if salario < 100 {
		bonus = (salario + 100.0)
	}

	fmt.Println("Salario: ", bonus)
	
	salario = 1900.0
	if salario < 1000 {
		bonus = (salario + 100.0)
	}else{
		bonus = (salario - 100.0)
	}

	fmt.Println("Salario: ", bonus)

	var flag bool = false
	flag = bonus > salario
	fmt.Println("misio: ", flag)

	if !flag{
		bonus += 2000
	}
	flag = bonus > salario
	fmt.Println("misio?: ", flag)

	//loop
	for i := 0; i < 6; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//
	strText := "Fabian"
	lenght := len(strText)
	i := 0
	for i < lenght {
		fmt.Print(string(strText[i]))
		i++
	}
	fmt.Println()
}
package main

import ("fmt"
		"mymodgo/model"
		"time")


func main(){
	fmt.Println("Hello")

	myaddress := model.Address{
		Street : "Verissimo",
		Number : 12,
		City : "Porto",
	}

	myself := model.Person{
		Name : "Fabian",
		Address : myaddress,
		Birthday : time.Date(1992, 02, 19, 0, 0, 0, 0, time.Local),
	}

	fmt.Println(myself)
	fmt.Println("Age:", model.CalcAge(myself))
	myself.CalculateAge()
	fmt.Println("AgePerson:", myself.Age)


	automovil := model.Automovel{
		Ano : 2022,
		Placa : "XP",
		Modelo : "CG",
	}

	moto := model.Moto{
		Automovel : automovil,
		Cilindradas : 125,
	}

	fmt.Println(moto)
}
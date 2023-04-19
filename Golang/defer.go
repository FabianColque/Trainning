package main

import ("os"
		"fmt")

func ShowText(){
	fmt.Println("Closing...")
}

func main(){
	file, err := os.Create("./settings.txt")
	defer file.Close() //Significa que sera a ultima coisa a executar
	defer ShowText()

	if err != nil {
		panic(err)
	}

	file.Write([]byte("teste"))	
}
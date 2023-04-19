package main

import ("os"
		"fmt")

func ReadFile(){
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("Recover")
		}
	}()

	_, err := os.Open("./settings2.txt")
	if err != nil {
		panic("FileNotFound")
	}	
}

func main(){
	
	ReadFile()
	fmt.Println("Closing...")
}
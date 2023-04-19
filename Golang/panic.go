package main

import "os"

func main(){
	//panic("error")
	_, err := os.Open("f:/settings.txt")
	if err != nil {
		panic(err)
	}
}
package main

import ("fmt"
		"time"
		"sync")

func main(){
	var wg sync.WaitGroup
	wg.Add(3)

	go callDatabase(&wg)
	go callAPI(&wg)
	go callInternal(&wg)

	//time.Sleep(2*time.Second)
	wg.Wait()
}

func callDatabase(wg *sync.WaitGroup){
	time.Sleep(1*time.Second)
	fmt.Println("Finalizado database")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup){
	time.Sleep(1*time.Second)
	fmt.Println("Finalizado API")
	wg.Done()
}

func callInternal(wg *sync.WaitGroup){
	time.Sleep(1*time.Second)
	fmt.Println("Finalizado internal")
	wg.Done()
}
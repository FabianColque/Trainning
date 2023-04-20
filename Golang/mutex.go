package main

import ("fmt"
		"time"
		"sync")

func main(){
	//var wg sync.WaitGroup

	var m sync.Mutex
	i := 0
	for x := 0; x < 10000; x++{
		//go ChangeNumber(&i, x)		
		go func(){
			m.Lock()
			i++;
			m.Unlock()
		}()		
	}

	time.Sleep(time.Second * 2)

	fmt.Println(i)
	//wg.Wait()
}

/*func ChangeNumber(i *int, newNumber int){
	*i = newNumber
}*/
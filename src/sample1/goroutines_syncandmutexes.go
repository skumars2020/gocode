package main

import (
	"fmt"
	"sync"
	"time"
)

//wait method is sync
func main() {

	//Muiteaccess
	mut := &sync.Mutex{}
	//mut.Lock and mut.Unlock

	wg := &sync.WaitGroup{}
	//Below one was lamda funcation
	wg.Add(2)
	go func() { //Delcaration of lamda funcation
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			fmt.Println("Heavy")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 2)
			fmt.Println("Super - Heavy")
		}
		wg.Done()
	}()

	fmt.Println("Fin")
	wg.Wait()
	fmt.Println("Final Exe")
}

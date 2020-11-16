package counter

import (
	"fmt"
	"time"
)

func timer(d time.Duration) <-chan int {
	counter := make(chan int)
	go func() {
		time.Sleep(d)
		counter <- 1
	}()
	return counter
}

// Count ...
func Count() {
	for i := 9; i >= 0; i-- {
		counter := timer(1 * time.Second)
		<-counter
		//fmt.Printf("%d ", val+i)
		fmt.Print(".")
	}
	fmt.Println()
}

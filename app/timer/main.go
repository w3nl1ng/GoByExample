package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2*time.Second)
	<-timer1.C
	fmt.Println("timer1 fired")

	timer2 := time.NewTimer(2*time.Second)
	go func(){
		<-timer2.C
		fmt.Println("time2 fired")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stoped")
	}

	time.Sleep(3*time.Second)

}
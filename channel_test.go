package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T){
	chanel := make(chan string)
	
	go func ()  {
		time.Sleep(3 * time.Second)
		chanel <- "Ardi Firmansyah"
		fmt.Println("Chanel Finisihed")
	}()

	data := <- chanel
	close(chanel)
	fmt.Println(data)
}
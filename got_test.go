package main

import (
	"fmt"
	"testing"
	"time"
)

func DisplayN(s int) {
	fmt.Println("Tampil", s)
}

func TestManyGoRoutines(t *testing.T){
	for i:= 0; i <= 10000; i++{
		go DisplayN(i)
	}
	time.Sleep(1 * time.Second)
}
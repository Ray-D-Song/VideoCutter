package main

import (
	"VideoCutter/handler"
	"VideoCutter/helper"
	"fmt"
	"os"
	"sync"
)

func mainTask(wg *sync.WaitGroup) {
	if len(os.Args) < 4 {
		panic("Insufficient number of parameters")
	}
	if os.Args[2] != "-d" && os.Args[2] != "-r" {
		panic("Incorrect operation parameter, can only be -d or -r")
	}
	switch os.Args[2] {
	case "-r":
		fmt.Println(os.Args)
		handler.Retrieve(os.Args[3], os.Args[4], os.Args[1], os.Args[5])
	case "-d":
		fmt.Println(os.Args[2])
		handler.Divider(wg, os.Args[3], os.Args[1], os.Args[4])
	}
}

func main() {
	var wg sync.WaitGroup
	helper.TryCatch(func() {
		mainTask(&wg)
	}, func(R any) {
		fmt.Println(R)
	})
	wg.Wait()
}

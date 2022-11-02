package main

import (
	"fmt"
	"time"
)

func main2() {
	timeNow := time.Now()
	timeFormat := timeNow.Format("2006-01-02 15:04:05")

	fmt.Println(fmt.Sprintf("%s", timeFormat))
}

package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().Format(time.RFC3339), " - Hello Go!")
		time.Sleep(5 * time.Second)
	}
}

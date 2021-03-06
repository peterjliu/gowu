package main

import (
	"fmt"
	"github.com/peterjliu/gowu"
	"os"
)

func main() {
	key := os.Getenv("WU_API_KEY")
	fmt.Printf("Using api key: %s\n", key)
	c := gowu.NewClient(key)
	cond, err := c.GetConditions("seattle", "wa")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Temperature is in San Francisco is %g F\n", cond.CurrentObservation.TempF)
}

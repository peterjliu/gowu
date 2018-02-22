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
	fore, err := c.GetForecast("portland", "or")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Today high/low in in Portland is: %s/%s F\n", fore.Simpleforecast.Forecastday[0].High.Fahrenheit, fore.Simpleforecast.Forecastday[0].Low.Fahrenheit)
	fmt.Printf("Tomorrow high/low in in Portland is: %s/%s F\n", fore.Simpleforecast.Forecastday[1].High.Fahrenheit, fore.Simpleforecast.Forecastday[1].Low.Fahrenheit)
}

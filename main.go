package main

import (
	"fmt"
	"main/config"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	// Use "%s" to print the string value of ENVIRONMENT
	fmt.Println(conf.ENVIRONMENT)
	fmt.Println(conf.DEBUG)
	for _, item := range conf.SERVICES {
		fmt.Println(item)
	}
}

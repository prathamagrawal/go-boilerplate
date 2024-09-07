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
	fmt.Println(conf.ENVIRONMENT)
	fmt.Println(conf.DEBUG)
	for _, item := range conf.SERVICES {
		fmt.Println(item)
	}
}

package main

import (
	"fmt"
	"github.com/spf13/viper"
	"main/config"
)

func main() {
	config.LoadConfig()
	fmt.Println(viper.GetBool("DEBUG"))
}

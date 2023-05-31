package main

import (
	"fmt"
	"githun.com/captainstdin/airenche-temp/service"
)

func main() {
	err := service.LoadEnv()
	if err != nil {
		fmt.Println(err)
		return
	}

}

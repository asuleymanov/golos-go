package main

import (
	"fmt"

	"github.com/asuleymanov/golos-go"
)

func main() {
	cls, _ := golos.NewClient("wss://golos.solox.world/ws")
	defer cls.Close()

	ans, err := cls.API.GetOpsInBlock(35137873, true)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Answer :")
		for k, v := range ans {
			fmt.Printf("OP : %d -> %+v \n", k, v)
		}
	}
}

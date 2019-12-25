package basic

import (
	"fmt"
	"net"
)

func GetLocalIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		fmt.Println(addr.String())
	}
}

package main

import (
	"bytes"
	"fmt"
	"net"
)

func GetInterface() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func main() {
	interfaces, err := net.Interfaces()
	as, err := GetInterface()
	if err == nil {

		for _, i := range interfaces {
			fmt.Printf("Name: %v \n", i.Name)
			fmt.Println("type:", i.Flags.String())
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				for _, a := range as {

					fmt.Printf("MAC Addr: %v \n", a)

				}
				fmt.Println("============================================")
			}
		}
	}
}

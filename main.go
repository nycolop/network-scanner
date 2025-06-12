package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	interfaces, err := net.Interfaces()

	fmt.Println("Running network scan...")
	fmt.Println()

	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%-2s %-30s %-20s %-20s", "#", "Name", "Ip", "Mac"))
	fmt.Println(strings.Repeat("-", 70))

	for i, el := range interfaces {
		addrs, _ := el.Addrs()

		fmt.Println(fmt.Sprintf("%-2d %-30s %-20s %-20s", i+1, el.Name, addrs[1], el.HardwareAddr))
	}

	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please, put the id of the network to scan: ")
	input, err := reader.ReadString('\n')

	fmt.Println(input)
}

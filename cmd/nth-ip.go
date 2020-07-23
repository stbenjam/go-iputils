// Generates the nth IP in a subnet
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/openshift/installer/pkg/ipnet"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s NETWORK_CIDR INDEX\n", os.Args[0])
		os.Exit(1)

	}

	index, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("invalid index %q: %s\n", os.Args[2], err.Error())
		os.Exit(1)

	}

	ipnet, err := ipnet.ParseCIDR(os.Args[1])
	if err != nil {
		fmt.Printf("invalid CIDR %q: %s\n", os.Args[1], err.Error())
		os.Exit(1)

	}

	ip, err := cidr.Host(&ipnet.IPNet, index)
	if err != nil {
		fmt.Printf("could not generate ip: %s\n", err.Error())
		os.Exit(1)

	}

	fmt.Println(ip)
}

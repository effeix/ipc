package ipc

import (
	"fmt"
	"log"
	"math/big"
	"net"

	"github.com/rodaine/table"
)

func CountIPs(cidr string) (result int) {
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatalf("Invalid CIDR notation: %s", cidr)
		return 0
	}

	maskBits, _ := network.Mask.Size()
	maxMaskBits := 128
	if len(network.IP) == net.IPv4len {
		maxMaskBits = 32
	}

	return 1 << (maxMaskBits - maskBits)
}

func GenerateIPs(cidr string) []net.IP {
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatalf("Invalid CIDR notation: %s", cidr)
		return nil
	}

	isIPv4Network := len(network.IP) == net.IPv4len

	numIPs := CountIPs(cidr)

	ips := make([]net.IP, numIPs)
	newIPInt := big.NewInt(0)
	newIPInt.SetBytes(network.IP.To16())
	var newIP net.IP
	for i := int64(0); i < int64(numIPs); i++ {
		bytes := newIPInt.Bytes()
		
		if isIPv4Network { 
			newIP = net.IPv4(
				bytes[2],
				bytes[3],
				bytes[4],
				bytes[5],	
			).To4()
		} else {
			newIP = net.IP(bytes)
		}

		ips[i] = newIP
		newIPInt.Add(newIPInt, big.NewInt(1))
	}

	return ips
}

func ExplainNetwork(cidr string) error {
	// maxMaskSizeIPV4 := 32
	// maxMaskSizeIPV6 := 128

	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}
	maskSize, _ := network.Mask.Size()

	isIPv4 := len(network.IP) == net.IPv4len

	// 1. Explanation of Mask Part
	maxMaskSize := 128
	if isIPv4 {
		maxMaskSize = 32
	}

	fmt.Printf("CIDR: %s\n", cidr)
	fmt.Printf("The mask part '/%d' indicates that the first %d bits are fixed as the network part and the remaining %d bits are used for host addresses within the network.\n", maskSize, maskSize, maxMaskSize-maskSize)
	fmt.Printf("So, in the given CIDR, the IP address '%s' represents the network, and any address where the first %d bits match those of '%s' is part of this network.\n\n", network.IP.String(), maskSize, network.IP.String())

	// Initialize table with header
	tbl := table.New("Type", "IP", "Binary Representation")

	tbl.AddRow(
		"Network Address",
		network.IP.String(),
		ToBinary(network.IP),
	)

	tbl.AddRow(
		"Subnet Mask",
		net.IP(network.Mask).String(),
		ToBinary(network.Mask),
	)

	// Example IP - If IPv4, increment the last octet. For IPv6, we may increment the last 16-bit block for simplicity.
	exampleIP := net.IP(make([]byte, len(network.IP))) // Create a new IP address with the same length (IPv4 or IPv6)
	copy(exampleIP, network.IP)                        // Copy the original IP into the example IP
	if isIPv4 {
		exampleIP[3]++
	} else { // IPv6
		exampleIP[15]++
	}
	exampleIPBinary := ToBinary(exampleIP)
	tbl.AddRow(
		"Example Host IP",
		exampleIP.String(),
		exampleIPBinary,
	)

	tbl.Print()

	fmt.Printf(
		"\nThe IP '%s' is just an example IP within the network. All IPs where the first %d bits match '%s' are valid hosts in the network.\n",
		exampleIP.String(),
		maskSize,
		ToBinary(network.IP)[:maskSize],
	)

	fmt.Printf(
		"This gives us a total of %d possible host addresses in the network.\n",
		CountIPs(cidr),
	)

	return nil
}
package ipc

import (
	"fmt"
	"net"
)

const (
	MAX_IPv4_MASK = 32
	MAX_IPv6_MASK = 128
)

func ToBinary(ip []byte) string {
	binStr := ""

	for _, b := range ip {
		binStr += fmt.Sprintf("%08b", b)
	}

	return binStr
}


func IsValidCIDR(cidr string) bool {
	_, _, err := net.ParseCIDR(cidr)
	return err == nil
}

func GetPrintable(ips []net.IP, binary bool, oneline bool) []string {
	var ipStr string
	var printable []string
	
	for index, ip := range ips {
		ipStr = ip.String()
		if binary {
			ipStr = ToBinary(ip)
		}
		
		result := ipStr

		trailingComma := index < len(ips) - 1
		if oneline && trailingComma {
			result += ","
		} else if !oneline {
			result += "\n"	
		}

		printable = append(printable, result)
	}

	return printable
}
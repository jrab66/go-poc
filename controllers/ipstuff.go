package controllers

import (
	"fmt"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetClientIP(c *gin.Context) string {
	// Check X-Real-IP or X-Forwarded-For headers for proxy scenarios, ipv4 or ipv6
	if ip := c.Request.Header.Get("X-Real-IP"); ip != "" {
		return ip
	} else if forwarded := c.Request.Header.Get("X-Forwarded-For"); forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}
	//  edge case : Fallback to the default ClientIP method , if localhost ipv6 default
	if strings.HasPrefix(c.ClientIP(), "::1") {
		// Convert IPv6 loopback to IPv4 loopback (localhost)
		return "127.0.0.1"
	}
	return c.ClientIP()
}

// func reverseIP(ipString string) string {
// 	ip := net.ParseIP(ipString)
// 	if ip == nil {
// 		fmt.Printf("Invalid IP address: %s\n", ipString)
// 		return ""
// 	}

// 	// Reverse the IP address
// 	reversedIP := reverseString(ip.String())

// 	return reversedIP
// }

// func reverseString(s string) string {
// 	// Split the string into individual parts
// 	parts := strings.Split(s, ".")

// 	// Reverse the order of parts
// 	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
// 		parts[i], parts[j] = parts[j], parts[i]
// 	}

// 	// Join the reversed parts into a string
// 	reversedString := strings.Join(parts, ".")

//		return reversedString
//	}
func ReverseIP(ipString string) string {
	ip := net.ParseIP(ipString)
	if ip == nil {
		fmt.Printf("Invalid IP address: %s\n", ipString)
		return ""
	}

	// Check if the IP is IPv6
	isIPv6 := strings.Contains(ipString, ":")

	if isIPv6 {
		// Reverse the IPv6 address
		reversedIP := reverseIPv6(ip)
		return reversedIP
	}

	// Reverse the IPv4 address
	reversedIP := reverseIPv4(ip)
	return reversedIP
}

func reverseIPv4(ip net.IP) string {
	// Split the IPv4 string into individual parts
	parts := strings.Split(ip.String(), ".")

	// Reverse the order of parts
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	// Join the reversed parts into a string
	reversedString := strings.Join(parts, ".")
	return reversedString
}

func reverseIPv6(ip net.IP) string {
	// Split the IPv6 string into individual parts
	parts := strings.Split(ip.String(), ":")

	// Reverse the order of parts
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	// Join the reversed parts into a string
	reversedString := strings.Join(parts, ":")
	return reversedString
}

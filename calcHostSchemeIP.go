package main

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

// Calculate subnet details from CIDR notation
func calculateSubnetInfo(cidr string) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Invalid CIDR format:", err)
		return
	}

	ipAddr := ip.To4()
	maskSize, _ := ipnet.Mask.Size()

	networkAddr := ip.Mask(ipnet.Mask)

	broadcastAddr := make(net.IP, len(networkAddr))
	for i := 0; i < len(networkAddr); i++ {
		broadcastAddr[i] = networkAddr[i] | ^ipnet.Mask[i]
	}

	firstIP := make(net.IP, len(networkAddr))
	copy(firstIP, networkAddr)
	firstIP[3]++

	lastIP := make(net.IP, len(broadcastAddr))
	copy(lastIP, broadcastAddr)
	lastIP[3]--

	hostCount := int(math.Pow(2, float64(32-maskSize))) - 2

	// Convert subnet mask to dotted decimal format
	subnetMask := net.IP(ipnet.Mask).String()

	fmt.Println("\n----- Subnet Details -----")
	fmt.Println("IP Address:        ", ipAddr)
	fmt.Println("Subnet Mask:       ", subnetMask)
	fmt.Println("CIDR Notation:      /" + strconv.Itoa(maskSize))
	fmt.Println("First Usable IP:   ", firstIP)
	fmt.Println("Last Usable IP:    ", lastIP)
	fmt.Println("Total Hosts:       ", hostCount)

	printPossibleNetworks(ipnet, maskSize)
}

// Generate IP scheme based on required hosts
func generateIPScheme(baseIP string, requiredHosts int) {
	ip := net.ParseIP(baseIP)
	if ip == nil {
		fmt.Println("Invalid IP address format.")
		return
	}

	maskSize := calculateSubnetMask(requiredHosts)
	if maskSize == -1 {
		fmt.Println("Unable to calculate subnet mask for the given number of hosts.")
		return
	}

	ipNet := &net.IPNet{
		IP:   ip.Mask(net.CIDRMask(maskSize, 32)),
		Mask: net.CIDRMask(maskSize, 32),
	}

	broadcastAddr := make(net.IP, len(ipNet.IP))
	for i := 0; i < len(ipNet.IP); i++ {
		broadcastAddr[i] = ipNet.IP[i] | ^ipNet.Mask[i]
	}

	firstIP := make(net.IP, len(ipNet.IP))
	copy(firstIP, ipNet.IP)
	firstIP[3]++

	lastIP := make(net.IP, len(broadcastAddr))
	copy(lastIP, broadcastAddr)
	lastIP[3]--

	totalHosts := int(math.Pow(2, float64(32-maskSize))) - 2

	fmt.Println("\n----- IP Scheme -----")
	fmt.Println("IP Address:        ", baseIP)
	fmt.Println("Broadcast Address: ", broadcastAddr)
	fmt.Println("Subnet Mask:       ", net.IP(ipNet.Mask).String())
	fmt.Println("CIDR Notation:     /" + strconv.Itoa(maskSize))
	fmt.Println("First Usable IP:   ", firstIP)
	fmt.Println("Last Usable IP:    ", lastIP)
	fmt.Println("Total Hosts:       ", totalHosts)

	printPossibleNetworks(ipNet, maskSize)
}

// Calculate subnet mask based on required hosts
func calculateSubnetMask(requiredHosts int) int {
	for maskSize := 30; maskSize >= 0; maskSize-- {
		if int(math.Pow(2, float64(32-maskSize)))-2 >= requiredHosts {
			return maskSize
		}
	}
	return -1
}

// Print possible networks based on CIDR
func printPossibleNetworks(ipnet *net.IPNet, maskSize int) {
	subnetCount := int(math.Pow(2, float64(maskSize-24))) // Assuming /24 as base
	fmt.Println("\n----- All Possible /" + strconv.Itoa(maskSize) + " Networks for " + ipnet.IP.String() + " -----")
	fmt.Printf("%-20s %-30s %-20s\n", "Network Address", "Usable Host Range", "Broadcast Address")

	for i := 0; i < subnetCount; i++ {
		networkAddr := make(net.IP, len(ipnet.IP))
		copy(networkAddr, ipnet.IP)
		networkAddr[3] += byte(i * int(math.Pow(2, float64(32-maskSize))))

		broadcastAddr := make(net.IP, len(networkAddr))
		for j := 0; j < len(networkAddr); j++ {
			broadcastAddr[j] = networkAddr[j] | ^ipnet.Mask[j]
		}

		firstIP := make(net.IP, len(networkAddr))
		copy(firstIP, networkAddr)
		firstIP[3]++

		lastIP := make(net.IP, len(broadcastAddr))
		copy(lastIP, broadcastAddr)
		lastIP[3]--

		fmt.Printf("%-20s %-30s %-20s\n", networkAddr, firstIP.String()+" - "+lastIP.String(), broadcastAddr)
	}
}

func displayMenu() {
	fmt.Println("\n-----Options:-----")
	fmt.Println("1. Calculate subnet details from CIDR")
	fmt.Println("2. Generate IP scheme based on required hosts")
	fmt.Println("3. Exit")
	fmt.Print("Choose an option (1-3): ")
}

func handleOption(choice int) {
	switch choice {
	case 1:
		var cidr string
		fmt.Print("Enter IP address in CIDR notation (e.g., 192.168.1.1/24): ")
		fmt.Scanln(&cidr)
		calculateSubnetInfo(cidr)
	case 2:
		var baseIP string
		var requiredHosts int
		fmt.Print("Enter base IP address (e.g., 192.168.1.0): ")
		fmt.Scanln(&baseIP)
		fmt.Print("Enter the number of required hosts: ")
		fmt.Scanln(&requiredHosts)
		generateIPScheme(baseIP, requiredHosts)
	case 3:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 3.")
	}
}

func main() {
	firstRun := true
	for {
		if firstRun {
			fmt.Println("\nWelcome to the IP Subnet Calculator!")
			firstRun = false
		}

		displayMenu()

		var choice int
		fmt.Scanln(&choice)

		if choice == 3 {
			handleOption(choice)
			break
		}

		handleOption(choice)

		var anotherCalculation string
		fmt.Print("\nWould you like to perform another calculation? (y/n): ")
		fmt.Scanln(&anotherCalculation)
		if strings.ToLower(anotherCalculation) != "y" {
			fmt.Println("Exiting program. Goodbye!")
			return
		}
	}
}

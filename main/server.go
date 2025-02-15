package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	multicastAddress = "239.0.0.1:9999"
	serverPort       = "8000"
	checkMessage     = "--ETR--"
)

var serverIp string

func sendMulticast() {
	go waitForUsers()
	multicastAddr := "224.0.0.1:9999" // Example multicast address and port

	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error dialing UDP:", err)
		return
	}
	defer conn.Close()

	for {
		_, err := conn.Write([]byte(checkMessage))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func getLocalAddress() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("Ошибка при получении IP адресов:", err)
	}
	for _, addr := range addresses {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}
		if ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	log.Fatal("Не удалось найти локальный IP адрес")
	return ""
}

func waitForUsers() {
	listen, err := net.Listen("tcp", getLocalAddress()+":"+serverPort)
	if err != nil {
		log.Println(err)
		//os.Exit(1)
	}
	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			//os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// incoming request

	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}

		// write data to response
		time := time.Now().Format(time.ANSIC)
		responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", string(buffer[:]), time)
		conn.Write([]byte(responseStr))
	}

	// close conn
	conn.Close()
}

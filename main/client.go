package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func getServerIp() {
	addr, err := net.ResolveUDPAddr("udp", multicastAddress)
	if err != nil {
		fmt.Println("Ошибка при разборе адреса:", err)
		return
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Ошибка при прослушивании:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	var src *net.UDPAddr
	for {
		var n int
		n, src, err = conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении:", err)
			continue
		}
		fmt.Printf("Получено от %s: %s\n", src, string(buffer[:n]))
		if string(buffer[:n]) == checkMessage {
			break
		}
	}
	connectToServer(src.String())
}

func connectToServer(serverAddr string) {
	fmt.Println(serverAddr)
	tcpServer, err := net.ResolveTCPAddr("tcp", serverAddr[:strings.IndexByte(serverAddr, ':')]+":"+serverPort)

	if err != nil {
		fmt.Println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		fmt.Println("Dial failed:", err.Error())
		os.Exit(1)
	}

	for {

		_, err = conn.Write([]byte("This is a message"))
		if err != nil {
			fmt.Println("Write data failed:", err.Error())
			os.Exit(1)
		}

		// buffer to get data
		received := make([]byte, 1024)
		_, err = conn.Read(received)
		if err != nil {
			fmt.Println("Read data failed:", err.Error())
			os.Exit(1)
		}

		fmt.Println("Received message:", string(received))
	}

	conn.Close()
}

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
)

var serverIp string

func sendMulticast() {
	addr, err := net.ResolveUDPAddr("udp", multicastAddress)
	if err != nil {
		fmt.Println("Ошибка разрешения адреса:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Ошибка создания соединения:", err)
		return
	}
	defer conn.Close()

	for {
		message := []byte("--ETR--" + getLocalAddress() + serverPort + "--RTE--")
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println("Ошибка отправки сообщения:", err)
			return
		}

		fmt.Println("Сообщение отправлено:", string(message))
		time.Sleep(2 * time.Second) // Отправлять сообщение каждые 2 секунды
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

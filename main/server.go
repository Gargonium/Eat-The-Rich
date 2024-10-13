package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	multicastAddress = "224.0.0.1"
	multicastPort    = "9999"
	port             = "8886"
)

func SendMulticast() {

	// Резолвим адрес для мультикаста
	udpAddr, err := net.ResolveUDPAddr("udp", multicastAddress+":"+multicastPort)
	if err != nil {
		fmt.Println("Ошибка при разрешении адреса:", err)
		return
	}

	// Подключаемся к мультикастной группе
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Ошибка при подключении:", err)
		return
	}
	defer conn.Close()

	// Отправляем сообщения каждые 2 секунды
	for {
		message := "--ETR--" + getLocalAddress() + ":" + port + "--RTE--"
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Ошибка при отправке:", err)
			return
		}
		fmt.Println("Сообщение отправлено:", message)
		time.Sleep(2 * time.Second)
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

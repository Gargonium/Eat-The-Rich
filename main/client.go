package main

import (
	"fmt"
	"net"
)

func getServerIp() {
	addr, err := net.ResolveUDPAddr("udp", multicastAddress)
	if err != nil {
		fmt.Println("Ошибка разрешения адреса:", err)
		return
	}

	// Создаем сокет для прослушивания мультикаста
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Ошибка создания сокета:", err)
		return
	}
	defer conn.Close()

	// Устанавливаем буфер на соединении
	conn.SetReadBuffer(1024)

	for {
		buffer := make([]byte, 1024)
		n, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении:", err)
			return
		}

		fmt.Printf("Сообщение получено от %v: %s\n", src, string(buffer[:n]))
	}
}

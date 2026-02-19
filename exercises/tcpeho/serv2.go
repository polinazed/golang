package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Клиент подключен")

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Соединение закрыто или ошибка:", err)
			break
		}
		fmt.Print("Сообщение: " + string(buffer[:n]))
	}
}
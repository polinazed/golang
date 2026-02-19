package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	var message string
	for {
		fmt.Print("Введите сообщение: ")
		_, err := fmt.Fscan(os.Stdin, &message)
		if err != nil {
			fmt.Println("Ошибка ввода")
			break
		}
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Ошибка отправки")
			break
		}
	}
}
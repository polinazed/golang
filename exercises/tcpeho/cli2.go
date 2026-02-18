    // Клиент отправляет сообщение серверу
    // Сервер выводит это сообщение в консоль
	package main 

	import(
		"fmt"
		"net"
		"os"
	)

	func main(){
		var message string
		conn, _ := net.Dial("tcp", "127.0.0.1:4545")
		for{
			fmt.Fscan(os.Stdin, &message)           // читаем с клавиатуры
			conn.Write([]byte(message))              // отправляем
		}
		conn.Close()
		
	}
    // Клиент отправляет сообщение серверу
    // Сервер выводит это сообщение в консоль

	package main 

	import(
		"fmt"
		"net"
	)
	func main(){
		listener, _ := net.Listen("tcp", ":4545")
		for {
			conn, _ := listener.Accept()
			buffer := make([]byte, 1024)
			n, _ := conn.Read(buffer)
			fmt.Println(string(buffer[:n]))
			conn.Close()
		}
		
		
	}
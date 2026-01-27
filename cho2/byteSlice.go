package main
import "fmt"

func main(){
	b := make([]byte, 12) //байтовый срез
	fmt.Println("Byte slice:", b)

	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b)

	//вывод в виде строк
	fmt.Printf("Byte slice as text: %s\n", b)
    fmt.Println("Byte slice as text:", string(b))
	fmt.Println("Lenght of b:", len(b))
}
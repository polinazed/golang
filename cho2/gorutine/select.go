package main 
import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	ch1 := make (chan int)
	ch2 := make (chan string)
	
	go func (){
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for range ticker.C{
			ch1 <- rand.Intn(100)
		}
	}()

	go func(){
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for range ticker.C{
			ch2 <- "тик-так"
		}
	}()

	stop := time.After(5*time.Second)
	for{
		select{
		case num := <- ch1:
			fmt.Println("Число:", num)
		case msg := <- ch2:
			fmt.Println("Сообщение:", msg)
		case <- stop:
			fmt.Println("Время вышло")
			return
		}
	}
}
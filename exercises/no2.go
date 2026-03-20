package main 
import (
	"math/rand"
	"time"
	"os"
	"fmt"
)

func main(){
	rand.Seed(time.Now().UnixNano()) // инициализация генератора
	secret := rand.Intn(100) + 1      // число от 1 до 100

	var n int 
	i := 0
	for {
		i++
		fmt.Print("Введите число: ")
		_, err := fmt.Fscan(os.Stdin, &n)

		if err != nil{
			fmt.Println("Ошибка, введите число")
			fmt.Scanln()
			continue
		}

		if n == secret {
			fmt.Println("Правильно, количество попыток:", i)
			return
		}

		if n > secret{
			fmt.Println("Нет, число должно быть меньше")
		}else {
			fmt.Println("Число должно быть больше")
		}

	}
}
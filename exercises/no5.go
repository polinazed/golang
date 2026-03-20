package main 
import(
	"fmt"
	"os"
)

func main(){
	var num []int
	i := 0
	var n int 
	for {
		fmt.Print("Введите число:")
		_, err := fmt.Fscan(os.Stdin, &n)

		if err != nil{
			fmt.Println("Ошибка, введите число")
			fmt.Scanln()
			continue
		}
		if n != 0 {
			i++
			num = append(num, n)
		}else {
			break
		}
	}
	max, min := n, n 
	for _, value := range num{
		if value > max {
			max = value 
		}
		if value < min {
			min = value 
		}
	}
	fmt.Println("Результат: ")
	fmt.Println("Минимум:", min)
	fmt.Println("Максимум:", max)
	fmt.Println("Всего чисел: ", i)
}
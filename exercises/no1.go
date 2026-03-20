package main
import(
	"fmt"
	"os"
)

func main(){
	var n int
	fmt.Print("Введите число: ")
	_, err := fmt.Fscan(os.Stdin, &n)

	if err != nil {
		fmt.Println("Ошибка, надо ввести число")
		return
	}

	for i:= 1; i <= 10; i++{
		fmt.Printf("%d x %d = %d\n",n, i, n*i)
	}
}
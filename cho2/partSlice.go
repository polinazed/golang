package main
import "fmt"
func main(){
	aSlice := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(aSlice)
	l := len(aSlice)
	fmt.Println(aSlice[0:5]) //первые пять элементов 
	fmt.Println(aSlice[:5]) //тоже первые пять элементов
	fmt.Println(aSlice[l-2 : l]) //последние два элемента 
	fmt.Println(aSlice[l-2:]) //последние два элемента 

	t := aSlice[0:5:10] //первые пять элементов
	fmt.Println(len(t), cap(t))
	//элементы с индексами 2,3,4
	//емкость будет 10-2
	t = aSlice[2:5:10]
	fmt.Println(len(t), cap(t))
	//элементы с индексами 0,1,2,3,4
	//новая емкость составит 6-0
	t = aSlice[:5:6]
	fmt.Println(len(t), cap(t))
}
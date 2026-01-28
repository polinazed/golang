package main 
import "fmt"

func main(){
	a := [4]int{1,2,3,4}
	b := [4]int{5,6,6,7}
	slice := []int{}
	// объединить в срез
	slice = append(slice, a[:]...)
	slice = append(slice, b[:]...)
	fmt.Println(slice)
	// объединить в массив 
	aslice := a[:] //преобразуем в срез 
	bslice := b[:]

	abslice := append(aslice, bslice...)
	var arrab [8]int
	copy(arrab[:], abslice)
	fmt.Println(arrab)

}
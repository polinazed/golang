package main
import (
	"fmt"
)

func change(s []string){
	s[0] = "Change_function"
}

func main(){
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)
	var S0 = a[0:1]
	fmt.Println(S0)
	S0[0] = "S0"
	var S12 = a[1:3]
	fmt.Println(S12)

	S12[0] = "S12_0"
	S12[1] = "S12_1"
	fmt.Println("a:", a)
	//изменения в срезе -> изменения в массиве 
	change(S12)
	fmt.Println("a:", a)

	//емкость S0
	fmt.Println("Capacity of S0:", cap(S0), "Lenght of S0:", len(S0))
	//добавление 4 элементов в S0
	S0 = append(S0, "n1")
	S0 = append(S0, "n2")
	S0 = append(S0, "n3")
	a[0] = "-n1"
	//изменение емкости S0
	//уже другой базовый массив 
	S0 = append(S0, "n4")
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))
	
	// это изменение не относится к S0
    a[0] = "-N1-"
    // это изменение не относится к S12
    a[1] = "-N2-"

	fmt.Println("S0:", S0)
    fmt.Println("a: ", a)
    fmt.Println("S12:", S12)
}
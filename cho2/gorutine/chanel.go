package main 
import (
	"fmt"
	"math/rand"
	"sync"
)

func main(){
	var wg sync.WaitGroup 
	ch := make (chan int)
	wg.Add(1)
	go func (){
		defer wg.Done()
		for i := 0; i < 5; i++{
		val := rand.Intn(100) + 1
		if val %2 == 0 {
			ch <- val
		}
	}
	close(ch)
	}()

	ch2 := make (chan int)
	wg.Add(1)
	go func (){
		defer wg.Done()
		for num := range ch {
			ch2 <- num*2
		}
		close(ch2)
	}()
	
	ch3 := make (chan int)
	wg.Add(1)
	go func (){
		defer wg.Done()
		for num := range ch{
			ch3 <- num * 10
		}
		close(ch3)
	}()

	for num:= range ch2 {
		fmt.Println(num)
	}

	for num:= range ch3 {
		fmt.Println(num)
	}

	wg.Wait()
}
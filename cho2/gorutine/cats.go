package main 
import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup 
	names := []string{"Барсик", "Муся", "Кексик"}

	for _, n := range names{
		wg.Add(1)
		go func(name string){
			defer wg.Done()
			fmt.Printf("Мяу! Я %s\n",name )
		}(n)
	}
	wg.Wait()
}
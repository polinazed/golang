package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Я main. Старт!")
	
	var wg sync.WaitGroup  // создаём счётчик горутин
	
	// Запускаем 3 горутины
	wg.Add(1)  // говорим: будет 1 горутина
	go func() {
		defer wg.Done()  // в конце уменьшаем счётчик
		fmt.Println("Я горутина 1!")
	}()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Я горутина 2!")
	}()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Я горутина 3!")
	}()
	
	fmt.Println("Я main. Жду горутины...")
	wg.Wait()  // ждём, пока счётчик не станет 0
	
	fmt.Println("Я main. Финиш! Все горутины завершились")
}
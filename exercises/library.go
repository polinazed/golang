package main

import "fmt"

type Cat struct {
    Name     string
    Breed    string   // порода
    Age      int      // возраст в месяцах
    Mood     string   // "игривый", "сонный", "сердитый"
    IsSleeping bool
}

type Visitor struct {
    Name     string
    Age      int
    Coffee   string   // любимый кофе
}

type Cafe struct {
    Name     string
    Cats     []Cat     // срез котиков
    Visitors []Visitor // срез посетителей
    Cash     int       // выручка
}

func main(){
	cat1 := Cat{Name: "Барсик", Breed: "дворянин",Age: 8,Mood: "иргивый",IsSleeping: false}
	cat2 := Cat{Name: "Муся", Breed: "сиамская",Age: 14,Mood: "сонная",IsSleeping: true}
	cat3 := Cat{Name: "Кексик", Breed: "мейн-кун",Age: 20,Mood: "добрый",IsSleeping: false}

	visit1 := Visitor{Name: "Анна",Age: 25,Coffee: "капучино"}
	visit2 := Visitor{Name: "Петр",Age: 32,Coffee: "раф"}

	cafe1 := Cafe{ Name: "Мур-мур", Cats: []Cat{cat1, cat2, cat3}, Visitors: []Visitor{visit1, visit2}, Cash: 0}

	fmt.Printf("Кафе: %s\n", cafe1.Name)
	fmt.Println("Количество котиков:", len(cafe1.Cats))
	fmt.Printf("Имена котиков: %s, %s, %s\n", cafe1.Cats[0].Name, cat2.Name, cat3.Name)

}
package main

import (
	"fmt"
)

type Cat struct{
	name string
	breed string 
	age int
	isHungry bool 
}
type Dog struct {
	name string 
}

type Cafe struct 
{
	name string 
	cats []Cat
}

func (c Cat) Meow(){
	fmt.Printf("Мяу! Я %s\n", c.name)
}

func (c Cafe) ListCats(){
	fmt.Printf("В кафе живут: ")
	for _, cat := range c.cats{
		fmt.Printf("%s, ", cat.name)
	} 
	fmt.Println()
}

func (c Cafe) CountCats() int{
	return len(c.cats)
}

func (c *Cat) Feed(){
	if c.isHungry == true{
		c.isHungry = false 
		fmt.Printf("%s покушал и доволен\n", c.name)
	}else {
		fmt.Printf("%s не хочет есть :P\n",c.name )
	}
}

func (c *Cafe) FeedAll(){
	for i := range c.cats {
		c.cats[i].Feed()
	}
}

func (c Cafe) FindCat(name string) *Cat{
	for i, cat := range c.cats {
		if cat.name == name{
			return &c.cats[i]
		}
	}
	return nil
}

func (c Cafe) HungryCats() []Cat{
	var hungry []Cat
	for _, cat := range c.cats{
		if cat.isHungry == true{
			hungry = append(hungry, cat)
		}
	}
	return hungry
}

func (c *Cafe) RemoveCat(name string) {
	for i, cat := range c.cats{
		if cat.name == name{
			c.cats = append(c.cats[:i], c.cats[i+1:]...)
			return
		} 
	}
	fmt.Println("Такого котика нет")
}

func (c Cafe) FindMinAge() *Cat{
	minAge := 1000
	var p *Cat 
	for i, cat := range c.cats{
		if cat.age < minAge {
			minAge = cat.age
			p  = &c.cats[i]
		}
	}
	return p
}

func (d Dog) Bark(){
	fmt.Println("Гав-гав, ", d.name)
}

func (c Cat) MakeSound(){
	c.Meow()
}
func (d Dog) MakeSound(){
	d.Bark()
}

type Sounder interface {
	MakeSound()
}
func MakeNoise(s Sounder){
	fmt.Println("Разрешите пошуметь:")
	s.MakeSound()
	s.MakeSound()
}
func main(){
	cat1 := Cat{name: "Барсик", breed: "дворянин",age: 8,isHungry: false}
	cat2 := Cat{name: "Муся", breed: "сиамская",age: 14,isHungry: true}
	cat3 := Cat{name: "Кексик", breed: "мейн-кун",age: 20,isHungry: false}

	cafe := Cafe{name: "Мурка", cats: []Cat{cat1, cat2, cat3}}

	cafe.ListCats()
	fmt.Println("Всего котиков: ", cafe.CountCats())
	cat1.Meow()
	cat2.Meow()
	cat3.Meow()

	cafe.FeedAll()

	fmt.Println("Введите имя котика:")
	var nameFind string
	fmt.Scanln(&nameFind) 

	foundCat := cafe.FindCat(nameFind)

	if foundCat != nil{
		fmt.Println("Мы нашли его!")
	}else{
		fmt.Println("У нас такого нет :(")
	}

	hungryCats := cafe.HungryCats()
	if len(hungryCats) == 0{
		fmt.Println("Голодных котиков нет :Р")
	}else { 
		fmt.Printf("Голодных котиков: %d\n", len(hungryCats))
		for _, cat := range hungryCats{
			fmt.Printf("%s - голоден\n", cat.name)
		}
	}

	var catremove string 
	fmt.Println("Введите имя котика, которого хотите удалить :")
	fmt.Scanln(&catremove)
	cafe.RemoveCat(catremove)

	fmt.Println("\n--- Самый молодой котик ---")
	youngest := cafe.FindMinAge()
	if youngest != nil {
	    fmt.Printf("Самый молодой: %s, возраст %d месяцев\n", youngest.name, youngest.age)
	} else {
	    fmt.Println("В кафе нет котиков")
	}

	fmt.Println("\n--- Проверка интерфейсов ---")
	MakeNoise(cat1)  // котик должен мявкнуть дважды
	
	dog := Dog{name: "Шарик"}
	MakeNoise(dog)   // собака должна гавкнуть дважды

	animals := []Sounder{cat1,dog}
	for _, a := range animals{
		a.MakeSound()
	}
}
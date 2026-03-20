package main 

import "fmt"

type Leon struct {
	name string
	age int
	maneLength int
}

type Monkey struct{
	name string
	age int
	bananaEaten int
}

type Penguin struct{
	name string 
	age int
	swimSpeed int
}

func (l Leon) Speak(){
	fmt.Println("P-p-p, это рычит лев", l.name)
}

func (m Monkey) Speak(){
	fmt.Println("А-а-а-а, играется обезъянка", m.name)
}

func (p Penguin) Speak(){
	fmt.Println("А-а, болтает пингвин", p.name)
}

type Speaker interface{
	Speak()
}

func (l Leon) Eat (food string) string{
	return food
}

func (m Monkey) Eat (food string) string{
	return food
}

func (p Penguin) Eat (food string) string{
	return food
}

type Eater interface{
	Eat(food string) string
}

func (l Leon) Move(){
	fmt.Printf("%s, бегает\n", l.name)
}

func (m Monkey) Move(){
	fmt.Printf("%s, прыгает\n", m.name)
}

func (p Penguin) Move(){
	fmt.Printf("%s, катается и плавает\n", p.name)
}

type Mover interface{
	Move()
}

type Animals interface{
	Speaker
	Eater
	Mover
}
func main(){
	leon := Leon{name: "Гриша", age: 6, maneLength: 50}
	monk := Monkey{name: "Оля", age: 16, bananaEaten: 7}
	peng := Penguin{name: "Шкипер", age: 24, swimSpeed: 40}

	animals := []Animals{leon, monk, peng}

	for _, a := range animals{
		a.Speak()
		switch a.(type){
		case Leon: a.Eat("мясо")
		case Monkey: a.Eat("банан")
		case Penguin: a.Eat("рыба")
		}
		a.Move()
	}
}
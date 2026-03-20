package main

import (
	"database/sql"
	"fmt"
	"log"
	
	_ "github.com/lib/pq"
)

type Cat struct {
	ID       int
	Name     string
	Breed    string
	Age      int
	IsHungry bool
}

func main() {
	// 1. Подключение к базе (как в прошлый раз)
	connStr := "user=gouser password=gouser123 dbname=gocats sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer db.Close()
	
	// 2. Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal("Не могу достучаться до базы:", err)
	}
	fmt.Println("✅ Подключено к базе данных")
	
	// 3. Бесконечный цикл с меню
	for {
		fmt.Println("\n=== МЕНЮ ===")
		fmt.Println("1. Показать всех котиков")
		fmt.Println("2. Добавить котика")
		fmt.Println("3. Найти котика по имени")
		fmt.Println("4. Выйти")
		fmt.Print("Выбери действие: ")
		
		var choice int
		fmt.Scanln(&choice)
		
		switch choice {
		case 1:
			showAllCats(db)
		case 2:
			addCat(db)
		case 3:
			findCatByName(db)
		case 4:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный выбор, попробуй снова")
		}
	}
}

func showAllCats(db *sql.DB) {
	fmt.Println("Показываем всех котиков...")
	rows, err:= db.Query("SELECT id, name, breed, age, is_hungry FROM cats")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next(){
		var id, age int
		var name, hungryS, breed string 
		var hungry bool 

		err = rows.Scan(&id, &name,&breed, &age, &hungry)
		if err != nil{
			log.Fatal(err)
		}

		if hungry{
			hungryS ="голоден"
		}else {
			
			hungryS = "сыт"
		}

		fmt.Printf("ID: %d, %s (%s), возраст: %d мес, %s\n",
			id, name, breed, age, hungryS)

	}
	if err = rows.Err(); err != nil{
		log.Fatal(err)
	}
}

func addCat(db *sql.DB) {
	var name, breed string 
	var age int
	var hungryInt int

	fmt.Println("Введите данные для кота:")
	fmt.Print("Имя: ")
	fmt.Scan(&name)
	fmt.Print("Порода: ")
	fmt.Scan(&breed)
	fmt.Print("Возраст: ")
	fmt.Scan(&age)
	fmt.Print("Голоден? 1 - да, 0 - нет")
	fmt.Scan(&hungryInt)
	hungry := hungryInt == 1 // true если 1, false если 0
	fmt.Println("Добавляем котика...")

	_, err := db.Exec(
			"INSERT INTO cats (name, breed, age, is_hungry) VALUES ($1, $2, $3, $4)",
			name, breed, age, hungry,
		)
	if err != nil {
		log.Printf("Ошибка вставки котика %s: %v", name, err)
	} else {
		fmt.Printf("Котик %s добавлен в базу\n", name)
	}
}

func findCatByName(db *sql.DB) {
	fmt.Println("Ищем котика...")
	var nameF string
	fmt.Print("Имя котика: ")
	fmt.Scan(&nameF)
	rows, err:= db.Query("SELECT id, name, breed, age, is_hungry FROM cats WHERE name = $1", nameF)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	found := false

	for rows.Next(){
		var id, age int
		var name, hungryS, breed string 
		var hungry bool 

		found = true
		err = rows.Scan(&id, &name,&breed, &age, &hungry)
		if err != nil{
			log.Fatal(err)
		}

		if hungry{
			hungryS ="голоден"
		}else {
			
			hungryS = "сыт"
		}

		fmt.Printf("ID: %d, %s (%s), возраст: %d мес, %s\n",
			id, name, breed, age, hungryS)

	}

	if ! found {
		fmt.Println("Таких котиков у нас нет")
	}
	
	if err = rows.Err(); err != nil{
		log.Fatal(err)
	}
}
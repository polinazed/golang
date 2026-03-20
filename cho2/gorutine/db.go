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
	connStr := "user=gouser password=gouser123 dbname=gocats sslmode=disable"
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		log.Fatal("Не могу достучаться до базы:", err)
	}
	fmt.Println("Успешно подключились к PostgreSQL!")
	
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS cats (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		breed TEXT,
		age INT,
		is_hungry BOOLEAN DEFAULT false
	);`
	
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}
	fmt.Println("Таблица создана (или уже существовала)")
	
	cats := []Cat{
		{Name: "Барсик", Breed: "дворянин", Age: 8, IsHungry: false},
		{Name: "Муся", Breed: "сиамская", Age: 14, IsHungry: true},
		{Name: "Кексик", Breed: "мейн-кун", Age: 20, IsHungry: false},
	}
	
	for _, cat := range cats {
		_, err = db.Exec(
			"INSERT INTO cats (name, breed, age, is_hungry) VALUES ($1, $2, $3, $4)",
			cat.Name, cat.Breed, cat.Age, cat.IsHungry,
		)
		
	}
	
	rows, err := db.Query("SELECT id, name, breed, age, is_hungry FROM cats")
	if err != nil {
		log.Fatal("Ошибка запроса:", err)
	}
	defer rows.Close()
	
	fmt.Println("\nКотики из базы данных:")
	for rows.Next() {
		var cat Cat
		err = rows.Scan(&cat.ID, &cat.Name, &cat.Breed, &cat.Age, &cat.IsHungry)
		if err != nil {
			log.Fatal("Ошибка сканирования:", err)
		}
		hungryStatus := "сыт"
		if cat.IsHungry {
			hungryStatus = "голоден"
		}
		fmt.Printf("ID: %d, %s (%s), возраст: %d мес, %s\n",
			cat.ID, cat.Name, cat.Breed, cat.Age, hungryStatus)
	}
}
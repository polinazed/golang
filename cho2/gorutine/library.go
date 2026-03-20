package main

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

type Library struct{
	ID int
	Title string
	Author string
	Year int
	Rating int
	IsRead bool
}

func addBook(db *sql.DB){
	var book Library
	var read int 
	fmt.Print("Введите название книги:")
	fmt.Scan(&book.Title)
	fmt.Print("Автор:")
	fmt.Scan(&book.Author)
	fmt.Print("Год:")
	fmt.Scan(&book.Year)
	fmt.Print("Ретинг:")
	fmt.Scan(&book.Rating)
	fmt.Print("Прочитали? (1 - да , 0 - нет)")
	fmt.Scan(&read)

	book.IsRead = read == 1

	_, err := db.Exec(
			"INSERT INTO library (title, author, year, rating, is_read) VALUES ($1, $2, $3, $4)",
			book.Title, book.Author, book.Year, book.Rating, book.IsRead,
		)
	if err != nil {
		log.Printf("Ошибка вставки книги %s: %v", book.Title, err)
	} else {
		fmt.Printf("Книга %s добавлена в базу\n", book.Title)
	}


}

func showAllBooks(db *sql.DB){
	var book Library
	var isread string
	fmt.Println("Выводим все ваши книги")
	rows, err:= db.Query("SELECT id, title, author, year, rating, is_read FROM library")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next(){

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.Rating, &book.IsRead)
		if err != nil{
			log.Fatal(err)
		}

		if book.IsRead{
			isread ="прочитана"
		}else {
			
			isread = "не прочитана"
		}

		fmt.Printf("ID: %d, %s автор: %s, год издания: %d , ваша оценка: %d %s\n",
			book.ID, book.Title, book.Author, book.Year, book.Rating, isread)

	}
	if err = rows.Err(); err != nil{
		log.Fatal(err)
	}
}

func findBookByAuthor(db *sql.DB){
	var book Library
	var nameF, isread string
	fmt.Print("Введите имя автора: ")
	fmt.Scan(&nameF)
	rows, err:= db.Query("SELECT id, title, author, year, rating, is_read FROM library WHERE author = $1", nameF)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	found := false

	for rows.Next(){

		found = true
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.Rating, &book.IsRead)
		if err != nil{
			log.Fatal(err)
		}

		if book.IsRead{
			isread ="прочитана"
		}else {
			
			isread = "не прочитана"
		}

		fmt.Printf("ID: %d, %s автор: %s, год издания: %d , ваша оценка: %d %s\n",
			book.ID, book.Title, book.Author, book.Year, book.Rating, isread)

	}

	if ! found {
		fmt.Println("Такой книги нет")
	}
	
	if err = rows.Err(); err != nil{
		log.Fatal(err)
	}
}

func markAsRead(db *sql.DB){
	var id int 

	fmt.Print("Введите id книги: ")
	fmt.Scan(&id)

	rows, err:= db.Query("UPDATE books SET is_read = true WHERE id = $1", id)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

}

func main(){
	connStr := "user=gouser password=gouser123 dbname=library sslmode=disable"
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

	// createTableSQL := `
	// CREATE TABLE books (
    // id SERIAL PRIMARY KEY,
    // title TEXT NOT NULL,
    // author TEXT,
    // year INT,
    // rating INT CHECK (rating >= 1 AND rating <= 5),
    // is_read BOOLEAN DEFAULT false
	// );`
	
	// _, err = db.Exec(createTableSQL)
	// if err != nil {
	// 	log.Fatal("Ошибка создания таблицы:", err)
	// }
	// fmt.Println("Таблица создана (или уже существовала)")
	for{
		fmt.Println("\n=== МЕНЮ ===")
		fmt.Println("1. Добавить книгу")
		fmt.Println("2. Показать все книги ")
		fmt.Println("3. Найти книгу по автору ")
		fmt.Println("4. Пометить книгу как прочитанную")
		fmt.Println("5. Выход")
		fmt.Print("Выбери действие: ")
		
		var choice int
		fmt.Scanln(&choice)
		
		switch choice {
		case 1:
			addBook(db)
		case 2:
			showAllBooks(db)
		case 3:
			findBookByAuthor(db)
		case 4:
			markAsRead(db)
		case 5:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный выбор, попробуй снова")
		}
	}

}
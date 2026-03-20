package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    
    // Простой путь
    r.HandleFunc("/", homePage)
    
    // Путь с параметром {id}
    r.HandleFunc("/books/{id}", bookHandler)
    
    // Query-параметры тоже работают как обычно
    r.HandleFunc("/books", booksHandler)
    
    fmt.Println("Сервер на :8080")
    http.ListenAndServe(":8080", r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Добро пожаловать!")
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r) // получаем параметры из пути
    id := vars["id"]
    fmt.Fprintf(w, "Запрошена книга с ID: %s", id)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
    // Здесь можно читать query-параметры как обычно
    author := r.URL.Query().Get("author")
    fmt.Fprintf(w, "Книги автора: %s", author)
}
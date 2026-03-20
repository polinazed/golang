package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

type Quote struct {
    ID     int    `json:"id"`
    Text   string `json:"text"`
    Author string `json:"author"`
}

func main(){
	q1 := Quote{1, "Ты права, конец истории в книге уже предопределен. Но в нашей жизни мы не читатели, мы писатели. Мы можем изменить концовку, верно?", "Гинтама"}
	q2 := Quote{2, "Не бывает безвыходных ситуаций. Бывают ситуации, выход из которых тебя не устраивает.", "Нара Шикамару"}
	q3 := Quote{3, "Никогда не сдаваться... Встать, когда все рухнуло — вот настоящая сила.", "Наруто"}

	for 

}

func quoteH (w http.ResponseWriter, r *http.Request) {
    books := []Book{{"Война и мир", "Толстой", 1869}}
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}
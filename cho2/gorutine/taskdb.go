package main
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

type Tasks struct{
	id int
	title string
	description string
	priority int
	deadline
}

func main(){

}
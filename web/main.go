package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Name string
	Age  int
}

func main() {
	server := http.Server{
		Addr: ":9000",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, 世界!")
	})

	db, err := sql.Open("mysql", "ussii39:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	defer db.Close()
	// select文
	// rows, err := db.Query("SELECT * FROM users")
	// if err != nil{
	// 	panic(err.Error())
	// }

	server.ListenAndServe()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

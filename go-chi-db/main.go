package main

import (
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Get("/hello", func(res http.ResponseWriter, req *http.Request) {
        // var result string

        // db, err := sql.Open("mysql", "id:pass@tcp:port/dbName")
        db, err := sql.Open("mysql", "user1:1234@tcp(172.17.0.1:3306)/test")
        if err != nil {
            log.Fatal(err)
        }
        defer db.Close()
        log.Println(db)
        // db.QueryRow("select * from testTable").Scan(&value)
        insert := db.QueryRow("INSERT INTO testTable(Name) VALUES (1+1)")
        log.Println(insert)
	    // row := db.QueryRow("SELECT Name FROM testTable")
        // log.Println(row)
        // err = row.Scan(&result)
        // if err != nil && err != sql.ErrNoRows {
        //     log.Fatal(err)
        // }
        // log.Printf(result)
        // res.Write([]byte(result))
        })
        http.ListenAndServe(":8080", router)
}
    
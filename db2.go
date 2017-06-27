package main

    import (
        "database/sql"
        "fmt"
        _ "github.com/lib/pq"
        //"time"
    )

    const (
        DB_USER     = "awaresystems"
        DB_PASSWORD = "3ee798d8"
        DB_NAME     = "workshop"
    )

    func main() {
        dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
        db, err := sql.Open("postgres", dbinfo)
        checkErr(err)
        defer db.Close()

        fmt.Println("# Inserting values")
        var lastInsertId int
        err = db.QueryRow("INSERT INTO ventas(nombre,producto,precio) VALUES($1,$2,$3) returning id;", "crone", "golang", 1500).Scan(&lastInsertId)
        checkErr(err)
        fmt.Println("Ultima id insertada =", lastInsertId)
    }

    func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }

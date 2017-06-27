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
        DB_NAME     = "awaresystems"
    )

    func main() {
        dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
        db, err := sql.Open("postgres", dbinfo)
        checkErr(err)
        defer db.Close()

        fmt.Println("# Querying")
        // rows, err := db.Query("SELECT * FROM apis")
        rows, err := db.Query("SELECT url FROM apis")
        checkErr(err)

        for rows.Next() {
            // var id int
            // var nombre string
            // var descripcion string
            var url string
            // var create_at string
            // var update_at string
            // err = rows.Scan(&id, &nombre, &descripcion, &url, &create_at, &update_at)
            err = rows.Scan(&url)
            checkErr(err)
            //fmt.Println("uid | nombre | descripcion | url ")
            // fmt.Printf("%3v | %6v | %45v | %20v\n", id, nombre, descripcion, url)
            fmt.Printf("%15v |\n",url)
        }
    }

    func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }
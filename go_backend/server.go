package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        db, err := sql.Open("mysql", "user:userpassword@tcp(mysql:3306)/mydatabase")
        if err != nil {
            log.Fatal(err)
        }
        defer db.Close()

        rows, err := db.Query("SELECT name, population, region, capital FROM countries")
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var countries []map[string]interface{}
        for rows.Next() {
            var name, region, capital string
            var population int64
            if err := rows.Scan(&name, &population, &region, &capital); err != nil {
                log.Fatal(err)
            }
            country := map[string]interface{}{
                "name":      name,
                "population": population,
                "region":    region,
                "capital":   capital,
            }
            countries = append(countries, country)
        }

        if err := rows.Err(); err != nil {
            log.Fatal(err)
        }

        json.NewEncoder(w).Encode(countries)
    })

    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

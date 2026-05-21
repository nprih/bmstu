package main

import (
	"database/sql"
	"fmt"
)
import _ "modernc.org/sqlite"

func main() {
	db, err := sql.Open("sqlite", "phones")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

}

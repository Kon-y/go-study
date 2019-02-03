package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// わからなすぎて、鼻血でそうｗｗｗｗ　2019/12/3　22時37分のことである。

func main() {

	ryo, err := sql.Open("mysql", "go_user:go_user@tcp(localhost:3306)/go_apps")
	if err != nil {
		panic(err.Error())
	}
	defer ryo.Close()

	/*
		for rows.Next() {
			var id int
			var name string

			if err := rows.Scan(&id, &name); err != nil {
				panic(err.Error())
			}
			fmt.Println(id, name)
		}
	*/
}

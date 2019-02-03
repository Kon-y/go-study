package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	ryo, _ := sql.Open("mysql", "go_user:go_user@tcp(localhost:3306)/go_apps")
	defer ryo.Close()

	rows, _ := ryo.Query(
		`SELECT * FROM "tasks"`,
	)

	defer rows.Close()
	for rows.Next() {
		id, hoge := 0, ""

		fmt.Println(id, hoge)
	}
}

/*
   fmt.Println(id, name)
   	rows, _ := db.Query("SELECT * FROM tasks")
   	//fmt.Println

   	takky, _ := rows.Columns()
   	takky1 := make([]sql.RawBytes, len(takky))
   	//takky2 := make([]interface{}, len(takky1))

   	for rows.Next() {

   		// 不要なカラムデータを読み捨てる。
   		//err := rows.Scan(takky2...)

   		for i, tokuda := range takky1 {
   			fmt.Println(takky[i], ":", string(tokuda))
   		}
   	}
   }
*/

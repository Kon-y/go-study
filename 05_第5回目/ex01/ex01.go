package main

import (
	"database/sql"
	"fmt"
    "os/exec"

	_ "github.com/go-sql-driver/mysql"
)

cmd := exec.Command("docker build -t docker_sample:001 .").Run()

//err := exec.Command("docker run --rm -d --name mysql_001 --hostname mysql_001 -e MYSQL_ROOT_PASSWORD=root docker_sample:001").Run()

func main() {
	db, _ := sql.Open("mysql", "go_user:go_user@tcp(localhost:3306)/go_apps")
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM tasks")

	takky, _ := rows.Columns()
	takky1 := make([]sql.RawBytes, len(takky))
	takky2 := make([]interface{}, len(takky1))

	for i := range takky {
		takky2[i] = &takky[i]
	}

	for rows.Next() {
		err := rows.Scan(takky2...)
		if err != nil {
			fmt.Printf("Unable to Scan values from Select Query")
		}
		for i, col := range takky1 {
			fmt.Println(takky[i], ":", string(col))
		}
	}
}

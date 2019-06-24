package main

//go get gopkg.in/goracle.v2
import (
	"database/sql"
	"fmt"
)

type Result struct {
	Codigo string
}

func execute() (string, error) {
	db, err := sql.Open("goracle", "scott/tiger@127.0.0.1:1521/orclpdb1")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer db.Close()

	rows, err := db.Query("Select sysdate from dual")
	if err != nil {
		fmt.Println("error ejecutando query")
		fmt.Println(err)

		return "", err
	}
	defer rows.Close()

	var date string
	for rows.Next() {
		rows.Scan(&date)
	}
	return date, nil
}

func main() {
	fmt.Println("iniciando programa")
	resultado, err := execute()
	if err != nil {
		fmt.Println("error al ejecutar query")
		return
	}
	println("el resultado de la query es %s", resultado)
}

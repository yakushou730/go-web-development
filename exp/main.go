package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "TsengYaoShang"
	password = ""
	dbname   = "go_web_dev"
)

// func main() {
// 	t, err := template.ParseFiles("hello.gohtml")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	data := struct {
// 		Name    string
// 		Phone   string
// 		Points  []int
// 		Mapping map[string]string
// 	}{
// 		Name:   "<script>alert('Howdy!')</script>",
// 		Phone:  "3345678",
// 		Points: []int{1, 2, 3, 4, 5},
// 		Mapping: map[string]string{
// 			"A": "100",
// 			"B": "200",
// 		},
// 	}
//
// 	err = t.Execute(os.Stdout, data)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	// _, err = db.Exec(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2)`, "Shou", "yakushou730@gmail.com")
	// if err != nil {
	// 	panic(err)
	// }

	var id int
	row := db.QueryRow(`
		INSERT INTO users(name, email)
		VALUES($1, $2) RETURNING id`,
		"shou", "yakushou730@gmail.com")
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)

	db.Close()
}

// CREATE TABLE users (
// id SERIAL PRIMARY KEY,
// name TEXT,
// email TEXT NOT NULL
// );
//
// CREATE TABLE orders (
// id SERIAL PRIMARY KEY,
// user_id INT NOT NULL,
// amount INT,
// description TEXT
// );

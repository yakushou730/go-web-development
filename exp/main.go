package main

import (

	// "database/sql"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/yakushou730/go-web-development/hash"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "TsengYaoShang"
	password = ""
	dbname   = "go_web_dev"
)

type User struct {
	gorm.Model
	Name   string
	Age    int
	Email  string `gorm:"not null;unique index"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func CreateOrder(db *gorm.DB, user User, amount int, desc string) {
	db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	})
	if db.Error != nil {
		panic(db.Error)
	}
}

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
	hmac := hash.NewHMAC("my-secret-key")
	fmt.Println(hmac.Hash("this is smy string to hash"))

	// fmt.Println(rand.String(10))
	// fmt.Println(rand.RememberToken())

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"dbname=%s sslmode=disable",
	// 	host, port, user, dbname)
	// us, err := models.NewUserService(psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer us.Close()
	// us.DestructiveReset()
	//
	// user := models.User{
	// 	Name:  "Shou",
	// 	Email: "yakushou730@gmail.com",
	// }
	// if err := us.Create(&user); err != nil {
	// 	panic(err)
	// }
	//
	// user.Name = "Updated Name"
	// user.Age = 8
	// if err := us.Update(&user); err != nil {
	// 	panic(err)
	// }
	//
	// user = models.User{
	// 	Name:  "Shou",
	// 	Email: "yakushou730+1@gmail.com",
	// }
	// if err := us.Create(&user); err != nil {
	// 	panic(err)
	// }
	//
	// user.Name = "Updated Name la"
	// user.Age = 10
	// if err := us.Update(&user); err != nil {
	// 	panic(err)
	// }

	// if err := us.Delete(user.ID); err != nil {
	// 	panic(err)
	// }

	// foundUser, err := us.ByAge(8)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(foundUser)
	//
	// foundUsers, err := us.InAgeRange(0, 100)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, user := range foundUsers {
	// 	fmt.Println(user)
	// }

	// db.LogMode(true)
	// db.AutoMigrate(&User{}, &Order{})
	//
	// var user User
	// db.Preload("Orders").First(&user)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }
	// fmt.Println("Email:", user.Email)
	// fmt.Println("Number of orders:", len(user.Orders))
	// fmt.Println("Orders:", user.Orders)

	// var user User
	// db.First(&user)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }
	// CreateOrder(db, user, 1001, "Fake Description #1")
	// CreateOrder(db, user, 9999, "Fake Description #2")
	// CreateOrder(db, user, 8800, "Fake Description #3")

	// name, email := getInfo()
	//
	// u := &User{
	// 	Name:  name,
	// 	Email: email,
	// }
	//
	// if err = db.Create(u).Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", u)

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"dbname=%s sslmode=disable",
	// 	host, port, user, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	//
	// if err != nil {
	// 	panic(err)
	// }
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Successfully connected!")

	// _, err = db.Exec(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2)`, "Shou", "yakushou730@gmail.com")
	// if err != nil {
	// 	panic(err)
	// }

	// var id int
	// row := db.QueryRow(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2) RETURNING id`,
	// 	"shou", "yakushou730@gmail.com")
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(id)

	// var id int
	// var name, email string
	// rows, err := db.Query(`
	// 	SELECT id, name, email
	// 	FROM users
	// 	WHERE email=$1
	// 	OR ID > $2`, "yakushou730@gmail.com", 3)
	// if err != nil {
	// 	panic(err)
	// }
	// for rows.Next() {
	// 	rows.Scan(&id, &name, &email)
	// 	fmt.Println("ID:", id, "Name:", name, "Email:", email)
	// }

	// var id int
	// for i := 1; i < 6; i++ {
	// 	userId := 1
	// 	if i > 3 {
	// 		userId = 2
	// 	}
	// 	amount := 1000 * i
	// 	description := fmt.Sprintf("USB-C Adapter x%d", i)
	//
	// 	err := db.QueryRow(`
	// 		INSERT INTO orders (user_id, amount, description)
	// 		VALUES ($1, $2, $3)
	// 		RETURNING id`,
	// 		userId, amount, description).Scan(&id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Created an order with the ID:", id)
	// }
	//
	// db.Close()
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

// func getInfo() (name, email string) {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("What is your name?")
// 	name, _ = reader.ReadString('\n')
// 	name = strings.TrimSpace(name)
// 	fmt.Println("What is your email?")
// 	email, _ = reader.ReadString('\n')
// 	email = strings.TrimSpace(email)
// 	return name, email
// }

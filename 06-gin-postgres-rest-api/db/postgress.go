package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "users"
)

var ErrorUserNotFound error = errors.New("user not found")

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil

}

func CreateTable() {
	db, err := GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	sqlStatement := `
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INT
	);
	`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
}

func DropTable() {
	db, err := GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	sqlStatement := `
	DROP TABLE users;
	`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
}
func GetUsers() ([]User, error) {
	db, err := GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	sqlStatement := `SELECT id, name, age FROM users;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)

	}
	return users, nil
}

func GetUserByID(id int) (User, error) {
	db, err := GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	sqlStatement := `SELECT id, name, age FROM users WHERE id=$1;`
	row := db.QueryRow(sqlStatement, id)
	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return User{}, ErrorUserNotFound
	}
	return user, nil
}
func InsertUser(name string, age int) (int, error) {
	db, err := GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	sqlStatement := `
	INSERT INTO users (name, age)
	VALUES ($1, $2)
	RETURNING id
	`
	id := 0
	err = db.QueryRow(sqlStatement, name, age).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateUser(user User) error {
	db, err := GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	sqlStatement := `
	UPDATE users
	SET name = $2, age = $3
	WHERE id = $1
	`
	_, err = db.Exec(sqlStatement, user.ID, user.Name, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByID(id int) error {
	db, err := GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	sqlStatement := `
	DELETE FROM users
	WHERE id = $1
	`
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

// NOTE: Test only
// func InitTable() {
// 	DropTable()
// 	CreateTable()
// 	InsertUser("Alice", 25)
// 	InsertUser("Bob", 30)
// 	InsertUser("Charlie", 35)
// }

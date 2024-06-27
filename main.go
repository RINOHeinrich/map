package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type UserHandler struct {
	Identity map[string]string
	Password string
	Db       *sql.DB
}

func (u *UserHandler) Insert() {
	//i := 1
	keys := ""
	variables := ""
	values := ""
	// Remove the declaration and assignment of 'values'
	i := 1
	for k, v := range u.Identity {
		keys += k + ","
		variables += "$" + fmt.Sprint(i) + ","
		values += v + ","
		i++
	}
	// Remove the trailing comma from the key and value strings
	if len(keys) > 0 {
		keys += "password"
		variables += "$" + fmt.Sprint(i)
		values += u.Password
	}
	log.Println("values:", values)
	insertCmd := fmt.Sprintf("INSERT INTO users1 (%s) VALUES (%s)", keys, variables)
	log.Println("insertCmd:", insertCmd)
	// Insert the User data into the database
	valuesSlice := make([]interface{}, len(strings.Split(values, ",")))
	for i, v := range strings.Split(values, ",") {
		valuesSlice[i] = v
	}
	_, err := u.Db.Exec(insertCmd, valuesSlice...)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User data inserted successfully")
}

type UserTable struct {
	Db      *sql.DB
	Columns []string
}

func (ut *UserTable) CreateTable() {
	columns := "(id SERIAL PRIMARY KEY,"
	// Remove the declaration and assignment of 'values'
	for _, k := range ut.Columns {
		columns += fmt.Sprint(k) + " VARCHAR(250),"
	}
	columns += "password VARCHAR(50))"
	// Create the users table
	_, err := ut.Db.Exec("CREATE TABLE IF NOT EXISTS users1 (id SERIAL PRIMARY KEY, email VARCHAR(50), username VARCHAR(50), phone VARCHAR(15), password VARCHAR(50))")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users1 table created successfully")
}
func (ut *UserTable) DropTable() {
	// Drop the users table
	_, err := ut.Db.Exec("DROP TABLE users1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users table dropped successfully")
}
func (ut *UserTable) LoadColumns() {
	err := godotenv.Load("payload.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ut.Columns = strings.Split(os.Getenv("COLUMNS"), "")
}
func main() {
	//personne1 := make(map[string]string, 0)
	var handler UserHandler
	handler.Identity = make(map[string]string) // Initialize the Identity map
	handler.Password = "123456"
	// jsonDataStr is the JSON string to be mapped
	jsonDataStr := `{"email": "rabe@gmail.com", "username": "Rabe12345", "phone": "1234567890"}`
	// Open a database connection
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	handler.Db = db
	var userTable UserTable
	userTable.Db = db
	userTable.LoadColumns()
	userTable.CreateTable()
	// Unmarshal the JSON string into User.Identity map
	err = json.Unmarshal([]byte(jsonDataStr), &handler.Identity)

	if err != nil {
		log.Println("Error unmarshalling JSON data:", err)
		// Handle error
	}
	handler.Insert()
	//println("personne1:", Rabe.Identity["phone"])
}

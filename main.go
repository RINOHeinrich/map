package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type UserHandler struct {
	Identity map[string]string
	Password string
}

func (u *UserHandler) InsertUser() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
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
	insertCmd := fmt.Sprintf("INSERT INTO users (%s) VALUES (%s)", keys, variables)
	log.Println("insertCmd:", insertCmd)
	// Insert the User data into the database
	_, err = db.Exec(insertCmd, values)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User data inserted successfully")
}

func main() {
	//personne1 := make(map[string]string, 0)
	var handler UserHandler
	handler.Identity = make(map[string]string) // Initialize the Identity map
	handler.Password = "123456"
	// jsonDataStr is the JSON string to be mapped
	jsonDataStr := `{"email": "rabe@gmail.com", "username": "Rabe12345", "phone": "1234567890"}`

	// Unmarshal the JSON string into User.Identity map
	err := json.Unmarshal([]byte(jsonDataStr), &handler.Identity)

	if err != nil {
		log.Println("Error unmarshalling JSON data:", err)
		// Handle error
	}
	handler.InsertUser()
	//println("personne1:", Rabe.Identity["phone"])
}

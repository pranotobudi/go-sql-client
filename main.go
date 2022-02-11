package main

import (
	"database/sql"
	"errors"
	"fmt"

	sqlclient "github.com/pranotobudi/go-sql-client/sql-client"
)

var (
	// dbClient *sql.DB
	dbClient *sqlclient.Client
)

type User struct {
	// Id   int64
	Name string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "sql_client"
)

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	if sqlclient.IsMocked && !sqlclient.IsProduction() {
		// client := &ClientMock{}
		// return client, nil
		dbClient, err = sqlclient.Open("postgres", psqlInfo)
	} else {
		dbClient, err = sql.Open("mysql", "this is the connection string")
	}
	if err != nil {
		panic(nil)
	}
}
func main() {
	user, err := GetUser(123)
	if err != nil {
		panic(err)
	}
	// fmt.Println(user.Id)
	fmt.Println(user.Name)
}

func GetUser(id int64) (*User, error) {
	rows, err := dbClient.Query("SELECT * FROM users WHERE name='budi';")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user User
	for rows.HasNext() {
		if err := rows.Scan(&user.Name); err != nil {
			panic(err)
		}
		return &user, nil
	}
	return nil, errors.New("user ot found")

}

package session

import (
	"log"

	"github.com/srinathgs/mysqlstore"
)

var Store *mysqlstore.MySQLStore

func InitSession() {
	var err error
	Store, err = mysqlstore.NewMySQLStore("jack:password1@tcp(127.0.0.1:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local", "user_session", "/", 3600, []byte("super-secret"))
	if err != nil {
		log.Print(err)
		panic(err)
	}
}

func GetHashAndSalt() string {
	return ""
}

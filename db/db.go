package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID           uint
	Username     string `gorm:"column:username"`
	PasswordHash string `gorm:"column:passwordHash"`
	Email        string `gorm:"column:email"`
}

var (
	SqlDB *gorm.DB
)

func InitDB() {
	var err error
	dsn := "jack:password1@tcp(localhost:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local"
	SqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetUser(username string) *User {
	var user *User
	SqlDB.Where("username = ?", username).First(&user)
	return user
}

func GetUsers(username string) []User {
	var users []User
	SqlDB.Where("username = ?", username).Find(&users)
	return users
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "user"
}

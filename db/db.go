package db

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	ID           uint
// 	Username     string
// 	PasswordHash string
// 	Email        string
// }

// var (
// 	SqlDB *gorm.DB
// )

// func InitDB() {
// 	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	SqlDB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// }

// func GetUser(username string) *User {
// 	var user *User
// 	SqlDB.Where("name = ?", username).First(&user)
// 	return user
// }

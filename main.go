package main

import (
	"log"
	"net/http"

	"github.com/JackNLewis/auth-project/db"
	"github.com/srinathgs/mysqlstore"
	"golang.org/x/crypto/bcrypt"
)

var (
	store *mysqlstore.MySQLStore
)

func main() {

	// log.Print(hashAndSalt([]byte("password")))
	//pass ports as the flag

	//Initialise db connection
	db.InitDB()

	//Initialise session store
	var err error
	store, err = mysqlstore.NewMySQLStore("jack:password1@tcp(127.0.0.1:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local", "user_session", "/", 3600, []byte("super-secret"))
	if err != nil {
		log.Print(err)
		panic(err)
	}
	defer store.Close()

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/home", defaultHandler)
	http.ListenAndServe(":8080", nil)
}

// loginHandler is responsible for authenticating a user and setting the cookie
// in the response headers
func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "invalid HTTP method", http.StatusInternalServerError)
		return
	}
	r.ParseForm()

	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := db.GetUser(r.Form.Get("username"))
	if user == nil { //couldn't find user with that username
		log.Printf("could not find user %v", r.Form.Get("username"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ok := comparePasswords(user.PasswordHash, []byte(r.Form.Get("password")))
	if !ok {
		log.Printf("invalid password for %v", r.Form.Get("username"))
		http.Error(w, "invalid password", http.StatusInternalServerError)
		return
	}

	session.Values["UserID"] = user.ID
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "unable to save session", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("successfully logged in"))

}

// signupHandler checks the credentials are correct and inserts the new user into the database
func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "invalid HTTP method", http.StatusInternalServerError)
		return
	}
	r.ParseForm()

	//checks if username is already in database
	users := db.GetUsers(r.Form.Get("username"))
	if len(users) != 0 {
		http.Error(w, "username already exists", http.StatusInternalServerError)
		return
	}

	hash, err := hashAndSalt([]byte(r.Form.Get("password")))
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}

	// insert into database
	u := db.User{
		Username:     r.Form.Get("username"),
		PasswordHash: hash,
	}

	db.SqlDB.Create(&u)

	w.Write([]byte("signup successful"))
}

// defaultHandler returns the main json data which will be used to render the main page
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["UserID"] != 1234 {
		http.Redirect(w, r, "http://localhost:8080/login", http.StatusTemporaryRedirect)
		return
	}

	w.Write([]byte("home handler"))
}

func hashAndSalt(pwd []byte) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Print(err)
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

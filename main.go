package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	// log.Print(hashAndSalt([]byte("password")))
	//pass ports as the flag

	//Initialise db connection

	//Initialise session store
	// session.InitSession()
	// defer session.Store.Close()

	http.HandleFunc("/login", loginHandler)
	// http.HandleFunc("/signup", signupHandler)
	// http.HandleFunc("/home", defaultHandler)
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

	// // Get a session. Get() always returns a session, even if empty.
	// session, err := session.Store.Get(r, "user")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// user := db.GetUser(r.Form.Get("username"))
	// if user == nil { //couldn't find user with that username
	// 	log.Printf("could not find user %v", r.Form.Get("username"))
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// ok := comparePasswords(user.PasswordHash, []byte(r.Form.Get("password")))
	// if !ok {
	// 	log.Printf("invalid password for %v", r.Form.Get("username"))
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// session.Values["UserID"] = user.ID
	// // Save it before we write to the response/return from the handler.
	// err = session.Save(r, w)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

}

// signupHandler checks the credentials are correct and inserts the new user into the database
// func signupHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("signup handler"))
// }

// // defaultHandler returns the main json data which will be used to render the main page
// func defaultHandler(w http.ResponseWriter, r *http.Request) {
// 	session, err := session.Store.Get(r, "user")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	if session.Values["UserID"] != 1234 {
// 		http.Redirect(w, r, "http://localhost:8080/login", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	w.Write([]byte("home handler"))
// }

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
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

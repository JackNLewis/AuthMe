package main

import (
	"log"
	"net/http"

	"github.com/JackNLewis/auth-backend/db"
	"github.com/JackNLewis/auth-backend/session"
)

func main() {

	//pass ports as the flag

	//Initialise db connection

	//Initialise session store
	session.InitSession()
	defer session.Store.Close()

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
	session, err := session.Store.Get(r, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//retrieve username and password from database
	// if r.Form.Get("username") != "username" || r.Form.Get("password") != "password" {
	// 	log.Print("user authentification failed")
	// 	http.Error(w, "user login failed", http.StatusInternalServerError)
	// 	return
	// }

	user := db.GetUser(r.Form.Get("username"))
	if user == nil { //couldn't find user with that username
		log.Printf("could not find user %v", r.Form.Get("username"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// signupHandler checks the credentials are correct and inserts the new user into the database
func signupHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup handler"))
}

// defaultHandler returns the main json data which will be used to render the main page
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	session, err := session.Store.Get(r, "user")
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

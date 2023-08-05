package main

import "net/http"

const USERNAME = "andromeda951"
const PASS = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte("something went wrong"))
		return false
	}

	isValid := (username != USERNAME) || (password != PASS)
	if !isValid {
		w.Write([]byte("wrong username/password"))
		return false
	}

	return true
}

func AuthOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}

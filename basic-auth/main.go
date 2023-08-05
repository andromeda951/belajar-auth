package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}

	if !AuthOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, i interface{}) {
	bytes, err := json.Marshal(i)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// bytes = append(bytes, []byte("\n")[0]) // appending enter in terminal
	w.Write(bytes)
}

// RUNNING
// cd basic-auth
// go run *.go

// TESTING
// curl -X GET --user andromeda:secret http://localhost:9000/student
// output: [{"Id":"01","Name":"Andromeda","Grade":3},{"Id":"02","Name":"Budi","Grade":3},{"Id":"03","Name":"Joko","Grade":3}]

// curl -X GET --user andromeda:secret http://localhost:9000/student?id=1
// output: {"Id":"01","Name":"Andromeda","Grade":3}

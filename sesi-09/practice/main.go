package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const USERNAME = "batman"
const PASSWORD = "secret"

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}
	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}
func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/student", ActionStudent)

	var handler http.Handler = mux
	handler = Auth(handler)
	handler = AllowOnlyGET(handler)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = handler

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte(`something went wrong`))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte(`wrong username/password`))
			return
		}
		next.ServeHTTP(w, r)

	})
}

func AllowOnlyGET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("only get is allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "bourne", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "ethan", Grade: 4})
	students = append(students, &Student{Id: "s003", Name: "wick", Grade: 3})
}

package task_2

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/user", GetUsers)
	r.Post("/user", CreateUser)
	return r
}

var users []User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error read request", http.StatusInternalServerError)
		return
	}

	var user User

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "JSON error", http.StatusBadRequest)
		return
	}

	users = append(users, user)

	fmt.Fprint(w, "Create User Successful:", user.Name, user.Age)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println("JSON error", err)
		return
	}
	fmt.Fprint(w, "Users:", string(data))
}

func Server2() {
	r := Routes()

	fmt.Println("Server started on host: http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("Server NOT started: %s\n", err)
	}
}

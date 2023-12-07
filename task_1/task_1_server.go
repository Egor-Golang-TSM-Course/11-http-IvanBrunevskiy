package task_1

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func TimePage(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Fprint(w, "Current Time:", formattedTime)
}

func CustomNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not Found!")
}

func Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", HomePage)
	r.Get("/time", TimePage)
	r.NotFound(CustomNotFoundHandler)
	return r
}

func StartServer() {
	r := Routes()

	fmt.Println("Server started on host: http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("Server NOT started: %s\n", err)
	}
}

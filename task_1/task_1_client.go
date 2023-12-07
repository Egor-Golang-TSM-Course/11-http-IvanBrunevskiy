package task_1

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetHomePage() {
	url := "http://localhost:8080"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(body))
}

func GetTimePage() {
	url := "http://localhost:8080/time"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(body))
}

func GetNotFoundPage() {
	url := "http://localhost:8080/test"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(body))
}

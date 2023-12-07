package task_2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAllUsers() {
	url := "http://localhost:8080/user"

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

func CreateNewUser(UserName string, UserAge int) {
	url := "http://localhost:8080/user"

	userInfo := map[string]interface{}{
		"name": UserName,
		"age":  UserAge,
	}

	data, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("JSON error", err)
		return
	}
	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Server error", err)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(body))
}

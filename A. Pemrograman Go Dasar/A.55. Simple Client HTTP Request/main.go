package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("GET", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUser(ID string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	// Use GET request with ID as a query parameter
	request, err := http.NewRequest("GET", baseURL+"/user?id="+url.QueryEscape(ID), nil)
	if err != nil {
		return data, err
	}

	// Execute the request
	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	// Debug: Check status code
	fmt.Println("Response Status:", response.Status)
	if response.StatusCode != http.StatusOK {
		return data, fmt.Errorf("failed to fetch user, status code: %d", response.StatusCode)
	}

	// Decode JSON response
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	// var users, err = fetchUsers()
	// if err != nil {
	// 	fmt.Println("Error!", err.Error())
	// 	return
	// }

	// for _, each := range users {
	// 	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	// }

	var user1, err = fetchUser("W001")
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}

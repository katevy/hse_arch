package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	log.Println("Running tests...")
	if !runTests() {
		log.Fatal("Tests failed")
	}
	log.Println("OK")
}

func runTests() bool {
	if !testSignupEndpoint() {
		return false
	}
	if !testLoginEndpoint() {
		return false
	}
	return true
}

func testSignupEndpoint() bool {
	user := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	userBytes, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "http://localhost:8082/signup", bytes.NewBuffer(userBytes))
	if err != nil {
		log.Println(err)
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Expected status code %d, got %d\n", http.StatusOK, resp.StatusCode)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}

	expected := `{"message":"User created successfully"}`
	if string(body) != expected {
		log.Printf("Expected response %s, got %s\n", expected, string(body))
		return false
	}

	return true
}

func testLoginEndpoint() bool {
	user := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	userBytes, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "http://localhost:8082/login", bytes.NewBuffer(userBytes))
	if err != nil {
		log.Println(err)
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Expected status code %d, got %d\n", http.StatusOK, resp.StatusCode)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}

	expected := `{"token":.*}`
	matched, err := regexp.MatchString(expected, string(body))
	if err != nil {
		log.Println("Error matching string:", err)
		return false
	}
	if !matched {
		log.Printf("Expected response starting with %s, got %s\n", expected, string(body))
		return false
	}

	return true
}

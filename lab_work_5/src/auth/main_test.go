package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)

func TestSignupEndpoint(t *testing.T) {
	user := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	userBytes, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "http://host.docker.internal:8082/signup", bytes.NewBuffer(userBytes))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"message":"User created successfully"}`
	if string(body) != expected {
		t.Errorf("Expected response %s, got %s", expected, string(body))
	}
}

func TestLoginEndpoint(t *testing.T) {
	user := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	userBytes, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "http://host.docker.internal:8082/login", bytes.NewBuffer(userBytes))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"token":.*}`
	matched, err := regexp.MatchString(expected, string(body))
	if err != nil {
		t.Errorf("Error matching string:%s", err)
	}
	if !matched {
		t.Errorf("Expected response %s, got %s", expected, string(body))
	}
}

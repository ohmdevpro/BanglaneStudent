package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// Login credentials
	username := "your_username"
	password := "your_password"

	// URL for login
	url := "http://student.banglane.ac.th/home/login"

	// Create a new HTTP client
	client := &http.Client{}

	// Create the form data
	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)

	// Create the POST request
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Error:", err)
		// Handle the error as needed
		return
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		// Handle the error as needed
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Login successful!")
		// Perform further actions after successful login
	} else {
		fmt.Println("Login failed!")
		// Handle the failed login attempt
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		// Handle the error as needed
		return
	}

	// Print the response body
	fmt.Println(string(body))
}

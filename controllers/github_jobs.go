package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jobs []struct {
	githubID string `json:"githubID"`
	jobType string `json:"jobType"`
	url string     `json:"URL"`
	company string `json:"company""`
	location string `json:"location"`
	title string   `json:"title"`
	howToApply string `json:"howToApply"`
}

func  getGithubJobs(key string,pageNo int) error {
	// Build the request
	url := fmt.Sprintf("https://jobs.github.com/positions.json?description=%s&page=%d",key,pageNo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}

	// create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}

	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the data with the data from the JSON
	var data jobs

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, data)
}
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/debmalya/gomad/controllers/interfaces"
	model "github.com/debmalya/gomad/models"
	
)



// AzureController struct
type GitController struct {
	interfaces.GitController
}

func (a *GitController) GetGitHubJobs(c echo.Context) error {
    skill := c.Param("skill")
    pageNo := c.Param("pageNo")
    return c.JSON(http.StatusOK, GetGitHubJobs(skill,pageNo))
}

func GetGitHubJobs(skill string, pageNo string) []model.Jobs {
	// Build the request
	url := fmt.Sprintf("https://jobs.github.com/positions.json?description=%s&page=%s", skill, pageNo)
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
	results := []model.Jobs{}

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		fmt.Println(err)
	}

	return results
}



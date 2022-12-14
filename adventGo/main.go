package main

import (
	"adventGo/days"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getInput(day string) string {
	godotenv.Load()
	value := os.Getenv("SESSION_COOKIE")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://adventofcode.com/2022/day/"+day+"/input", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: value})

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	return sb
}

func main() {
	input := getInput("14")
	days.Day14(input)
}

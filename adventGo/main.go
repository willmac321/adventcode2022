package main

import (
	"adventGo/days"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getInput(day int) string {
	godotenv.Load()
	value := os.Getenv("SESSION_COOKIE")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://adventofcode.com/2022/day/7/input", nil)
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
	input := getInput(7)
	days.Day7(input)
}

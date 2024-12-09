package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	day, year, cookie, userAgent := ParseFlags()
	GetInput(day, year, cookie, userAgent)
}

func GetInput(day, year int, cookie string, userAgent string) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	input := RequestAOC(url, cookie, userAgent)

	filename := fmt.Sprintf("%d/day%02d/input.txt", year, day)
	WriteFile(filename, input)

	fmt.Println("Wrote input to", filename)
}

func ParseFlags() (day, year int, cookie string, userAgent string) {
	today := time.Now()

	flag.IntVar(&day, "day", today.Day(), "day number to fetch, 1-25")
	flag.IntVar(&year, "year", today.Year(), "year to fetch")

	flag.StringVar(&cookie, "cookie", os.Getenv("AOC_SESSION_TOKEN"), "your session cookie")
	flag.StringVar(&userAgent, "user-agent", os.Getenv("AOC_USER_AGENT"), "Your User-Agent header")
	flag.Parse()

	if day < 1 || day > 25 {
		log.Fatalf("day must be between 1 and 25")
	}

	if year < 2015 || year > today.Year() {
		log.Fatalf("year must be between 2015 and %d", today.Year())
	}

	if cookie == "" {
		log.Fatalf("AOC_SESSION_TOKEN environment variable is required")
	}

	return day, year, cookie, userAgent
}

func RequestAOC(url string, cookie string, userAgent string) []byte {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("making request: %s", err)
	}

	req.Header.Add("User-Agent", userAgent)

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: cookie,
	})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("requesting input: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("reading input body: %s", err)
	}

	return body
}

func WriteFile(filename string, content []byte) {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}

	err = os.WriteFile(filename, content, os.FileMode(0644))
	if err != nil {
		log.Fatalf("writing file: %s", err)
	}
}

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	arg := flag.String("url", "", "URL to shorten")
	flag.Parse()
	if *arg == "" {
		fmt.Println("Please provide a URL to shorten")
		os.Exit(1)
	}

	if !isValidURL(*arg) {
		fmt.Println("Invalid URL provided")
		os.Exit(1)
	}

	shortenURL(*arg)
}

func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}

func shortenURL(url string) {
	shortCode := generateRandomString(6)
	fmt.Printf("Shortened URL: %s -> %s\n", url, shortCode)
}

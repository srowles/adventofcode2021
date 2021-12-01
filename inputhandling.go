package adventofcode2021

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// MustStringListFromReader parses new line separated text into a list of strings
func MustStringList(data string) []string {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	return lines
}

func MustIntList(data string) []int {
	lines := MustStringList(data)
	result := make([]int, len(lines))
	for i, v := range lines {
		result[i] = MustInt(v)
	}
	return result
}

func MustInt(value string) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return v
}

func mustReadStringData(reader io.Reader) string {
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Failed to read all from reader: %v", err)
	}
	return string(data)
}

// MustReaderFromFile creates an io.Reader from supplied file or panics
func MustReaderFromFile(path string) io.Reader {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", path, err)
	}
	return f
}

// MustInputFromWebsite reads the AOC_SESSION env variable
// to form the session cookie and then reads the input for
// the appropriate day. NB says are not zero prefixed so day
// one is just "1"
//
// TODO oauth login via github etc. so I don't have to steal
// the session cookie from my browser
func MustInputFromWebsite(day string) string {
	session := strings.TrimSpace(os.Getenv("AOC_SESSION"))
	if session == "" {
		log.Fatal("AOC_SESSION env var must be ser")
	}
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var client = &http.Client{
		Timeout:   time.Second * 5,
		Transport: netTransport,
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2021/day/%s/input", day), nil)
	if err != nil {
		log.Fatalf("failed to create new request for day %s: %v", day, err)
	}
	req.Header.Set("cookie", fmt.Sprintf("session=%s", session))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to get input for day %s: %v", day, err)
	}
	defer resp.Body.Close()
	data := mustReadStringData(resp.Body)
	return data
}

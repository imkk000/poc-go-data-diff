package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/exp/slices"
)

type Data struct {
	EventID    int
	CardNumber string
}

func main() {
	currentData := make([]Data, 0)
	var mx time.Duration
	mn := time.Hour
	for t := range time.Tick(1) {
		incomingData := generateData()
		start := time.Now()
		d := diff(incomingData, currentData)
		elapsed := time.Since(start)
		currentData = append(currentData, d...)
		mx = maxDur(mx, elapsed)
		mn = minDur(mn, elapsed)
		fmt.Println(t.Format("2006-01-02 15:04:05"))
		fmt.Println("diff:", len(d))
		fmt.Println("curr:", len(currentData))
		fmt.Printf("elapsed: %s / %s / %s\n", elapsed, mn, mx)
		fmt.Println()
	}
}

func minDur(d1, d2 time.Duration) time.Duration {
	if d1 < d2 {
		return d1
	}
	return d2
}

func maxDur(d1, d2 time.Duration) time.Duration {
	if d1 < d2 {
		return d2
	}
	return d1
}

func generateData() []Data {
	rand.Seed(time.Now().UnixNano())
	l := rand.Intn(1000)
	d := make([]Data, 0)
	for l > 0 {
		l--
		d = append(d, Data{
			EventID:    10,
			CardNumber: strconv.Itoa(rand.Int()),
		})
	}
	return d
}

// diff O(n^2)
func diff(s1, s2 []Data) []Data {
	var d []Data
	for _, d1 := range s1 {
		if !slices.Contains(s2, d1) {
			d = append(d, d1)
		}
	}
	return d
}

func print(s []Data) {
	fmt.Println("total:", len(s))
	if s != nil {
		for _, d := range s {
			b, _ := json.Marshal(d)
			fmt.Println(string(b))
		}
	}
}

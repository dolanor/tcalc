package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func printUsage() {
	fmt.Printf(`Compute duration between clock times
Usage:
	%s INTERVAL [INTERVAL]...

INTERVAL are 2 clock times formatted as hh:mm and separated by a "-" dash. (eg. 9:10-11:38)
`, os.Args[0])
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Bad format of INTERVAL")
		printUsage()
		os.Exit(1)
	}
}
func main() {
	intervals := os.Args[1:]
	if len(intervals) < 1 {
		printUsage()
		os.Exit(1)
	}

	// avoid too much error handling for index access
	defer func() {
		if r := recover(); r != nil {
			checkErr(errors.New("panicked: " + fmt.Sprintf("%+v", r)))
		}
	}()

	var sum time.Duration
	for _, interval := range intervals {
		i := strings.Split(interval, "-")
		h1 := strings.Split(i[0], ":")
		h2 := strings.Split(i[1], ":")

		h1s, err := strconv.Atoi(h1[0])
		checkErr(err)
		h1e, err := strconv.Atoi(h1[1])
		checkErr(err)

		h2s, err := strconv.Atoi(h2[0])
		checkErr(err)
		h2e, err := strconv.Atoi(h2[1])
		checkErr(err)

		t1 := time.Date(0, 0, 0, h1s, h1e, 0, 0, time.UTC)
		t2 := time.Date(0, 0, 0, h2s, h2e, 0, 0, time.UTC)

		diff := t2.Sub(t1)
		fmt.Printf("%s-%s: %v\n", t1.Format("15:04"),
			t2.Format("15:04"),
			diff)
		sum += diff
	}
	fmt.Printf("Total: %v (%0.2f)\n", sum, sum.Hours())
}

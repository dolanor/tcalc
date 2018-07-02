package main

import (
	"errors"
	"fmt"
	"os"
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

	var sum time.Duration
	for _, interval := range intervals {
		i := strings.Split(interval, "-")
		if len(i) < 2 {
			checkErr(errors.New("bad interval"))
		}

		t1, err := time.Parse("15:04", i[0])
		checkErr(err)
		t2, err := time.Parse("15:04", i[1])
		checkErr(err)

		diff := t2.Sub(t1)
		fmt.Printf("%s-%s: %v\n", t1.Format("15:04"),
			t2.Format("15:04"),
			diff)
		sum += diff
	}
	fmt.Printf("Total: %v (%0.2f)\n", sum, sum.Hours())
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const layout string = "2006-01-02"

func main() {
	today := time.Now()
	count := 10

	birth_date, err := time.Parse(layout, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		count, err = strconv.Atoi(os.Args[2])
	}

	print_header()
	print_date(birth_date.Year(), birth_date)
	for i := 0; i < count; i++ {
		print_date(today.Year()+i, birth_date)
	}
}

func print_date(year int, date time.Time) {
	d := time.Date(year, date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Println(d.Format(layout) + "\t" + strconv.Itoa(year-date.Year()) + "\t" + d.Weekday().String())
}

func print_header() {
	fmt.Println("GoLang")
	fmt.Println("Date      \tAge\tDay")
	fmt.Println("----------\t---\t----------")
}

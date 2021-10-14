package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
	"time"
)

func printIssues(items []*github.Issue, desc string) {
	fmt.Println(desc)
	for _, item := range items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var monthIssues []*github.Issue
	var yearIssues []*github.Issue
	var yearsIssues []*github.Issue

	for _, item := range result.Items {
		dur := time.Since(item.CreatedAt).Hours()
		if dur <= float64(24*30) {
			monthIssues = append(monthIssues, item)
		} else if dur <= float64(24*365) {
			yearIssues = append(yearIssues, item)
		} else {
			yearsIssues = append(yearsIssues, item)
		}
	}

	printIssues(monthIssues, "=== 1ヶ月未満 ===")
	printIssues(yearIssues, "=== 1年未満 ===")
	printIssues(yearsIssues, "=== 1年以上 ===")

}

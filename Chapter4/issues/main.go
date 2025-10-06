package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func categorizeIssue(result *IssuesSearchResult) ([]*Issue, []*Issue, []*Issue) {
	var lessThanAMonth, lessThanAYear, moreThanAYear []*Issue

	for _, item := range result.Items {
		switch issueAge := daysAgo(item.CreatedAt); {
		case issueAge <= 30:
			lessThanAMonth = append(lessThanAMonth, item)
		case issueAge <= 365:
			lessThanAYear = append(lessThanAYear, item)
		default:
			moreThanAYear = append(moreThanAYear, item)
		}

	}

	return lessThanAMonth, lessThanAYear, moreThanAYear
}

// !+
func main() {
	result, err := SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	lessThanAMonth, lessThanAYear, moreThanAYear := categorizeIssue(result)

	fmt.Printf(
		"%d issues:\t %d are less than a month,\t %d are less than a year, \t %d are more than a year\n",
		result.TotalCount,
		len(lessThanAMonth),
		len(lessThanAYear),
		len(moreThanAYear),
	)

	fmt.Println()
	fmt.Println("===== issues less than a month old =====")
	for _, item := range lessThanAMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Println()
	fmt.Println("===== issues less than a year old =====")
	for _, item := range lessThanAYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Println()
	fmt.Println("===== issues more than a year old ===== ")
	for _, item := range moreThanAYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}

//!-

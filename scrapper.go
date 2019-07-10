package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type team struct {
	name  string
	stars float64
}

func getTeams() {
	teams := []team{}

	c := colly.NewCollector(
		colly.AllowedDomains("www.fifaindex.com"),
		colly.Async(true),
	)
	c.OnHTML("table.table-teams tbody tr", func(e *colly.HTMLElement) {
		name := e.ChildText(`td[data-title="Name"]>a`)
		stars := e.ChildAttrs(`td[data-title="Team Rating"]>span>i`, "class")
		teams = append(teams, team{name: name, stars: countStars(stars)})
	})

	c.Visit("https://www.fifaindex.com/teams/")
	c.Wait()
	fmt.Println(teams)
}

func countStars(classes []string) float64 {
	completeStars := 0.0
	halfStars := 0.0

	for _, c := range classes {
		if c == "fas fa-star fa-lg" {
			completeStars++
		}
		if c == "fas fa-star-half-alt fa-lg" {
			halfStars++
		}
	}

	return completeStars + (halfStars * 0.5)
}

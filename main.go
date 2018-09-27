package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	s "strings"
)

var monthMap = map[string]string{
	"January":   "01",
	"February":  "02",
	"March":     "03",
	"April":     "04",
	"May":       "05",
	"June":      "06",
	"July":      "07",
	"August":    "08",
	"September": "09",
	"October":   "10",
	"November":  "11",
	"December":  "12",
}

var mapMap = map[string]string{
	"inf":  "Inferno",
	"d2":   "Dust 2",
	"mrg":  "Mirage",
	"ovp":  "Overpass",
	"nuke": "Nuke",
	"cch":  "Cache",
	"trn":  "Train",
	"bo3":  "Bo3",
}

type match struct {
	date      []string
	matchUrl  string
	winner    string
	loser     string
	winScore  string
	loseScore string
	event     string
	num       int
	id        string
	mapName   string
}

func (m match) String() string {
	return fmt.Sprintf("%s %s > %s %s :: %s", m.winner, m.winScore, m.loseScore, m.loser, mapMap[m.mapName])
}

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("div.results-sublist", func(e *colly.HTMLElement) {
		// matches := []match{}
		matchDate := e.ChildText(".standard-headline")

		if matchDate == "" {
			return
		}
		parsedDate := parseDate(s.Split(matchDate, " "))

		fmt.Println()
		fmt.Println("[" + parsedDate[3] + " " + parsedDate[0] + ", " + parsedDate[2] + "]")

		e.ForEach("div.result-con", func(n int, el *colly.HTMLElement) {

			match := match{
				date:      parsedDate,
				matchUrl:  el.ChildAttr("a", "href"),
				winner:    el.ChildText("div.team-won"),
				loser:     el.ChildText("div.team:not(div.team-won)"),
				winScore:  el.ChildText("span.score-won"),
				loseScore: el.ChildText("span.score-lost"),
				event:     el.ChildText("span.event-name"),
				num:       n,
				id:        el.Attr("data-zonedgrouping-entry-unix"),
				mapName:   el.ChildText("div.map"),
			}
			fmt.Println(match)
			// fmt.Println("\t" + match.id)
		})
	})

	c.Visit("https://www.hltv.org/results?stars=1")
}

func parseDate(input []string) []string {
	// date format of [dd, mm, yyyy, monthName]
	date := make([]string, 4)

	r, _ := regexp.Compile("[0-9]")
	day := r.FindAllString(input[3], -1)
	if len(day) == 2 {
		date[0] = day[0] + day[1]
	} else {
		date[0] = "0" + day[0]
	}

	date[1] = monthMap[input[2]]
	date[2] = input[len(input)-1]
	date[3] = input[2]

	return date
}

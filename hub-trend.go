package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/andygrunwald/go-trending"
	"github.com/fatih/color"
)

var (
	fLang = flag.String("lang", "", "language to search: go, swift, ...")
	fTime = flag.String("time", "today", "time frame to search: today, week, month")
	fItem = flag.String("item", "proj", "item to search: proj, dev, lang")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: %v flags:[item|lang|time]\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(1)
	}

	flag.Parse()
}

func main() {
	trend := trending.NewTrending()

	optTime := trending.TimeToday
	switch *fTime {
	case "today":
		optTime = trending.TimeToday
	case "week":
		optTime = trending.TimeWeek
	case "month":
		optTime = trending.TimeMonth
	default:
		fmt.Printf("time(%s) is unknown.\n", *fTime)
		os.Exit(1)
	}

	var optLang = *fLang
	var optItem = *fItem

	switch optItem {
	case "proj":
		fmt.Println(color.RedString("Projects"), optTime, optLang)

		projects, err := trend.GetProjects(optTime, optLang)
		if err != nil {
			log.Fatal(err)
		}
		for index, project := range projects {
			no := index + 1
			if len(project.Language) > 0 {
				fmt.Printf("%d: %s (written in %s with %d ★ ) %v\n", no, color.BlueString(project.Name), color.YellowString(project.Language), project.Stars, project.URL)
			} else {
				fmt.Printf("%d: %s (with %d ★ ) %v\n", no, color.BlueString(project.Name), project.Stars, project.URL)
			}
		}

	case "dev":
		fmt.Println(color.RedString("Developers"), optTime, optLang)

		developers, err := trend.GetDevelopers(optTime, optLang)
		if err != nil {
			log.Fatal(err)
		}
		for index, developer := range developers {
			no := index + 1
			fmt.Printf("%d: %s (%s)\n", no, color.BlueString(developer.DisplayName), developer.FullName)
		}

	case "lang":
		fmt.Println(color.RedString("Languages"))

		languages, err := trend.GetLanguages()
		if err != nil {
			log.Fatal(err)
		}
		for index, language := range languages {
			no := index + 1
			fmt.Printf("%d: %s (%s)\n", no, color.BlueString(language.Name), language.URLName)
		}
	}
}

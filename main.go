package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hokita/corona_slack/corona"
	"github.com/hokita/corona_slack/webhook"
)

const (
	ExitCodeOK  = 0
	ExitCodeErr = 1
)

func main() {
	os.Exit(run())
}

func run() int {
	webhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	if webhookUrl == "" {
		fmt.Fprintln(os.Stderr, "please set SLACK_WEBHOOK_URL")
		return ExitCodeErr
	}

	crn := corona.New()

	jp, err := crn.JP()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ExitCodeErr
	}

	ranking, err := crn.WorldRanking()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ExitCodeErr
	}

	if err := webhook.Post(webhookUrl, buildTxt(jp)); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ExitCodeErr
	}

	for i, country := range ranking[:5] {
		if err := webhook.Post(webhookUrl, buildRankingTxt(country, i+1)); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return ExitCodeErr
		}
	}

	fmt.Println("Done!!")
	return ExitCodeOK
}

func buildTxt(country *corona.Country) string {
	name := fmt.Sprintf("*%s*", country.Name)
	cases := fmt.Sprintf("cases: %s", strconv.Itoa(country.Cases))
	todayCases := fmt.Sprintf("today cases: %s", strconv.Itoa(country.TodayCases))
	deaths := fmt.Sprintf("deaths: %s", strconv.Itoa(country.Deaths))
	todayDeaths := fmt.Sprintf("today deaths: %s", strconv.Itoa(country.TodayDeaths))
	text := fmt.Sprintf(`
%s
----------
%s
%s
%s
%s
----------
`,
		name, cases, todayCases, deaths, todayDeaths)

	return text
}

func buildRankingTxt(country *corona.Country, rank int) string {
	r := fmt.Sprintf("%s", strconv.Itoa(rank))
	name := fmt.Sprintf("*%s*", country.Name)
	cases := fmt.Sprintf("cases: %s", strconv.Itoa(country.Cases))
	todayCases := fmt.Sprintf("today cases: %s", strconv.Itoa(country.TodayCases))
	deaths := fmt.Sprintf("deaths: %s", strconv.Itoa(country.Deaths))
	todayDeaths := fmt.Sprintf("today deaths: %s", strconv.Itoa(country.TodayDeaths))
	text := fmt.Sprintf(`
%s. %s
----------
%s
%s
%s
%s
----------`,
		r, name, cases, todayCases, deaths, todayDeaths)

	return text
}

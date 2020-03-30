package main

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	coronaJPEndpoint      = "https://corona.lmao.ninja/countries/JP"
	coronaRankingEndpoint = "https://corona.lmao.ninja/countries?sort=cases"
	userAgent             = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"
	webhookURL            = ""
)

// Corona struct
type Corona struct {
	Country     string `json:"country"`
	Cases       int    `json:"cases"`
	TodayCases  int    `json:"todayCases"`
	Deaths      int    `json:"deaths"`
	TodayDeaths int    `json:"todayDeaths"`
}

func main() {
	jPByte := getByteFromAPI(coronaJPEndpoint)
	var jp Corona
	if err := json.Unmarshal(jPByte, &jp); err != nil {
		log.Fatal(err)
	}

	rankingByte := getByteFromAPI(coronaRankingEndpoint)
	var ranking []Corona
	if err := json.Unmarshal(rankingByte, &ranking); err != nil {
		log.Fatal(err)
	}

	sendCorona(jp)
	sendCoronas(ranking[:5])
}

func getByteFromAPI(endpoint string) []byte {
	endpointURL, _ := url.Parse(endpoint)
	req, _ := http.NewRequest("GET", endpointURL.String(), nil)
	req.Header.Add("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// fmt.Println(reflect.TypeOf(resp.Body))

	byte, _ := ioutil.ReadAll(resp.Body)

	return byte
}

func sendCorona(corona Corona) {
	country := fmt.Sprintf("*%s*", corona.Country)
	cases := fmt.Sprintf("cases: %s", strconv.Itoa(corona.Cases))
	todayCases := fmt.Sprintf("today cases: %s", strconv.Itoa(corona.TodayCases))
	deaths := fmt.Sprintf("deaths: %s", strconv.Itoa(corona.Deaths))
	todayDeaths := fmt.Sprintf("today deaths: %s", strconv.Itoa(corona.TodayDeaths))

	text := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", country, cases, todayCases, deaths, todayDeaths)

	postSlack(text)
}

func sendCoronas(coronas []Corona) {
	text := "*World Ranking*\n"
	for i, corona := range coronas {
		rank := fmt.Sprintf("rank: %s", strconv.Itoa(i+1))
		country := corona.Country
		cases := fmt.Sprintf("cases: %s", strconv.Itoa(corona.Cases))
		todayCases := fmt.Sprintf("today cases: %s", strconv.Itoa(corona.TodayCases))
		deaths := fmt.Sprintf("deaths: %s", strconv.Itoa(corona.Deaths))
		todayDeaths := fmt.Sprintf("today deaths: %s", strconv.Itoa(corona.TodayDeaths))
		text += fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n\n", rank, country, cases, todayCases, deaths, todayDeaths)
	}

	postSlack(text)
}

// postSlack function
func postSlack(text string) error {
	payload := &slack.WebhookMessage{
		Attachments: []slack.Attachment{
			{
				Blocks: []slack.Block{
					getResultSectionBlock(text),
				},
			},
		},
	}

	err := slack.PostWebhook(webhookURL, payload)
	if err != nil {
		return err
	}

	return nil
}

// getResultSectionBlock function
func getResultSectionBlock(text string) slack.Block {
	textBlockObject := slack.NewTextBlockObject(
		"mrkdwn",
		text,
		false,
		false,
	)
	section := slack.NewSectionBlock(
		textBlockObject,
		nil,
		nil,
	)
	return section
}

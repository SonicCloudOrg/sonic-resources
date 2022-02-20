package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func main() {
	var url = "https://www.theiphonewiki.com/wiki/Models"
	p, _ := goquery.NewDocument(url)
	tLen := p.Find(".wikitable").Length()
	var generationMap = make(map[string]interface{})
	for j := 0; j < tLen; j++ {
		table := p.Find(".wikitable").Eq(j)
		tbody := table.Find("tbody")
		tr := tbody.Find("tr")
		trLen := tr.Length()
		var maxIndex = 0
		var indexText string
		var oldIndexText string
		for k := 1; k < trLen; k++ {
			td := tr.Eq(k).Find("td")
			td.Each(func(i int, s *goquery.Selection) {
				if i == 0 && k > 0 {
					indexText = strings.ReplaceAll(s.Text(), "\n", "")
				}
				if k > 0 && strings.Contains(s.Text(), ",") && (strings.HasPrefix(s.Text(), "AirPods") || strings.HasPrefix(s.Text(), "AppleTV") || strings.HasPrefix(s.Text(), "Watch") || strings.HasPrefix(s.Text(), "iPhone") || strings.HasPrefix(s.Text(), "iPad") || strings.HasPrefix(s.Text(), "iPod")) {
					ss, _ := s.Html()
					if !strings.Contains(ss, "href") {
						if maxIndex == 0 {
							maxIndex = i
						}
						if i >= maxIndex {
							oldIndexText = indexText
						} else {
							indexText = oldIndexText
						}
						a := strings.Split(strings.ReplaceAll(ss, "\n", ""), "<br/>")
						for _, aa := range a {
							generationMap[aa] = indexText
						}
					}
				}
			})
		}
	}
	result, _ := json.MarshalIndent(generationMap, "\t", "")
	fmt.Println(string(result))
}

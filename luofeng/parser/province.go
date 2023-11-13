package parser

import (
	"crawlerConcurrent/engine"
	"regexp"
)

var provinceRe = regexp.MustCompile(`<a href=".+cid=+([0-9]+)" target="_blank">[\s]*?<div class="C-InfoCard G-Field">[\s]*?<div class="Content">[\s]*?<div class="Title">([^<]+)</div>`)
var provincePageRe = regexp.MustCompile(`<a title="下一页".+sheng=([0-9]+)&page=([0-9]+)">»</a>`)

func ParseProvince(contents string) engine.ParseResult {
	provinceMatches := provinceRe.FindAllStringSubmatch(contents, -1)
	pageMatches := provincePageRe.FindAllStringSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range provinceMatches {
		title := m[2]
		id := m[1]
		url := "https://www.lfgvip.com/content.php?cid=" + id
		//result.Items = append(result.Items, "Title : "+title)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(s string) engine.ParseResult {
				return ParseProfile(s, title, url, id)
			},
		})
	}

	for _, m := range pageMatches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://www.lfgvip.com/index.php?sheng=" + m[1] + "&page=" + m[2],
			ParserFunc: ParseProvince,
		})
	}

	return result
}

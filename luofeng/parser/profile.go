package parser

import (
	"crawlerConcurrent/engine"
	"crawlerConcurrent/model"
	"regexp"
	"time"
)

var (
	releaseTimeRe        = regexp.MustCompile(`<small class="pr-2"><i class="fa fa-clock-o px5" aria-hidden="true"></i>([^<]+)</small>`)
	regionRe             = regexp.MustCompile(`<span class="font-weight-bold">所属地区</span>： ([^-]+)-([^<]+)</div>`)
	ageRe                = regexp.MustCompile(`<span class="font-weight-bold">小姐年龄</span>： ([^<]+)</div>`)
	faceValueRe          = regexp.MustCompile(`<span class="font-weight-bold">小姐颜值</span>： ([^<]+)</div>`)
	expensesRe           = regexp.MustCompile(`<span class="font-weight-bold">消费水平</span>： ([^<]+)</div>`)
	serviceItemsRe       = regexp.MustCompile(`<span class="font-weight-bold">服务项目</span>： ([^<]+)</div>`)
	detailPicRe          = regexp.MustCompile(`<img class="Img G-CursorPointer w-100" src="([^"]+)" alt />`)
	detailIntroductionRe = regexp.MustCompile(`<p class="mb-2">\n([^<]+)</p>`)
	expensesReSec        = regexp.MustCompile(`([^(]+)`)
)

func ParseProfile(contents string, title string, url string, id string) engine.ParseResult {
	profile := model.Profile{}

	currentTime := time.Now()

	profile.Url = url
	profile.Id = id
	profile.Title = title
	profile.ReleaseTime = checkingProfile(contents, releaseTimeRe)
	profile.FetchTime = currentTime.Format("2006-01-02 12:10")
	profile.Province, profile.City = checkingRegion(contents, regionRe)
	profile.Age = checkingProfile(contents, ageRe)
	profile.FaceValue = checkingProfile(contents, faceValueRe)
	profile.Expenses = checkingProfile(checkingProfile(contents, expensesRe), expensesReSec)
	profile.Services = checkingProfile(contents, serviceItemsRe)
	profile.Pics = detailPics(contents, detailPicRe)
	profile.Introduction = checkingProfile(contents, detailIntroductionRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func checkingProfile(contents string, re *regexp.Regexp) string {
	cs := []byte(contents)
	match := re.FindSubmatch(cs)
	if len(match) >= 2 {
		str := string(match[1])
		return str
	} else {
		return ""
	}
}

func checkingRegion(contents string, re *regexp.Regexp) (province, city string) {
	cs := []byte(contents)
	match := re.FindSubmatch(cs)
	if len(match) >= 3 {
		province = string(match[1])
		city = string(match[2])
		return province, city
	} else {
		return "", ""
	}
}

func detailPics(contents string, re *regexp.Regexp) []string {
	var picPath []string
	matches := re.FindAllStringSubmatch(contents, -1)
	if len(matches) >= 1 {
		for _, match := range matches {
			picPath = append(picPath, match[1])
		}
	}
	return picPath
}

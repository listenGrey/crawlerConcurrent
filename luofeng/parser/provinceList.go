package parser

import (
	"crawlerConcurrent/engine"
	"regexp"
)

const provinceListRe = `.+" tag="+([1-9][0-9]*){1,3}"+.+[^>]*>([^<]+)</span></a>`

func ParseProvinceList(contents string) engine.ParseResult {
	compile := regexp.MustCompile(provinceListRe)
	matches := compile.FindAllStringSubmatch(contents, -1)

	result := engine.ParseResult{}
	//limit := 5
	for _, m := range matches {
		//result.Items = append(result.Items, "Province : "+m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://www.lfgvip.com/index.php?sheng=" + m[1],
			ParserFunc: ParseProvince,
		})
		/*limit--
		if limit == 0{
			break
		}*/
	}
	return result
}

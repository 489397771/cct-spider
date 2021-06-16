package miit

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/respont"
	"net/http"
)

func GetFirstUrl(url string) {
	fmt.Println(url) // frist url
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: "div[class='fl w100 text-center'] script",
		},
	}
	num := pr.GetPageNum()
	if num == 0 {
		num = 100
	}
	for i := 1; i < num; i++ {
		url := fmt.Sprintf("%s_%v.html", url[:len(url)-len(".html")], i)
		fmt.Println(url) // add other url
	}
}

func GetDetailPageUrl(url string, baseUrl string) {
	baseUrl = "https://xxgk.mot.gov.cn/2020/jigou/"
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse:   parse.Parse{
			BaseUrl: baseUrl,
			UrlSelector: "ul[class='fl w100 right_list'] a",
		},
	}
	pr.GetPageUrl("href")
	//req := request.Request{
	//	Url:    url,
	//	Method: http.MethodGet,
	//}
	//html, err := req.Visit()
	//if err != nil {
	//	return
	//}
	//p := parse.Parse{
	//	Html:        string(html),
	//	BaseUrl: baseUrl,
	//	UrlSelector: "ul[class='fl w100 right_list'] a",
	//}
	//hrefL := p.GetAllUrlByParseHtml("href")
	//for _, href := range hrefL{
	//	fmt.Println(baseUrl + href)
	//}
}

func GetHtmlInfo(url string) (infoMap map[string]string){
	infoMap = make(map[string]string)
	pr := respont.PR{
		Request: request.Request{
			Url : url,
			Method: http.MethodGet,
		},
		Parse:   parse.Parse{
			TitleSelector: "h1>span",
			TextSelector: "#Zoom p",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}

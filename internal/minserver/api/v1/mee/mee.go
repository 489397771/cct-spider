package mee

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/store"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
	"strings"
	"sync"
	"time"
)

func GetFirstUrl(url string, urlChan chan<- *store.UrlChan, wg *sync.WaitGroup) {
	defer wg.Done()
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     url,
			UrlSelector: "div[class='outBox zcwj']>div>a",
		},
	}
	for _, link := range pr.GetPageUrl("href"){
		urlChan <- &store.UrlChan{
			Url:     link,
			GetUrlF: GetSecondUrl,
		}
		time.Sleep(time.Second*1)
	}

}

func GetSecondUrl(url string, urlChan chan<- *store.UrlChan, infoChan chan<- *store.InfoChan) {
	s := strings.Split(url, "/")
	fmt.Println(s[4])
	switch s[4] {
	case "zyygwj", "gwywj", "xzspwj":
		urlChan <- &store.UrlChan{
			Url:     url,
			GetUrlF: GetPageUrlList,
		}
		//fmt.Println(url)
	default:
		pr := response.PR{
			Request: request.Request{
				Url:    url,
				Method: http.MethodGet,
			},
			Parse: parse.Parse{
				BaseUrl:     url,
				UrlSelector: "span[class='mobile_none']>a",
			},
		}
		for _, link := range pr.GetPageUrl("href"){
			fmt.Println(link)
			urlChan <- &store.UrlChan{
				Url:     link,
				GetUrlF: GetPageUrlList,
			}
		}
	}
}

func GetPageUrlList(url string, urlChan chan<- *store.UrlChan, infoChan chan<- *store.InfoChan) {
	fmt.Println(url) // frist url
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: ".slideTxtBoxgsf script",
		},
	}
	num := pr.GetPageNum("var countPage = [0-9]+//")
	if num == 0 {
		num = 40
	}
	for i := 1; i < num; i++ {
		//url := fmt.Sprintf("%sindex_%v.shtml", url, i)
		//fmt.Println(url) // other url
		urlChan <- &store.UrlChan{
			Url:     fmt.Sprintf("%sindex_%v.shtml", url, i),
			GetUrlF: GetDetailPageUrl,
		}
	}
}

func GetDetailPageUrl(url string, urlChan chan<- *store.UrlChan, infoChan chan<- *store.InfoChan) {
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse:   parse.Parse{
			BaseUrl: url,
			UrlSelector: "#div>li>a",
		},
	}
	//pr.GetPageUrl("href")
	for _, link := range pr.GetPageUrl("href") {
		infoChan <- &store.InfoChan{
			Url:      link,
			GetInfoF: GetHtmlInfo,
		}
	}
}


func GetHtmlInfo(url string, errChan chan <- *store.InfoChan, info chan <-map[string]string){
	pr := response.PR{
		Request: request.Request{
			Url : url,
			Method: http.MethodGet,
		},
		Parse:   parse.Parse{
			TitleSelector: "h1, h2",
			TextSelector: ".Custom_UnionStyle p, .Custom_UnionStyle div, .content_body_box>p, .content_body_box>div, .neiright_JPZ_GK_CP>p",
			DomainName: "http://www.mee.gov.cn",
		},
	}
	info <- pr.GetHtmlInfo()

	//infoMap := pr.GetHtmlInfo()
	//if len(infoMap) == 0 {
	//	errChan <- &store.InfoChan{
	//		Url:      url,
	//		GetInfoF: GetHtmlInfo,
	//	}
	//}else {
	//	info <- infoMap
	//}
}


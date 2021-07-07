package executor

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/cook"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/response"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"strings"
)

// end return the last value of a slice
func end(s []string) string {
	length := len(s)
	return s[length-1]
}

// Executor
// 查询年度指标数据的步骤
// 1、访问https://data.stats.gov.cn/easyquery.htm?cn=%s&zb=%s获取cookie，
// 其中cn代表查询类型的代码(年度、季度、月度等)，zb代表某个指标的代码(GDP、PPI、PMI等)
// 2、带着上一步产生的cookie访问URL，只需修改变量LAST-N，即查询多少年的数据
func Executor(url urllib.Param, cn, zb string) (row [][]string) {
	cookie := cook.Cookie(cn, zb)
	req := request.Request{
		URL: url.Encode(),
		Cookie: cookie,
	}
	resBody, err := req.Visit()
	if err != nil {
		fmt.Println(err)
		return
	}
	var res response.Response
	if err = json.Unmarshal(resBody, &res); err != nil {
		fmt.Println(err)
		return
	}
	nodes := res.ReturnData.DataNodes
	for _, node := range nodes {
		// Value := node.Data.Data
		StrValue := node.Data.StrData
		date := end(strings.Split(node.Code, "."))
		row = append(row, []string{date, StrValue})
	}
	return
}
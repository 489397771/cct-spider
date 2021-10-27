package response

import (
	"cct-spider-s/internal/indicator/industry/v2/pkg/net/http/request"
	"cct-spider-s/pkg/logger"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"unsafe"
)

type Respond struct {
	Date        string
	TargetValue string
}

func CrawlEastMoney(targetCode string) (row []Respond) {
	row = make([]Respond, 0)
	b, err := request.VisitEastMoney(targetCode)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var frontEnd FrontEastMoney
	if err = json.Unmarshal(b, &frontEnd); err != nil {
		logger.Error(err.Error())
		return
	}
	for _, front := range frontEnd {
		var respond Respond
		respond.TargetValue = fmt.Sprintf("%.0f", front.Value)
		respond.Date = strings.ReplaceAll(strings.Split(front.Date, "T")[0], "-", "")
		row = append(row, respond)
	}
	return
}

func CrawlSCI(pd request.PostData) (row []Respond) {
	row = make([]Respond, 0)
	b, err := request.VisitSCI(pd)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var frontEnd FrontSCI
	if err = json.Unmarshal(b, &frontEnd); err != nil {
		logger.Error(err.Error())
		return
	}
	for _, front := range frontEnd.List {
		var respond Respond
		respond.TargetValue = fmt.Sprintf("%.0f", front.MDataValue)
		respond.Date = strings.ReplaceAll(front.DataDate, "/", "")
		row = append(row, respond)
	}
	return
}

func CrawlSina(symbol string) (row []Respond) {
	row = make([]Respond, 0)
	b, err := request.VisitSina(symbol)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	str := (*string)(unsafe.Pointer(&b))
	reg := regexp.MustCompile(`".*"`)
	all := reg.FindString(*str)
	all = strings.ReplaceAll(all, `"`, "")
	allArray := strings.Split(all, "|")
	for _, v := range allArray {
		var respond Respond
		vs := strings.Split(v, ",")
		respond.TargetValue = vs[3]
		respond.Date = strings.ReplaceAll(vs[0], "-", "")
		row = append(row, respond)
	}
	return
}

// CrawlShiBor 上海银行间同业拆放利率
// https://data.eastmoney.com/shibor/shibor.aspx?m=sh&t=99&d=99221&cu=cny&type=009016&p=3
func CrawlShiBor() (row []Respond) {
	row = make([]Respond, 0)
	data := [][]string{
		{"2021-08-12", "1.9400"},
		{"2021-08-11", "2.0320"},
		{"2021-08-10", "2.2500"},
		{"2021-08-09", "2.2120"},
		{"2021-08-06", "1.8590"},
		{"2021-08-05", "1.7090"},
		{"2021-08-04", "1.7100"},
		{"2021-08-03", "1.8360"},
		{"2021-08-02", "1.8760"},
		{"2021-07-30", "2.1780"},
		{"2021-07-29", "1.6440"},
		{"2021-07-28", "2.0920"},
		{"2021-07-27", "2.2700"},
		{"2021-07-26", "2.1190"},
		{"2021-07-23", "2.0570"},
		{"2021-07-22", "2.1140"},
		{"2021-07-21", "2.1820"},
		{"2021-07-20", "2.1780"},
		{"2021-07-19", "2.1110"},
		{"2021-07-16", "2.1150"},
		{"2021-07-15", "2.0980"},
		{"2021-07-14", "1.9840"},
		{"2021-07-13", "2.0500"},
		{"2021-07-12", "1.9400"},
		{"2021-07-09", "2.2060"},
		{"2021-07-08", "1.7910"},
		{"2021-07-07", "2.0680"},
		{"2021-07-06", "1.9080"},
		{"2021-07-05", "1.6650"},
		{"2021-07-02", "1.6140"},
		{"2021-07-01", "1.7340"},
		{"2021-06-30", "2.1770"},
		{"2021-06-29", "1.7950"},
		{"2021-06-28", "1.5580"},
		{"2021-06-25", "1.5500"},
		{"2021-06-24", "1.8530"},
		{"2021-06-23", "2.2080"},
		{"2021-06-22", "2.3010"},
		{"2021-06-21", "2.2470"},
		{"2021-06-18", "2.0140"},
		{"2021-06-17", "1.8760"},
		{"2021-06-16", "2.0060"},
		{"2021-06-15", "2.1040"},
		{"2021-06-11", "1.9980"},
		{"2021-06-10", "1.8920"},
		{"2021-06-09", "2.1880"},
		{"2021-06-08", "2.1930"},
		{"2021-06-07", "2.2990"},
		{"2021-06-04", "2.1750"},
		{"2021-06-03", "1.8520"},
		{"2021-06-02", "2.0400"},
		{"2021-06-01", "2.1870"},
		{"2021-05-31", "2.2270"},
		{"2021-05-28", "2.1660"},
		{"2021-05-27", "2.1170"},
		{"2021-05-26", "1.9900"},
		{"2021-05-25", "2.1940"},
		{"2021-05-24", "2.1710"},
		{"2021-05-21", "1.9960"},
		{"2021-05-20", "2.0920"},
	}
	for _, v := range data {
		var respond Respond
		respond.Date = strings.ReplaceAll(v[0], "-", "")
		respond.TargetValue = v[1]
		row = append(row, respond)
	}
	return
}

// CrawlTBI 国债指数
// http://app.finance.ifeng.com/hq/stock_daily.php?code=sh000012
func CrawlTBI() (row []Respond) {
	row = make([]Respond, 0)
	data := [][]string{
		{"2021-08-12", "188.64", "188.67", "188.64", "188.66"},
		{"2021-08-11", "188.63", "188.64", "188.61", "188.62"},
		{"2021-08-10", "188.60", "188.63", "188.58", "188.61"},
		{"2021-08-09", "188.71", "188.72", "188.58", "188.59"},
		{"2021-08-06", "188.62", "188.67", "188.56", "188.66"},
		{"2021-08-05", "188.49", "188.61", "188.48", "188.61"},
		{"2021-08-04", "188.42", "188.48", "188.41", "188.48"},
		{"2021-08-03", "188.46", "188.47", "188.40", "188.41"},
		{"2021-08-02", "188.26", "188.54", "188.26", "188.45"},
		{"2021-07-30", "188.12", "188.22", "188.00", "188.21"},
		{"2021-07-29", "188.08", "188.10", "188.05", "188.10"},
		{"2021-07-28", "188.05", "188.07", "188.03", "188.07"},
		{"2021-07-27", "187.97", "188.09", "187.96", "188.04"},
		{"2021-07-26", "188.02", "188.10", "187.93", "187.95"},
		{"2021-07-23", "187.96", "187.97", "187.94", "187.97"},
		{"2021-07-22", "187.86", "187.95", "187.86", "187.94"},
		{"2021-07-21", "187.80", "187.85", "187.73", "187.85"},
		{"2021-07-20", "187.74", "187.79", "187.74", "187.78"},
		{"2021-07-19", "187.69", "187.73", "187.62", "187.73"},
		{"2021-07-16", "187.55", "187.65", "187.55", "187.64"},
		{"2021-07-15", "187.51", "187.53", "187.48", "187.53"},
		{"2021-07-14", "187.34", "187.50", "187.34", "187.50"},
		{"2021-07-13", "187.20", "187.33", "187.19", "187.33"},
		{"2021-07-12", "187.10", "187.19", "187.10", "187.18"},
		{"2021-07-09", "187.04", "187.06", "187.03", "187.05"},
		{"2021-07-08", "186.92", "187.03", "186.92", "187.03"},
		{"2021-07-07", "186.86", "186.89", "186.85", "186.89"},
		{"2021-07-06", "186.82", "186.84", "186.82", "186.84"},
		{"2021-07-05", "186.85", "186.85", "186.80", "186.80"},
		{"2021-07-02", "186.83", "186.83", "186.79", "186.79"},
		{"2021-07-01", "186.87", "186.89", "186.74", "186.81"},
		{"2021-06-30", "186.77", "186.93", "186.61", "186.86"},
		{"2021-06-29", "186.95", "186.98", "186.75", "186.75"},
		{"2021-06-28", "186.96", "186.96", "186.92", "186.93"},
		{"2021-06-25", "186.84", "186.91", "186.83", "186.91"},
		{"2021-06-24", "186.82", "186.83", "186.79", "186.82"},
		{"2021-06-23", "186.80", "186.81", "186.79", "186.81"},
		{"2021-06-22", "186.69", "186.80", "186.69", "186.78"},
		{"2021-06-21", "186.72", "186.73", "186.58", "186.73"},
		{"2021-06-18", "186.68", "186.70", "186.64", "186.67"},
		{"2021-06-17", "186.60", "186.68", "186.59", "186.66"},
		{"2021-06-16", "186.52", "186.58", "186.51", "186.58"},
		{"2021-06-15", "186.58", "186.59", "186.46", "186.50"},
		{"2021-06-11", "186.54", "186.54", "186.49", "186.51"},
		{"2021-06-10", "186.45", "186.53", "186.34", "186.52"},
		{"2021-06-09", "186.57", "186.59", "186.30", "186.43"},
		{"2021-06-08", "186.59", "186.59", "186.52", "186.55"},
		{"2021-06-07", "186.66", "186.66", "186.54", "186.57"},
		{"2021-06-04", "186.68", "186.69", "186.60", "186.61"},
		{"2021-06-03", "186.73", "186.73", "186.65", "186.66"},
		{"2021-06-02", "186.70", "186.72", "186.70", "186.72"},
		{"2021-06-01", "186.64", "186.69", "186.63", "186.69"},
		{"2021-05-31", "186.65", "186.67", "186.56", "186.62"},
		{"2021-05-28", "186.56", "186.60", "186.55", "186.60"},
		{"2021-05-27", "186.56", "186.58", "186.54", "186.55"},
		{"2021-05-26", "186.52", "186.55", "186.51", "186.54"},
		{"2021-05-25", "186.55", "186.57", "186.47", "186.50"},
		{"2021-05-24", "186.55", "186.56", "186.52", "186.54"},
		{"2021-05-21", "186.44", "186.50", "186.44", "186.50"},
		{"2021-05-20", "186.38", "186.43", "186.38", "186.43"},
		{"2021-05-19", "186.34", "186.37", "186.34", "186.37"},
		{"2021-05-18", "186.29", "186.33", "186.27", "186.33"},
		{"2021-05-17", "186.29", "186.30", "186.25", "186.28"},
		{"2021-05-14", "186.28", "186.30", "186.22", "186.25"},
		{"2021-05-13", "186.17", "186.27", "186.13", "186.26"},
		{"2021-05-12", "186.12", "186.15", "186.11", "186.15"},
		{"2021-05-11", "186.10", "186.25", "186.05", "186.10"},
		{"2021-05-10", "186.06", "186.08", "186.03", "186.08"},
		{"2021-05-07", "186.01", "186.05", "186.01", "186.01"},
		{"2021-05-06", "186.01", "186.02", "185.98", "186.00"},
		{"2021-04-30", "185.87", "185.92", "185.86", "185.92"},
		{"2021-04-29", "185.85", "185.88", "185.84", "185.85"},
		{"2021-04-28", "185.80", "185.84", "185.79", "185.84"},
		{"2021-04-27", "185.78", "185.85", "185.78", "185.78"},
		{"2021-04-26", "185.87", "185.90", "185.73", "185.77"},
		{"2021-04-23", "185.81", "185.83", "185.78", "185.82"},
		{"2021-04-22", "185.79", "185.80", "185.75", "185.79"},
		{"2021-04-21", "185.70", "185.78", "185.65", "185.78"},
		{"2021-04-20", "185.80", "185.83", "185.68", "185.68"},
		{"2021-04-19", "185.76", "185.78", "185.70", "185.78"},
		{"2021-04-16", "185.68", "185.73", "185.67", "185.71"},
		{"2021-04-15", "185.79", "186.05", "185.47", "185.66"},
		{"2021-04-14", "185.78", "185.88", "185.61", "185.78"},
		{"2021-04-13", "185.67", "185.83", "185.52", "185.77"},
		{"2021-04-12", "185.42", "185.79", "185.42", "185.65"},
		{"2021-04-09", "185.36", "185.59", "185.29", "185.38"},
		{"2021-04-08", "185.38", "185.61", "185.32", "185.35"},
		{"2021-04-07", "185.41", "185.58", "185.25", "185.37"},
		{"2021-04-06", "185.31", "185.58", "185.31", "185.40"},
		{"2021-04-02", "185.34", "185.36", "185.23", "185.25"},
		{"2021-04-01", "185.26", "185.34", "185.24", "185.33"},
		{"2021-03-31", "185.27", "185.28", "185.25", "185.26"},
		{"2021-03-30", "185.22", "185.26", "185.18", "185.25"},
		{"2021-03-29", "185.18", "185.22", "185.13", "185.21"},
		{"2021-03-26", "185.16", "185.17", "185.09", "185.14"},
		{"2021-03-25", "185.13", "185.23", "185.12", "185.14"},
		{"2021-03-24", "185.11", "185.16", "185.10", "185.11"},
		{"2021-03-23", "185.05", "185.13", "185.03", "185.10"},
		{"2021-03-22", "185.05", "185.06", "184.99", "185.03"},
		{"2021-03-19", "184.88", "185.01", "184.88", "185.00"},
		{"2021-03-18", "184.87", "184.87", "184.83", "184.87"},
		{"2021-03-17", "184.86", "184.88", "184.83", "184.85"},
		{"2021-03-16", "184.86", "184.86", "184.76", "184.84"},
		{"2021-03-15", "184.80", "184.84", "184.78", "184.84"},
		{"2021-03-12", "184.78", "184.79", "184.74", "184.75"},
		{"2021-03-11", "184.81", "184.81", "184.65", "184.77"},
		{"2021-03-10", "184.73", "184.80", "184.72", "184.80"},
		{"2021-03-09", "184.71", "184.72", "184.69", "184.71"},
		{"2021-03-08", "184.69", "184.73", "184.65", "184.69"},
		{"2021-03-05", "184.63", "184.65", "184.61", "184.64"},
		{"2021-03-04", "184.65", "184.70", "184.60", "184.61"},
		{"2021-03-03", "184.57", "184.66", "184.56", "184.63"},
		{"2021-03-02", "184.53", "184.55", "184.53", "184.55"},
		{"2021-03-01", "184.53", "184.54", "184.48", "184.52"},
		{"2021-02-26", "184.48", "184.52", "184.47", "184.49"},
		{"2021-02-25", "184.53", "184.53", "184.47", "184.47"},
		{"2021-02-24", "184.59", "184.63", "184.50", "184.51"},
		{"2021-02-23", "184.59", "184.62", "184.54", "184.57"},
		{"2021-02-22", "184.67", "184.67", "184.58", "184.58"},
		{"2021-02-19", "184.60", "184.62", "184.56", "184.62"},
		{"2021-02-18", "184.63", "18 .63", "184.55", "184.58"},
	}
	for _, v := range data {
		var respond Respond
		respond.Date = strings.ReplaceAll(v[0], "-", "")
		respond.TargetValue = v[4]
		row = append(row, respond)
	}
	return
}

// CrawlLPR 贷款基准利率
// https://mip.yinhang123.net/yhll/fangdaililv/1254559.html
func CrawlLPR() (row []Respond) {
	row = make([]Respond, 0)
	data := [][]string{
		{"19960501", "9.72", "0.00", "13.14", "14.94", "15.12"},
		{"19960823", "9.18", "10.08", "10.98", "11.70", "12.42"},
		{"19971023", "7.65", "8.64", "9.36", "9.90", "10.53"},
		{"19980325", "7.02", "7.92", "9.00", "9.72", "10.35"},
		{"19980701", "6.57", "6.93", "7.11", "7.65", "8.01"},
		{"19981207", "6.12", "6.39", "6.66", "7.20", "7.56"},
		{"19990610", "5.58", "5.85", "5.94", "6.03", "6.21"},
		{"20020221", "5.04", "5.31", "5.49", "5.58", "5.76"},
		{"20041029", "5.22", "5.58", "5.76", "5.82", "6.12"},
		{"20060428", "5.40", "5.85", "6.03", "6.12", "6.39"},
		{"20060819", "5.58", "6.12", "6.30", "6.48", "6.84"},
		{"20070318", "5.67", "6.39", "6.57", "6.75", "7.11"},
		{"20070519", "5.85", "6.57", "6.75", "6.93", "7.20"},
		{"20070721", "6.03", "6.84", "7.02", "7.20", "7.38"},
		{"20070822", "6.21", "7.02", "7.20", "7.38", "7.56"},
		{"20070915", "6.48", "7.29", "7.47", "7.65", "7.83"},
		{"20071221", "6.57", "7.47", "7.56", "7.74", "7.83"},
		{"20080916", "6.21", "7.20", "7.29", "7.56", "7.74"},
		{"20081008", "6.12", "6.93", "7.02", "7.29", "7.47"},
		{"20081030", "6.03", "6.66", "6.75", "7.02", "7.20"},
		{"20081127", "5.04", "5.58", "5.67", "5.94", "6.12"},
		{"20081223", "4.86", "5.31", "5.40", "5.76", "5.94"},
		{"20101020", "5.10", "5.56", "5.60", "5.96", "6.14"},
		{"20101226", "5.35", "5.81", "5.85", "6.22", "6.40"},
		{"20110209", "5.60", "6.06", "6.10", "6.45", "6.60"},
		{"20110406", "5.85", "6.31", "6.40", "6.65", "6.80"},
		{"20110707", "6.10", "6.56", "6.65", "6.90", "7.05"},
		{"20120608", "5.85", "6.31", "6.40", "6.65", "6.80"},
		{"20120706", "5.60", "6.00", "6.15", "6.40", "6.55"},
		{"20141122", "5.60", "5.60", "6.00", "6.00", "6.15"},
		{"20150301", "5.35", "5.35", "5.75", "5.75", "5.90"},
		{"20150511", "5.10", "5.10", "5.50", "5.50", "5.65"},
		{"20150628", "4.85", "4.85", "5.25", "5.25", "5.40"},
		{"20150826", "4.60", "4.60", "5.00", "5.00", "5.15"},
		{"20151024", "4.35", "4.35", "4.75", "4.75", "4.90"},
	}
	for _, v := range data {
		var respond Respond
		respond.Date = v[0]
		respond.TargetValue = v[5]
		row = append(row, respond)
	}
	return
}
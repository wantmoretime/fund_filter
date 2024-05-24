package fund_ranking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// 地址： https://fund.eastmoney.com/data/fundranking.html#tall;c0;r;s1nzf;pn50;ddesc;qsd20230523;qed20240523;qdii;zq;gg;gzbd;gzfs;bbzt;sfbb
// https://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=gp&rs=&gs=0&sc=1nzf&st=desc&sd=2023-05-23&ed=2024-05-23&qdii=&tabSubtype=,,,,,&pi=1&pn=50&dx=1&v=0.5311045589248209

var RankUrl = "https://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=gp&rs=&gs=0&sc=1nzf&st=desc&qdii=&tabSubtype=,,,,,&pi=1&pn=1000&dx=1&v=0.5311045589248209"

var Referer = "https://fund.eastmoney.com/data/fundranking.html"

type FundInc struct {
	Code           string
	Name           string
	ShortName      string
	Data           string
	NetWorth       string
	IncWorth       string
	DayIncRate     string
	WekIncRate     string
	MonIncRate     string
	SeaIncRate     string
	HYeaIncRate    string
	YeaIncRate     string
	TwoYeaIncRate  string
	YhrYeaIncRate  string
	ThisYeaIncRate string
	AllTimeIncRate string
}

type ByMonIncRate []FundInc

func (a ByMonIncRate) Len() int           { return len(a) }
func (a ByMonIncRate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMonIncRate) Less(i, j int) bool { return a[i].MonIncRate < a[j].MonIncRate }

func GetRankingData(url string) error {

	// 创建一个新的HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// 错误处理
	}
	// 添加cookie到请求中
	req.Header.Set("Referer", Referer)
	//req.AddCookie()
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// 错误处理
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	left := strings.Index(string(body), "[")
	right := strings.Index(string(body), "]")
	if left != -1 && right != -1 {
		body = body[left : right+1]
	}

	var mp []string
	json.Unmarshal(body, &mp)
	//fmt.Println(mp)
	res := make(map[string]FundInc, 100)
	var res1 []FundInc
	for _, s := range mp {
		x := strings.Split(s, ",")
		if len(x) > 0 {

			fund := FundInc{
				Code: x[0],
				Name: x[1],
				//ShortName:      x[2],
				//Data:           x[3],
				//NetWorth:       x[4],
				//IncWorth:       x[5],
				//DayIncRate:     x[6],
				//WekIncRate:     x[7],
				MonIncRate: x[8],
				//SeaIncRate:     x[9],
				//HYeaIncRate:    x[10],
				//YeaIncRate:     x[11],
				//TwoYeaIncRate:  x[12],
				//YhrYeaIncRate:  x[13],
				//ThisYeaIncRate: x[14],
				//AllTimeIncRate: x[15],
			}
			res[x[0]] = fund
			res1 = append(res1, fund)
		}
		//fmt.Println(x[0])
	}
	fmt.Println(res1)
	sort.Sort(ByMonIncRate(res1))
	fmt.Println(res1)
	return nil
}

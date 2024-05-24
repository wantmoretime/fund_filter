package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"robotwang/fund_filter/aspx"
)

// FundInfo 用于存储基金信息的结构体
type FundInfo struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// FundValue 用于存储基金净值的结构体
type FundValue struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}

func main() {
	// 假设这是天天基金网提供的获取所有基金信息的API
	//fundInfoAPI := "http://fund.eastmoney.com/js/fundcode_search.js"

	funds, err := filterByType()
	if err != nil {
		fmt.Println("fund info err:", err)
		return
	}
	fmt.Println("fund info:", funds[0].Name, funds[0].Code)

	aspx.GetNetWorth(funds[0].Code, 1)

	//fund_ranking.GetRankingData(fund_ranking.RankUrl)
}

// getFundInfo 发送HTTP请求获取基金信息列表
func getFundInfo(url string) ([]FundInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 假设返回的是一个JSON数组
	var fundInfoArray []FundInfo
	err = json.Unmarshal(body, &fundInfoArray)
	if err != nil {
		return nil, err
	}

	return fundInfoArray, nil
}

func filterByType() ([]FundInfo, error) {
	// 发起请求获取所有基金信息
	funds, err := loadFundData("混合型-偏股.txt")
	if err != nil {
		fmt.Println("fund info err:", err)
		return nil, err
	}
	fmt.Println("fund info:", funds[0].Name, funds[0].Code)
	//jsonStr, _ := json.Marshal(funds)
	//saveFundInfo("混合型-偏股_code.json", jsonStr)
	return funds, nil
}

func loadFundData(path string) ([]FundInfo, error) {
	data, err := os.ReadFile(path)
	var tmps [][]string
	// 解码JSON数据到结构体切片
	err = json.Unmarshal(data, &tmps)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	var funds []FundInfo
	for _, k := range tmps {
		funds = append(funds, FundInfo{
			Code: k[0],
			Name: k[2],
		})
	}
	return funds, nil
}

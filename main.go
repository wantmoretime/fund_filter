package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
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

	// 假设我们需要获取前一天的基金净值数据
	// 这里需要根据实际API来确定如何获取前一天的日期和净值数据
	// 以下代码仅为示例
	yesterdayFundValue, err := getFundValue(funds[0].Code, "前一天的日期")
	if err != nil {
		fmt.Println("Error getting fund value:", err)
	}

	fmt.Printf("Yesterday's value: %f\n", yesterdayFundValue.Value)
	//// 遍历基金信息，获取每个基金的净值数据
	//for _, fund := range funds {
	//	fmt.Printf("Fund: %s (%s)\n", fund[0], fund[1])
	//
	//	// 假设我们需要获取前一天的基金净值数据
	//	// 这里需要根据实际API来确定如何获取前一天的日期和净值数据
	//	// 以下代码仅为示例
	//	yesterdayFundValue, err := getFundValue(fund[1], "前一天的日期")
	//	if err != nil {
	//		fmt.Println("Error getting fund value:", err)
	//		continue
	//	}
	//
	//	fmt.Printf("Yesterday's value: %f\n", yesterdayFundValue.Value)
	//}
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

// getFundValue 获取指定基金的净值数据
func getFundValue(fundCode, endDate string) (FundValue, error) {
	// 假设这是天天基金网提供的获取基金净值数据的API
	// 注意：这里的URL和参数需要根据实际API来确定
	apiURL := fmt.Sprintf("http://fund.eastmoney.com/f10/F10DataApi.aspx?type=lsjz&code=%s&page=1", fundCode)

	resp, err := http.Get(apiURL)
	if err != nil {
		return FundValue{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return FundValue{}, err
	}
	fmt.Println(string(body))

	// 使用正则表达式提取JSON数据部分
	re := regexp.MustCompile(`(?m)(\{.*\})`)
	matches := re.FindStringSubmatch(string(body))

	if len(matches) < 2 {
		return FundValue{}, fmt.Errorf("unable to extract JSON data")
	}

	var fundValue FundValue
	err = json.Unmarshal([]byte(matches[1]), &fundValue)
	if err != nil {
		return FundValue{}, err
	}

	return fundValue, nil
}

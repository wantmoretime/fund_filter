package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// TableData 用于存储表格行数据的结构体
type TableData struct {
	Columns PValue `json:"data"`
}

type PValue struct {
	Data   string `json:"data_time"`
	Price  string `json:"day_price"`
	CPrice string `json:"accumulate_price"`
	Rate   string `json:"rate"`
}

func main() {
	// 目标ASPX页面的URL
	url := "https://fundf10.eastmoney.com/F10DataApi.aspx?type=lsjz&code=970212&page=1"

	// 发起HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// 使用goquery加载HTML文档
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 存放解析出来的所有表格数据
	var tablesData []TableData

	// 选择页面上所有的表格
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		// 存放当前表格的数据
		var tableRows []TableData

		table.Find("tr").Each(func(j int, row *goquery.Selection) {
			// 提取单元格数据
			cells := row.Find("td,th")
			var rowData []string
			cells.Each(func(k int, cell *goquery.Selection) {
				rowData = append(rowData, cell.Text())
			})
			//fmt.Println(rowData)

			// 将行数据添加到当前表格的数据中
			tableRows = append(tableRows, TableData{Columns: PValue{
				Data:   rowData[0],
				Price:  rowData[1],
				CPrice: rowData[2],
				Rate:   rowData[3],
			}})
		})

		// 将当前表格的所有行数据添加到总的表格数据中
		tablesData = append(tablesData, tableRows...)
	})

	// 将表格数据编码为JSON
	jsonBytes, err := json.Marshal(tablesData)
	if err != nil {
		log.Fatal(err)
	}

	// 输出JSON格式的数据
	fmt.Println(string(jsonBytes))
}

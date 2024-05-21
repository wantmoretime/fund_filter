package ValuationModel

import (
	"errors"
	"fmt"
)

type Stock struct {
	Symbol       string  // 股票代码
	Name         string  // 公司名称
	CurrentPrice float64 // 当前股价
	EPS          float64 // 每股收益
}

// 检查Stock是否已经初始化了必要的字段
func (s *Stock) validate() error {
	if s.CurrentPrice == 0 || s.EPS == 0 {
		return errors.New("current price and EPS must be set")
	}
	return nil
}

// CalculatePE 计算股票的市盈率
func (s *Stock) CalculatePE() float64 {
	if err := s.validate(); err != nil {
		fmt.Println(err)
		return 0
	}
	return s.CurrentPrice / s.EPS
}

// EstimateFairPrice 根据行业平均市盈率估算股票的合理价格
func (s *Stock) EstimateFairPrice(industryPE float64) float64 {
	if err := s.validate(); err != nil {
		fmt.Println(err)
		return 0
	}
	return s.EPS * industryPE
}

func PERatio() {
	// 创建一个股票实例
	myStock := Stock{
		Symbol:       "AAPL",
		Name:         "Apple Inc.",
		CurrentPrice: 150,
		EPS:          3,
	}

	// 计算市盈率
	peRatio := myStock.CalculatePE()
	fmt.Printf("The P/E Ratio for %s is: %.2f\n", myStock.Name, peRatio)

	// 假设行业平均市盈率为20
	industryAveragePE := 20.0
	fairPrice := myStock.EstimateFairPrice(industryAveragePE)
	fmt.Printf("The estimated fair price for %s based on industry P/E is: %.2f\n", myStock.Name, fairPrice)
}

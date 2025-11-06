package model

// import "time" // 已移除未使用的 time

// APIResponse represents the overall structure of the API response.
type APIResponse struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data []TradeData `json:"data"`
}

// TradeData represents a single record of electricity trade data.
type TradeData struct {
	// 與你的 PostgreSQL table schema 保持一致，並考慮 JSON 標籤
	TranDate string `json:"tranDate"` // 先用 string，入庫時再轉 time.Time
	TranHour string `json:"tranHour"`

	MarginalPrice  float64 `json:"marginalPrice"`
	RegBid         float64 `json:"regBid"`
	RegBidQse      float64 `json:"regBidQse"`
	RegBidNontrade float64 `json:"regBidNontrade"`
	RegDemand      float64 `json:"regDemand"`
	RegOffering    float64 `json:"regOffering"`
	RegPrice       float64 `json:"regPrice"`
	RegRegistered  float64 `json:"regRegistered"`
	SrBid          float64 `json:"srBid"`
	SrBidQse       float64 `json:"srBidQse"`
	SrBidNontrade  float64 `json:"srBidNontrade"`
	SrDemand       float64 `json:"srDemand"`
	SrOffering     float64 `json:"srOffering"`
	SrPrice        float64 `json:"srPrice"`
	SrRegistered   float64 `json:"srRegistered"`
	SupBid         float64 `json:"supBid"`
	SupBidQse      float64 `json:"supBidQse"`
	SupBidNontrade float64 `json:"supBidNontrade"`
	SupDemand      float64 `json:"supDemand"`
	SupOffering    float64 `json:"supOffering"`
	SupPrice       float64 `json:"supPrice"`
	SupRegistered  float64 `json:"supRegistered"`
	EdregBid       float64 `json:"edregBid"`
	EdregPrice     float64 `json:"edregPrice"`
}

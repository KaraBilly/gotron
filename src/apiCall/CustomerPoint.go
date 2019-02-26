package apiCall

import "github.com/shopspring/decimal"

type CustomerPoint struct {
	CustomerId int64           `json:"customer_id"`
	RedeemType string          `json:"redeem_type"`
	Point      decimal.Decimal `json:"point"`
}

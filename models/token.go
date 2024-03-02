package models

type Token struct {
	TokenId int    `json:"tokenId"`
	Name    string `json:"name"`
	Fee     int    `json:"fee,omitempty"`
	Address string `json:"address"`
	Decimal uint64 `json:"decimal"`
	Icon    string `json:"icon"`
	Type    int    `json:"type,omitempty"`
}

package entities

type CountryLevelRevenueResponse struct {
	Page    int                         `json:"page"`
	Limit   int                         `json:"limit"`
	Data    []*CountryLevelRevenue      `json:"data"`
}
